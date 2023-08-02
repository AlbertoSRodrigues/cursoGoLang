package campaign

import (
	"projeto/internal/contract"
	internalerrors "projeto/internal/internal-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(idToSearch string) (*contract.ReadCampaign, error)
}
type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) GetBy(idToSearch string) (*contract.ReadCampaign, error) {
	selectedCampaign, err := s.Repository.GetBy(idToSearch)
	if err != nil {
		return nil, internalerrors.ErrInteral
	}

	return &contract.ReadCampaign{
		Id:      selectedCampaign.ID,
		Status:  selectedCampaign.Status,
		Name:    selectedCampaign.Name,
		Content: selectedCampaign.Content,
	}, nil

}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInteral
	}

	return campaign.ID, nil
}
