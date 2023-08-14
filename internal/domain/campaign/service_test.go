package campaign

import (
	"errors"
	"projeto/internal/contract"
	internalerrors "projeto/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
func (r *repositoryMock) Update(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
func (r *repositoryMock) Delete(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

func (r *repositoryMock) GetBy(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), nil
}

var (
	campanhaTeste = contract.NewCampaign{
		Name:    "campanha",
		Content: "Bodys",
		Emails:  []string{"email1@g.com", "email2@g.com", "email3@g.com"},
	}

	service = ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(campanhaTeste)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_ValidateDomainError_Campaign(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerrors.ErrInteral, err))
}

func Test_Save_Campaign(t *testing.T) {

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != campanhaTeste.Name ||
			campaign.Content != campanhaTeste.Content ||
			len(campaign.Contacts) != len(campanhaTeste.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repositoryMock

	service.Create(campanhaTeste)

	repositoryMock.AssertExpectations(t)
}

func Test_Save_Repository(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(internalerrors.ErrInteral)
	service.Repository = repositoryMock

	_, err := service.Create(campanhaTeste)

	assert.True(errors.Is(internalerrors.ErrInteral, err))
}

func Test_Get_CampaignById(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(campanhaTeste.Name, campanhaTeste.Content, campanhaTeste.Emails)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(campaign, nil)
	service.Repository = repositoryMock

	returnedCampaign, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, returnedCampaign.Id)
}

func Test_Get_CampignById_Error(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(campanhaTeste.Name, campanhaTeste.Content, campanhaTeste.Emails)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetBy", mock.Anything).Return(nil, errors.New("Deu erro"))
	service.Repository = repositoryMock

	_, err := service.GetBy(campaign.ID)

	assert.Equal(internalerrors.ErrInteral.Error(), err.Error())
}
