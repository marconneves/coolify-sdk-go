package coolify_sdk

import (
	"errors"
	"fmt"
	"time"

	client "github.com/marconneves/coolify-sdk-go/client"
)

type ProjectInstance struct {
	client *client.Client
}

type Project struct {
	ID           int64         `json:"id"`
	UUID         string        `json:"uuid"`
	Name         string        `json:"name"`
	Description  *string       `json:"description"`
	Environments []Environment `json:"environments"`
}

type Environment struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	ProjectId   int64     `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *ProjectInstance) List() (*[]Project, error) {
	body, err := t.client.HttpRequest("projects", "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &[]Project{})
}

func (t *ProjectInstance) Get(uuid string) (*Project, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.HttpRequest(fmt.Sprintf("projects/%v", uuid), "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &Project{})
}

type CreateProjectDTO struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type CreateProjectResponse struct {
	UUID string `json:"uuid"`
}

func (t *ProjectInstance) Create(server *CreateProjectDTO) (*string, error) {
	buf, err := client.EncodeRequest(server)
	if err != nil {
		return nil, err
	}

	body, err := t.client.HttpRequest("projects", "POST", *buf)
	if err != nil {
		return nil, err
	}

	response, err := client.DecodeResponse(body, &CreateProjectResponse{})
	if err != nil {
		return nil, err
	}

	return &response.UUID, nil
}

func (t *ProjectInstance) Delete(uuid string) error {
	if uuid == "" {
		return errors.New("uuid is required")
	}

	_, err := t.client.HttpRequest(fmt.Sprintf("projects/%v", uuid), "DELETE")
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

	buf, err := client.EncodeRequest(server)
	if err != nil {
		return err
	}

	_, err = t.client.HttpRequest(fmt.Sprintf("projects/%v", uuid), "PATCH", *buf)
	return err
}

type EnvironmentData struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProjectID   int64     `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *ProjectInstance) Environment(uuid string, environment string) (*EnvironmentData, error) {
	if uuid == "" {
		return nil, errors.New("uuid is required")
	}

	body, err := t.client.HttpRequest(fmt.Sprintf("projects/%v/%v/", uuid, environment), "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &EnvironmentData{})
}
