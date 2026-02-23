package database

import (
	"context"
	"fmt"

	"github.com/marconneves/coolify-sdk-go/client"
)

// CreateDatabaseRedisDTO represents the data required to create a Redis database.
type CreateDatabaseRedisDTO struct {
	ServerUUID      string  `json:"server_uuid"`
	ProjectUUID     string  `json:"project_uuid"`
	EnvironmentName string  `json:"environment_name"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	Image           *string `json:"image,omitempty"`
	IsPublic        *bool   `json:"is_public,omitempty"`
	PublicPort      *int    `json:"public_port,omitempty"`
	InstantDeploy   *bool   `json:"instant_deploy,omitempty"`
	DestinationUUID *string `json:"destination_uuid,omitempty"`

	RedisPassword *string `json:"redis_password,omitempty"`
	RedisConf     *string `json:"redis_conf,omitempty"`

	LimitsMemory            *string `json:"limits_memory,omitempty"`
	LimitsMemorySwap        *string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  *int    `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation *string `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs              *string `json:"limits_cpus,omitempty"`
	LimitsCPUSet            *string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares         *int    `json:"limits_cpu_shares,omitempty"`
}

// CreateDatabaseRedisResponse represents the response when creating a Redis database.
type CreateDatabaseRedisResponse struct {
	UUID string `json:"uuid"`
}

// CreateRedis creates a new Redis database instance.
func (d *DatabaseInstance) CreateRedis(ctx context.Context, data *CreateDatabaseRedisDTO) (*string, error) {
	buf, err := client.EncodeRequest(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	body, err := d.client.HttpRequestWithContext(ctx, "databases/redis", "POST", *buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create Redis database: %w", err)
	}

	response, err := client.DecodeResponse(body, &CreateDatabaseRedisResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response.UUID, nil
}
