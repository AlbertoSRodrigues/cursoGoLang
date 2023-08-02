package database

import (
	"fmt"
	"projeto/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)
	return tx.Rollback().Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Find(&campaigns)
	fmt.Println("a ...any")
	fmt.Print("b ...any")
	print("teste3")
	println("teste4")
	return campaigns, tx.Error
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := c.Db.First(&campaign, id)
	fmt.Println("a ...any")
	fmt.Print("b ...any")
	print("teste1")
	println("teste2")
	return &campaign, tx.Error
}
