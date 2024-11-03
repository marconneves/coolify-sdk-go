package coolify_sdk

import (
	"errors"
	"fmt"
	"time"

	client "github.com/marconneves/coolify-sdk-go/client"
)

type PrivateKeyInstance struct {
	client *client.Client
}

type PrivateKey struct {
	ID           int       `json:"id"`
	UUID         string    `json:"uuid"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	PrivateKey   string    `json:"private_key"`
	IsGitRelated bool      `json:"is_git_related"`
	TeamID       int       `json:"team_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t *PrivateKeyInstance) List() (*[]PrivateKey, error) {
	body, err := t.client.HttpRequest("security/keys", "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &[]PrivateKey{})
}

func (t *PrivateKeyInstance) Get(uuid string) (*PrivateKey, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.HttpRequest(fmt.Sprintf("security/keys/%v", uuid), "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &PrivateKey{})
}

type CreatePrivateKeyDTO struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	PrivateKey  string  `json:"private_key"`
}

type CreatePrivateKeyResponse struct {
	UUID string `json:"uuid"`
}

func (t *PrivateKeyInstance) Create(server *CreatePrivateKeyDTO) (*string, error) {
	buf, err := client.EncodeRequest(server)
	if err != nil {
		return nil, err
	}

	body, err := t.client.HttpRequest("security/keys", "POST", *buf)
	if err != nil {
		return nil, err
	}

	response, err := client.DecodeResponse(body, &CreatePrivateKeyResponse{})
	if err != nil {
		return nil, err
	}

	return &response.UUID, nil
}

func (t *PrivateKeyInstance) Delete(uuid string) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	_, err := t.client.HttpRequest(fmt.Sprintf("security/keys/%v", uuid), "DELETE")
	if err != nil {
		return err
	}

	return nil
}

type UpdatePrivateKeyDTO struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	PrivateKey  *string `json:"private_key,omitempty"`
}

func (t *PrivateKeyInstance) Update(uuid string, privateKey *UpdatePrivateKeyDTO) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	buf, err := client.EncodeRequest(privateKey)
	if err != nil {
		return err
	}

	_, err = t.client.HttpRequest(fmt.Sprintf("security/keys/%v", uuid), "PATCH", *buf)
	return err
}
