package campaign

import (
	"errors"
	"projeto/internal/contract"
	internalerrors "projeto/internal/internal-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(idToSearch string) (*contract.ReadCampaign, error)
	Cancel(idToCancel string) error
	Delete(idToDelete string) error
}
type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Create(campaign)
	if err != nil {
		return "", internalerrors.ErrInteral
	}

	return campaign.ID, nil
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
func (s *ServiceImp) Cancel(idToCancel string) error {
	selectedCampaign, err := s.Repository.GetBy(idToCancel)
	if err != nil {
		return internalerrors.ErrInteral
	}

	if selectedCampaign.Status != "pending" {
		return errors.New("Campaign status invalid")
	}

	selectedCampaign.Cancel()
	err = s.Repository.Update(selectedCampaign)
	if err != nil {
		return internalerrors.ErrInteral
	}
	return nil
}
func (s *ServiceImp) Delete(idToDelete string) error {
	selectedCampaign, err := s.Repository.GetBy(idToDelete)
	if err != nil {
		return internalerrors.ErrInteral
	}

	if selectedCampaign.Status != "pending" {
		return errors.New("Campaign status invalid")
	}

	selectedCampaign.Delete()
	err = s.Repository.Delete(selectedCampaign)
	if err != nil {
		return internalerrors.ErrInteral
	}
	return nil
}
