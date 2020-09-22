package repository

import (
	"blue-discount/internal/domain/model"

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
	var ms []model.Campaign
	err := r.db.Where(`active = ?`, active).Find(&ms).Error
	return ms, err
}
