package coolify_sdk

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

type ProjectInstance struct {
	client *Client
}

type Project struct {
	Description  *string       `json:"description,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
	ID           *int64        `json:"id,omitempty"`
	Name         *string       `json:"name,omitempty"`
	UUID         *string       `json:"uuid,omitempty"`
}

type Environment struct {
	CreatedAt   *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	ProjectID   *int64  `json:"project_id,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

func (t *ProjectInstance) List() (*[]Project, error) {
	body, err := t.client.httpRequest("projects", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]Project{})
}

func (t *ProjectInstance) Get(uuid string) (*Project, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.httpRequest(fmt.Sprintf("projects/%v", uuid), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &Project{})
}

type CreateProjectDTO struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type CreateProjectResponse struct {
	UUID string `json:"uuid"`
}

func (t *ProjectInstance) Create(server *CreateProjectDTO) (*string, error) {
	buf, err := encodeRequest(server)
	if err != nil {
		return nil, err
	}

	body, err := t.client.httpRequest("projects", "POST", *buf)
	if err != nil {
		return nil, err
	}

	response, err := decodeResponse(body, &CreateProjectResponse{})
	if err != nil {
		return nil, err
	}

	return &response.UUID, nil
}

func (t *ProjectInstance) Delete(uuid string) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	_, err := t.client.httpRequest(fmt.Sprintf("projects/%v", uuid), "DELETE")
	if err != nil {
		return err
	}

	return nil
}

type UpdateProjectDTO struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (t *ProjectInstance) Update(uuid string, server *UpdateProjectDTO) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	buf, err := encodeRequest(server)
	if err != nil {
		return err
	}

	_, err = t.client.httpRequest(fmt.Sprintf("projects/%v", uuid), "PATCH", *buf)
	return err
}

type EnvironmentData struct {
	Id        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `json:"status"`
}

func (t *ProjectInstance) Resources(uuid string, environment string) (*[]EnvironmentData, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.httpRequest(fmt.Sprintf("projects/%v/%v/", uuid, environment), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return decodeResponse(body, &[]EnvironmentData{})
}
