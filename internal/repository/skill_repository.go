package repository

import (
	"collab-be/internal/entity"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SkillRepository struct {
	Repository[entity.Skill]
	Log *logrus.Logger
}

func NewSkillRepository(log *logrus.Logger) *SkillRepository {
	return &SkillRepository{
		Log: log,
	}
}

// FindByName searches for a skill by name
func (r *SkillRepository) FindByName(tx *gorm.DB, name string) (*entity.Skill, error) {
	var skill entity.Skill
	if err := tx.Where("skill_name = ?", name).First(&skill).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		r.Log.Warnf("Failed to find skill by name: %+v", err)
		return nil, err
	}
	return &skill, nil
}

// FindOrCreateByName checks if a skill exists, and creates it if not
func (r *SkillRepository) FindOrCreateByName(tx *gorm.DB, name string) (*entity.Skill, error) {
	skill, err := r.FindByName(tx, name)
	if err != nil {
		return nil, err
	}
	if skill != nil {
		return skill, nil
	}
	newSkill := &entity.Skill{
		SkillName: name,
	}
	if err := r.Create(tx, newSkill); err != nil {
		return nil, err
	}
	return newSkill, nil
}
