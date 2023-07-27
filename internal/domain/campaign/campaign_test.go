package campaign

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Bodys"
	contacts = []string{"email1@g.com", "email2@g.com", "email3@g.com", "asdasd"}
	fake     = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t) //arrange

	campaign, _ := NewCampaign(name, content, contacts) //act

	assert.Equal(campaign.Name, name) //assert
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_NotEmpty(t *testing.T) {
	assert := assert.New(t) //arrange

	campaign, _ := NewCampaign(name, content, contacts) //act

	assert.NotEmpty(campaign.ID) //Assert
}

func Test_NewCampaign_MustValidate_MinChar_Name(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name requires 5 minimum characters", err.Error())
}
func Test_NewCampaign_MustValidate_MaxChar_Name(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(25), content, contacts)

	assert.Equal("name requires 24 maximum characters", err.Error())
}
func Test_NewCampaign_MustValidate_MinChar_Content(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content requires 5 minimum characters", err.Error())
}
func Test_NewCampaign_MustValidate_MaxChar_Content(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content requires 1024 maximum characters", err.Error())
}

func Test_NewCampaign_MustValidate_MinChar_Contacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("contacts requires 1 minimum characters", err.Error())
}
func Test_NewCampaign_MustValidate_ValidEmail_Contacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"emailinvalido"})

	assert.Equal("email is an invalid email", err.Error())
}
