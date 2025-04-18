package converter

import (
	"collab-be/internal/entity"
	"collab-be/internal/model"
)

func JobToResponse(job *entity.Job) *model.JobResponse {
	return &model.JobResponse{
		ID:          &job.ID,
		Title:       job.Title,
		Position:    job.Position,
		Description: job.Description,
		Location:    job.Location,
		Salary:      job.Salary,
		CreatedAt:   &job.CreatedAt,
		UpdatedAt:   &job.UpdatedAt,
		Company:     UserToResponse(&job.Company),
		Skills:      SkillsToResponse(job.Skills),
	}
}
func SkillsToResponse(skills []*entity.Skill) []model.SkillResponse {
	if skills == nil || len(skills) == 0 {
		return []model.SkillResponse{} // Return empty array, not null
	}

	responses := make([]model.SkillResponse, len(skills))
	for i, skill := range skills {
		// Create a copy of the ID to get its address
		id := skill.ID
		responses[i] = model.SkillResponse{
			ID:        &id, // Pass the address of the copy
			SkillName: skill.SkillName,
		}
	}
	return responses
}
