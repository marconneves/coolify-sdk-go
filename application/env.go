package application

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/marconneves/coolify-sdk-go/client"
)

// EnvironmentVariable represents a Coolify environment variable.
type EnvironmentVariable struct {
	ID               int     `json:"id"`
	UUID             string  `json:"uuid"`
	ResourceableType string  `json:"resourceable_type"`
	ResourceableID   int     `json:"resourceable_id"`
	IsLiteral        bool    `json:"is_literal"`
	IsMultiline      bool    `json:"is_multiline"`
	IsPreview        bool    `json:"is_preview"`
	IsRuntime        bool    `json:"is_runtime"`
	IsBuildtime      bool    `json:"is_buildtime"`
	IsShared         bool    `json:"is_shared"`
	IsShownOnce      bool    `json:"is_shown_once"`
	Key              string  `json:"key"`
	Value            string  `json:"value"`
	RealValue        string  `json:"real_value"`
	Comment          *string `json:"comment"`
	Version          string  `json:"version"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// EnvDTO represents a single environment variable payload accepted by the API.
type EnvDTO struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	IsPreview   *bool  `json:"is_preview,omitempty"`
	IsLiteral   *bool  `json:"is_literal,omitempty"`
	IsMultiline *bool  `json:"is_multiline,omitempty"`
	IsShownOnce *bool  `json:"is_shown_once,omitempty"`
}

type bulkEnvsRequest struct {
	Data []EnvDTO `json:"data"`
}

type createEnvResponse struct {
	UUID string `json:"uuid"`
}

// ListEnvs lists all environment variables for an application.
func (a *ApplicationInstance) ListEnvs(ctx context.Context, applicationUUID string) (*[]EnvironmentVariable, error) {
	if applicationUUID == "" {
		return nil, errors.New("application UUID is required")
	}

	body, err := a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v/envs", applicationUUID), "GET")
	if err != nil {
		return nil, fmt.Errorf("failed to list envs for application %s: %w", applicationUUID, err)
	}

	res, err := client.DecodeResponse(body, &[]EnvironmentVariable{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode envs for application %s: %w", applicationUUID, err)
	}

	return res, nil
}

// CreateEnv creates a new environment variable for an application.
func (a *ApplicationInstance) CreateEnv(ctx context.Context, applicationUUID string, env *EnvDTO) (*string, error) {
	if applicationUUID == "" {
		return nil, errors.New("application UUID is required")
	}

	buf, err := client.EncodeRequest(env)
	if err != nil {
		return nil, fmt.Errorf("failed to encode env: %w", err)
	}

	body, err := a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v/envs", applicationUUID), "POST", *buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create env for application %s: %w", applicationUUID, err)
	}

	response, err := client.DecodeResponse(body, &createEnvResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode env create response: %w", err)
	}

	return &response.UUID, nil
}

// UpdateEnv updates an existing environment variable for an application,
// identified by its key.
func (a *ApplicationInstance) UpdateEnv(ctx context.Context, applicationUUID string, env *EnvDTO) (*EnvironmentVariable, error) {
	if applicationUUID == "" {
		return nil, errors.New("application UUID is required")
	}

	buf, err := client.EncodeRequest(env)
	if err != nil {
		return nil, fmt.Errorf("failed to encode env: %w", err)
	}

	body, err := a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v/envs", applicationUUID), "PATCH", *buf)
	if err != nil {
		return nil, fmt.Errorf("failed to update env for application %s: %w", applicationUUID, err)
	}

	res, err := client.DecodeResponse(body, &EnvironmentVariable{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode env update response: %w", err)
	}

	return res, nil
}

// UpdateEnvsBulk performs a bulk upsert of environment variables for an
// application. The API accepts the full list and creates/updates entries by
// matching on the `key` field.
func (a *ApplicationInstance) UpdateEnvsBulk(ctx context.Context, applicationUUID string, envs []EnvDTO) (*[]EnvironmentVariable, error) {
	if applicationUUID == "" {
		return nil, errors.New("application UUID is required")
	}

	if envs == nil {
		envs = []EnvDTO{}
	}

	buf, err := client.EncodeRequest(&bulkEnvsRequest{Data: envs})
	if err != nil {
		return nil, fmt.Errorf("failed to encode bulk envs: %w", err)
	}

	body, err := a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v/envs/bulk", applicationUUID), "PATCH", *buf)
	if err != nil {
		return nil, fmt.Errorf("failed to bulk update envs for application %s: %w", applicationUUID, err)
	}

	res, err := client.DecodeResponse(body, &[]EnvironmentVariable{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode bulk envs response: %w", err)
	}

	return res, nil
}

// DeleteEnv deletes an environment variable from an application by its UUID.
func (a *ApplicationInstance) DeleteEnv(ctx context.Context, applicationUUID, envUUID string) error {
	if applicationUUID == "" {
		return errors.New("application UUID is required")
	}
	if envUUID == "" {
		return errors.New("env UUID is required")
	}

	_, err := a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v/envs/%v", applicationUUID, envUUID), "DELETE", bytes.Buffer{})
	if err != nil {
		return fmt.Errorf("failed to delete env %s for application %s: %w", envUUID, applicationUUID, err)
	}

	return nil
}
