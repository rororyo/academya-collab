package usecase

import (
	"collab-be/internal/entity"
	"collab-be/internal/model"
	"collab-be/internal/model/converter"
	"collab-be/internal/repository"
	"context"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JobUsecase struct {
	DB              *gorm.DB
	Log             *logrus.Logger
	Validate        *validator.Validate
	JobRepository   *repository.JobRepository
	SkillRepository *repository.SkillRepository
}

func NewJobUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, jobRepository *repository.JobRepository, skillRepository *repository.SkillRepository) *JobUsecase {
	return &JobUsecase{
		DB:              db,
		Log:             log,
		Validate:        validate,
		JobRepository:   jobRepository,
		SkillRepository: skillRepository,
	}
}

func (c *JobUsecase) Create(ctx context.Context, request *model.JobRequest) (*model.JobResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		c.Log.Warnf("Failed to start transaction: %+v", tx.Error)
		return nil, fiber.ErrInternalServerError
	}
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Create job entity
	job := &entity.Job{
		CompanyID:   *request.CompanyID,
		Title:       request.Title,
		Position:    request.Position,
		Description: request.Description,
		Location:    request.Location,
		Salary:      request.Salary,
	}
	var skillEntities []*entity.Skill

	for _, skillName := range request.Skills {
		skill, err := c.SkillRepository.FindByName(tx, skillName)
		if err != nil && err != gorm.ErrRecordNotFound {
			c.Log.Warnf("Failed to find skill: %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		if skill == nil {
			// Skill doesn't exist, create it
			skill = &entity.Skill{
				SkillName: skillName,
			}
			if err := c.SkillRepository.Create(tx, skill); err != nil {
				c.Log.Warnf("Failed to create skill: %+v", err)
				return nil, fiber.ErrInternalServerError
			}
		}

		skillEntities = append(skillEntities, skill)
	}

	job.Skills = skillEntities

	// Insert into DB
	if err := c.JobRepository.Create(tx, job); err != nil {
		c.Log.Warnf("Failed to create job: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Could not create job")
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction: %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Return created job
	return converter.JobToResponse(job), nil
}

func (c *JobUsecase) Get(ctx context.Context, request *model.GetJobRequest) (*model.JobResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}
	job := new(entity.Job)
	if err := c.JobRepository.FindById(tx, job, request.ID); err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		return nil, fiber.ErrNotFound
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.JobToResponse(job), nil

}

func (c *JobUsecase) Update(ctx context.Context, request *model.UpdateJobRequest) (*model.JobResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Find job by ID
	job := new(entity.Job)
	if err := c.JobRepository.FindById(tx, job, request.ID); err != nil {
		c.Log.Warnf("Failed to find job by ID : %+v", err)
		return nil, fiber.ErrNotFound
	}

	// Verify job ownership
	if err := c.JobRepository.VerifyJobOwnership(tx, request.ID, request.CompanyID); err != nil {
		c.Log.Warnf("Unauthorized attempt to update job : %+v", err)
		return nil, fiber.ErrForbidden
	}

	// Update fields if provided
	if request.Title != "" {
		job.Title = request.Title
	}
	if request.Position != "" {
		job.Position = request.Position
	}
	if request.Description != "" {
		job.Description = request.Description
	}
	if request.Location != "" {
		job.Location = request.Location
	}
	if request.Salary > 0 {
		job.Salary = request.Salary
	}
	now := time.Now()
	job.UpdatedAt = now

	// Update job in database
	if err := c.JobRepository.Update(tx, job); err != nil {
		c.Log.Warnf("Failed to update job : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.JobToResponse(job), nil
}

func (c *JobUsecase) Delete(ctx context.Context, request *model.DeleteJobRequest) (*model.JobResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Find job by ID
	job := new(entity.Job)
	if err := c.JobRepository.FindById(tx, job, request.ID); err != nil {
		c.Log.Warnf("Failed to find job by ID : %+v", err)
		return nil, fiber.ErrNotFound
	}

	// Verify job ownership
	if err := c.JobRepository.VerifyJobOwnership(tx, request.ID, request.CompanyID); err != nil {
		c.Log.Warnf("Unauthorized attempt to delete job : %+v", err)
		return nil, fiber.ErrForbidden // Returns fiber.ErrForbidden if user is not the owner
	}

	// Delete job
	if err := c.JobRepository.Delete(tx, job); err != nil {
		c.Log.Warnf("Failed to delete job : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed to commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.JobToResponse(job), nil
}

func (c *JobUsecase) Search(ctx context.Context, request *model.SearchJobRequest) ([]model.JobResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Warnf("Invalid request body")
		return nil, 0, fiber.ErrBadRequest
	}
	jobs, total, err := c.JobRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to search job")
		return nil, 0, fiber.ErrInternalServerError
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.JobResponse, len(jobs))
	for i, job := range jobs {
		responses[i] = *converter.JobToResponse(&job)
	}
	return responses, total, nil
}
