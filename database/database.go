package database

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/marconneves/coolify-sdk-go/client"
	"github.com/marconneves/coolify-sdk-go/server"
)

type DatabaseInstance struct {
	client *client.Client
}

func NewDatabaseInstance(client *client.Client) *DatabaseInstance {
	return &DatabaseInstance{client: client}
}

type Database struct {
	UUID          string  `json:"uuid"`
	Name          string  `json:"name"`
	Description   *string `json:"description"`
	PublicPort    int     `json:"public_port"`
	PortsMappings *string `json:"ports_mappings"`
	Image         string  `json:"image"`
	IsPublic      bool    `json:"is_public"`
	ExternalDbURL string  `json:"external_db_url"`
	InternalDbURL string  `json:"internal_db_url"`

	ServerStatus bool    `json:"server_status"`
	Status       string  `json:"status"`
	StartedAt    *string `json:"started_at"`

	LimitsCPUShares         int     `json:"limits_cpu_shares"`
	LimitsCpus              string  `json:"limits_cpus"`
	LimitsCpuset            *string `json:"limits_cpuset"`
	LimitsMemory            string  `json:"limits_memory"`
	LimitsMemoryReservation string  `json:"limits_memory_reservation"`
	LimitsMemorySwap        string  `json:"limits_memory_swap"`
	LimitsMemorySwappiness  int     `json:"limits_memory_swappiness"`

	DragonflyPassword *string `json:"dragonfly_password"`

	KeydbConf     *string `json:"keydb_conf"`
	KeydbPassword *string `json:"keydb_password"`

	ClickhouseAdminPassword *string `json:"clickhouse_admin_password"`
	ClickhouseAdminUser     *string `json:"clickhouse_admin_user"`

	MariadbConf         *string `json:"mariadb_conf"`
	MariadbDatabase     *string `json:"mariadb_database"`
	MariadbPassword     *string `json:"mariadb_password"`
	MariadbRootPassword *string `json:"mariadb_root_password"`
	MariadbUser         *string `json:"mariadb_user"`

	MongoConf               *string `json:"mongo_conf"`
	MongoInitdbInitDatabase *string `json:"mongo_initdb_init_database"`
	MongoInitdbRootPassword *string `json:"mongo_initdb_root_password"`
	MongoInitdbRootUsername *string `json:"mongo_initdb_root_username"`

	MysqlConf         *string `json:"mysql_conf"`
	MysqlDatabase     *string `json:"mysql_database,omitempty"`
	MysqlRootPassword *string `json:"mysql_root_password,omitempty"`
	MysqlUser         *string `json:"mysql_user,omitempty"`

	PostgresConf           *string `json:"postgres_conf"`
	PostgresDB             string  `json:"postgres_db"`
	PostgresHostAuthMethod *string `json:"postgres_host_auth_method"`
	PostgresInitdbArgs     *string `json:"postgres_initdb_args"`
	PostgresPassword       string  `json:"postgres_password"`
	PostgresUser           string  `json:"postgres_user"`

	RedisConf     *string `json:"redis_conf"`
	RedisPassword *string `json:"redis_password"`

	ConfigHash             string  `json:"config_hash"`
	CustomDockerRunOptions *string `json:"custom_docker_run_options"`
	DatabaseType           string  `json:"database_type"`

	Destination     Destination `json:"destination"`
	DestinationId   int         `json:"destination_id"`
	DestinationType string      `json:"destination_type"`

	EnvironmentID       int     `json:"environment_id"`
	InitScripts         *string `json:"init_scripts"`
	IsIncludeTimestamps bool    `json:"is_include_timestamps"`
	IsLogDrainEnabled   bool    `json:"is_log_drain_enabled"`

	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type Destination struct {
	CreatedAt string        `json:"created_at"`
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Network   string        `json:"network"`
	Server    server.Server `json:"server"`
	ServerID  int           `json:"server_id"`
	UpdatedAt string        `json:"updated_at"`
	UUID      string        `json:"uuid"`
}

func (d *DatabaseInstance) List() (*[]Database, error) {
	body, err := d.client.HttpRequest("databases", "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &[]Database{})
}

func (d *DatabaseInstance) Get(uuid string) (*Database, error) {
	if uuid == "" {
		return nil, errors.New("UUID is required")
	}

	body, err := d.client.HttpRequest(fmt.Sprintf("databases/%v", uuid), "GET")
	if err != nil {
		return nil, err
	}

	return client.DecodeResponse(body, &Database{})
}

func (d *DatabaseInstance) Start(uuid string) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}

	_, err := d.client.HttpRequest(fmt.Sprintf("databases/%v/start", uuid), "GET")

	return err
}

func (d *DatabaseInstance) Stop(uuid string) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}

	_, err := d.client.HttpRequest(fmt.Sprintf("databases/%v/stop", uuid), "GET")

	return err
}

func (d *DatabaseInstance) RestartDatabase(uuid string) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}

	_, err := d.client.HttpRequest(fmt.Sprintf("databases/%v/restart", uuid), "GET", bytes.Buffer{})

	return err
}

func (d *DatabaseInstance) Delete(uuid string) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}

	_, err := d.client.HttpRequest(fmt.Sprintf("databases/%v", uuid), "DELETE", bytes.Buffer{})
	return err
}

type UpdateDatabaseDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PublicPort  *int    `json:"public_port,omitempty"`
	Image       *string `json:"image,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`

	LimitsCPUShares         *int    `json:"limits_cpu_shares,omitempty"`
	LimitsCpus              *string `json:"limits_cpus,omitempty"`
	LimitsCpuset            *string `json:"limits_cpuset,omitempty"`
	LimitsMemory            *string `json:"limits_memory,omitempty"`
	LimitsMemoryReservation *string `json:"limits_memory_reservation,omitempty"`
	LimitsMemorySwap        *string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  *int    `json:"limits_memory_swappiness,omitempty"`

	DragonflyPassword *string `json:"dragonfly_password,omitempty"`

	KeydbConf     *string `json:"keydb_conf,omitempty"`
	KeydbPassword *string `json:"keydb_password,omitempty"`

	ClickhouseAdminPassword *string `json:"clickhouse_admin_password,omitempty"`
	ClickhouseAdminUser     *string `json:"clickhouse_admin_user,omitempty"`

	MariadbConf         *string `json:"mariadb_conf,omitempty"`
	MariadbDatabase     *string `json:"mariadb_database,omitempty"`
	MariadbPassword     *string `json:"mariadb_password,omitempty"`
	MariadbRootPassword *string `json:"mariadb_root_password,omitempty"`
	MariadbUser         *string `json:"mariadb_user,omitempty"`

	MongoConf               *string `json:"mongo_conf,omitempty"`
	MongoInitdbInitDatabase *string `json:"mongo_initdb_init_database,omitempty"`
	MongoInitdbRootPassword *string `json:"mongo_initdb_root_password,omitempty"`
	MongoInitdbRootUsername *string `json:"mongo_initdb_root_username,omitempty"`

	MysqlConf         *string `json:"mysql_conf,omitempty"`
	MysqlDatabase     *string `json:"mysql_database,omitempty"`
	MysqlRootPassword *string `json:"mysql_root_password,omitempty"`
	MysqlUser         *string `json:"mysql_user,omitempty"`

	PostgresConf           *string `json:"postgres_conf,omitempty"`
	PostgresDB             *string `json:"postgres_db,omitempty"`
	PostgresHostAuthMethod *string `json:"postgres_host_auth_method,omitempty"`
	PostgresInitdbArgs     *string `json:"postgres_initdb_args,omitempty"`
	PostgresPassword       *string `json:"postgres_password,omitempty"`
	PostgresUser           *string `json:"postgres_user,omitempty"`

	RedisConf     *string `json:"redis_conf,omitempty"`
	RedisPassword *string `json:"redis_password,omitempty"`
}

func (d *DatabaseInstance) Update(uuid string, data *UpdateDatabaseDTO) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}
	buf, err := client.EncodeRequest(data)
	if err != nil {
		return err
	}
	_, err = d.client.HttpRequest(fmt.Sprintf("databases/%v", uuid), "PATCH", *buf)
	return err
}
