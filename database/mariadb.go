package database

import (
	"context"
	"fmt"

	"github.com/marconneves/coolify-sdk-go/client"
)

// CreateDatabaseMariaDBDTO represents the data required to create a MariaDB database.
type CreateDatabaseMariaDBDTO struct {
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

	MariadbConf         *string `json:"mariadb_conf,omitempty"`
	MariadbRootPassword *string `json:"mariadb_root_password,omitempty"`
	MariadbUser         *string `json:"mariadb_user,omitempty"`
	MariadbPassword     *string `json:"mariadb_password,omitempty"`
	MariadbDatabase     *string `json:"mariadb_database,omitempty"`

	LimitsMemory            *string `json:"limits_memory,omitempty"`
	LimitsMemorySwap        *string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  *int    `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation *string `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs              *string `json:"limits_cpus,omitempty"`
	LimitsCPUSet            *string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares         *int    `json:"limits_cpu_shares,omitempty"`
}

// CreateDatabaseMariaDBResponse represents the response when creating a MariaDB database.
type CreateDatabaseMariaDBResponse struct {
	UUID string `json:"uuid"`
}

// CreateMariaDB creates a new MariaDB database instance.
func (d *DatabaseInstance) CreateMariaDB(ctx context.Context, data *CreateDatabaseMariaDBDTO) (*string, error) {
	buf, err := client.EncodeRequest(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	body, err := d.client.HttpRequestWithContext(ctx, "databases/mariadb", "POST", *buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create MariaDB database: %w", err)
	}

	response, err := client.DecodeResponse(body, &CreateDatabaseMariaDBResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response.UUID, nil
}
