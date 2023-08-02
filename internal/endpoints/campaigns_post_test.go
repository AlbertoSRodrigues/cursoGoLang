package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"projeto/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}
func (r *serviceMock) GetBy(id string) (*contract.ReadCampaign, error) {
	//args := r.Called(id)
	return nil, nil
}

func Test_CampaignPost_Save_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "Teste1",
		Content: "Body Teste",
		Emails:  []string{"teste1@g.com", "teste2@g.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content && len(request.Emails) == len(body.Emails) {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaignService: service}
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buffer)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CampaignPost_Error(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "Teste1",
		Content: "Body Teste",
		Emails:  []string{"teste1@g.com", "teste2@g.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buffer)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
