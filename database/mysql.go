package database

import (
	"context"
	"fmt"

	"github.com/marconneves/coolify-sdk-go/client"
)

// CreateDatabaseMySQLDTO represents the data required to create a MySQL database.
type CreateDatabaseMySQLDTO struct {
	ServerUUID      string  `json:"server_uuid"`
	ProjectUUID     string  `json:"project_uuid"`
	EnvironmentName string  `json:"environment_name"`
	EnvironmentUUID *string `json:"environment_uuid,omitempty"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	Image           *string `json:"image,omitempty"`
	IsPublic        *bool   `json:"is_public,omitempty"`
	PublicPort      *int    `json:"public_port,omitempty"`
	InstantDeploy   *bool   `json:"instant_deploy,omitempty"`
	DestinationUUID *string `json:"destination_uuid,omitempty"`

	MysqlRootPassword *string `json:"mysql_root_password,omitempty"`
	MysqlPassword     *string `json:"mysql_password,omitempty"`
	MysqlUser         *string `json:"mysql_user,omitempty"`
	MysqlDatabase     *string `json:"mysql_database,omitempty"`
	MysqlConf         *string `json:"mysql_conf,omitempty"`

	LimitsMemory            *string `json:"limits_memory,omitempty"`
	LimitsMemorySwap        *string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  *int    `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation *string `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs              *string `json:"limits_cpus,omitempty"`
	LimitsCPUSet            *string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares         *int    `json:"limits_cpu_shares,omitempty"`
}

// CreateDatabaseMySQLResponse represents the response when creating a MySQL database.
type CreateDatabaseMySQLResponse struct {
	UUID string `json:"uuid"`
}

// CreateMySQL creates a new MySQL database instance.
func (d *DatabaseInstance) CreateMySQL(ctx context.Context, data *CreateDatabaseMySQLDTO) (*string, error) {
	buf, err := client.EncodeRequest(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	body, err := d.client.HttpRequestWithContext(ctx, "databases/mysql", "POST", *buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create MySQL database: %w", err)
	}

	response, err := client.DecodeResponse(body, &CreateDatabaseMySQLResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response.UUID, nil
}
