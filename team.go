package coolify_sdk

import (
	"fmt"
	"time"

	client "github.com/marconneves/coolify-sdk-go/client"
)

type TeamInstance struct {
	client *client.Client
}

type Team struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *TeamInstance) List() (*[]Team, error) {
	body, err := t.client.HttpRequest("teams", "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &[]Team{})
}

func (t *TeamInstance) Get(id int) (*Team, error) {
	body, err := t.client.HttpRequest(fmt.Sprintf("teams/%v", id), "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &Team{})
}

type Member struct {
	Id                   int     `json:"id"`
	Name                 string  `json:"name"`
	Email                string  `json:"email"`
	EmailVerifiedAt      *string `json:"email_verified_at"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	TwoFactorConfirmedAt *string `json:"two_factor_confirmed_at"`
	ForcePasswordReset   bool    `json:"force_password_reset"`
	MarketingEmails      bool    `json:"marketing_emails"`
}

func (t *TeamInstance) Members(id int) (*[]Member, error) {
	body, err := t.client.HttpRequest(fmt.Sprintf("teams/%v/members", id), "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &[]Member{})
}
