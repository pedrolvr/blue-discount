package repository

import (
	"blue-discount/internal/domain/model"
	"fmt"

	"gorm.io/gorm"
)

type CampaignRepo struct {
	db *gorm.DB
}

func NewCampaignRepo(db *gorm.DB) model.CampaignRepo {
	return &CampaignRepo{
		db: db,
	}
}

func (r *CampaignRepo) FindByActive(active bool) ([]model.Campaign, error) {
	var m []model.Campaign

	err := r.db.Where(`active = ?`, active).
		Order("priority ASC").
		Find(&m).Error

	if err != nil {
		return m, fmt.Errorf("find by active: %w", err)
	}

	return m, nil
}
