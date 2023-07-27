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

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	campanhaTeste = contract.NewCampaign{
		Name:    "campanha",
		Content: "Body",
		Emails:  []string{"email1@g.com", "email2@g.com", "email3@g.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(campanhaTeste)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_ValidateDomainError_Campaign(t *testing.T) {
	assert := assert.New(t)
	campanhaTeste.Name = ""

	_, err := service.Create(campanhaTeste)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())
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
