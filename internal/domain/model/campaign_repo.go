package model

type CampaignRepo interface {
	FindByActive(bool) ([]Campaign, error)
}
