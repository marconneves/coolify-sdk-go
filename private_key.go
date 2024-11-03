package coolify_sdk

import (
	"bytes"
	"errors"
	"fmt"
)

type PrivateKeyInstance struct {
	client *Client
}

type PrivateKey struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PrivateKey   string `json:"private_key"`
	IsGitRelated bool   `json:"is_git_related"`
	TeamID       int    `json:"team_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (t *PrivateKeyInstance) List() (*[]PrivateKey, error) {
	body, err := t.client.httpRequest("security/keys", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]PrivateKey{})
}

func (t *PrivateKeyInstance) Get(uuid string) (*PrivateKey, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.httpRequest(fmt.Sprintf("security/keys/%v", uuid), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &PrivateKey{})
}

type CreatePrivateKeyDTO struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	PrivateKey  string  `json:"private_key"`
}

type CreatePrivateKeyResponse struct {
	UUID string `json:"uuid"`
}

func (t *PrivateKeyInstance) Create(server *CreatePrivateKeyDTO) (*string, error) {
	buf, err := encodeRequest(server)
	if err != nil {
		return nil, err
	}

	body, err := t.client.httpRequest("security/keys", "POST", *buf)
	if err != nil {
		return nil, err
	}

	response, err := decodeResponse(body, &CreatePrivateKeyResponse{})
	if err != nil {
		return nil, err
	}

	return &response.UUID, nil
}

func (t *PrivateKeyInstance) Delete(uuid string) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	_, err := t.client.httpRequest(fmt.Sprintf("security/keys/%v", uuid), "DELETE")
	if err != nil {
		return err
	}

	return nil
}

type UpdatePrivateKeyDTO struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	PrivateKey  string  `json:"private_key"`
}

func (t *PrivateKeyInstance) Update(uuid string, privateKey *UpdatePrivateKeyDTO) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	buf, err := encodeRequest(privateKey)
	if err != nil {
		return err
	}

	_, err = t.client.httpRequest(fmt.Sprintf("security/keys/%v", uuid), "PATCH", *buf)
	return err
}
