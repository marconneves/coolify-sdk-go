package database

import "github.com/marconneves/coolify-sdk-go/client"

type CreateDatabaseDTO struct {
	ServerUUID      string  `json:"server_uuid"`
	ProjectUUID     string  `json:"project_uuid"`
	Environment     string  `json:"environment_name"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	Image           *string `json:"image,omitempty"`
	IsPublic        *bool   `json:"is_public,omitempty"`
	PublicPort      *int    `json:"public_port,omitempty"`
	InstantDeploy   *bool   `json:"instant_deploy,omitempty"`
	DestinationUUID *string `json:"destination_uuid,omitempty"`

	PostgresUser           *string `json:"postgres_user,omitempty"`
	PostgresPassword       *string `json:"postgres_password,omitempty"`
	PostgresDB             *string `json:"postgres_db,omitempty"`
	PostgresInitdbArgs     *string `json:"postgres_initdb_args,omitempty"`
	PostgresHostAuthMethod *string `json:"postgres_host_auth_method,omitempty"`
	PostgresConf           *string `json:"postgres_conf,omitempty"`

	LimitsMemory            *string `json:"limits_memory,omitempty"`
	LimitsMemorySwap        *string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  *int    `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation *string `json:"limits_memory_reservation,omitempty"`
	LimitsCPUs              *string `json:"limits_cpus,omitempty"`
	LimitsCPUSet            *string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares         *int    `json:"limits_cpu_shares,omitempty"`
}

type CreateDatabaseResponse struct {
	UUID string `json:"uuid"`
}

func (d *DatabaseInstance) CreatePostgreSQL(data *CreateDatabaseDTO) (*string, error) {
	buf, err := client.EncodeRequest(data)
	if err != nil {
		return nil, err
	}

	body, err := d.client.HttpRequest("databases/postgresql", "POST", *buf)
	if err != nil {
		return nil, err
	}

	response, err := client.DecodeResponse(body, &CreateDatabaseResponse{})
	if err != nil {
		return nil, err
	}

	return &response.UUID, nil
}
