package endpoints

import "projeto/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
