package application

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/marconneves/coolify-sdk-go/client"
)

// ApplicationInstance provides methods to interact with application resources.
type ApplicationInstance struct {
	client *client.Client
}

// NewApplicationInstance creates a new ApplicationInstance.
func NewApplicationInstance(client *client.Client) *ApplicationInstance {
	return &ApplicationInstance{client: client}
}

// Application represents a Coolify application entity.
type Application struct {
	ID                                  int     `json:"id"`
	UUID                                string  `json:"uuid"`
	Name                                string  `json:"name"`
	Description                         *string `json:"description"`
	RepositoryProjectID                 *int    `json:"repository_project_id"`
	FQDN                                *string `json:"fqdn"`
	ConfigHash                          string  `json:"config_hash"`
	GitRepository                       string  `json:"git_repository"`
	GitBranch                           string  `json:"git_branch"`
	GitCommitSHA                        string  `json:"git_commit_sha"`
	GitFullURL                          *string `json:"git_full_url"`
	DockerRegistryImageName             *string `json:"docker_registry_image_name"`
	DockerRegistryImageTag              *string `json:"docker_registry_image_tag"`
	BuildPack                           string  `json:"build_pack"`
	StaticImage                         string  `json:"static_image"`
	InstallCommand                      string  `json:"install_command"`
	BuildCommand                        string  `json:"build_command"`
	StartCommand                        string  `json:"start_command"`
	PortsExposes                        string  `json:"ports_exposes"`
	PortsMappings                       *string `json:"ports_mappings"`
	CustomNetworkAliases                *string `json:"custom_network_aliases"`
	BaseDirectory                       string  `json:"base_directory"`
	PublishDirectory                    string  `json:"publish_directory"`
	HealthCheckEnabled                  bool    `json:"health_check_enabled"`
	HealthCheckPath                     string  `json:"health_check_path"`
	HealthCheckPort                     *string `json:"health_check_port"`
	HealthCheckHost                     *string `json:"health_check_host"`
	HealthCheckMethod                   string  `json:"health_check_method"`
	HealthCheckReturnCode               int     `json:"health_check_return_code"`
	HealthCheckScheme                   string  `json:"health_check_scheme"`
	HealthCheckResponseText             *string `json:"health_check_response_text"`
	HealthCheckInterval                 int     `json:"health_check_interval"`
	HealthCheckTimeout                  int     `json:"health_check_timeout"`
	HealthCheckRetries                  int     `json:"health_check_retries"`
	HealthCheckStartPeriod              int     `json:"health_check_start_period"`
	HealthCheckType                     string  `json:"health_check_type"`
	HealthCheckCommand                  *string `json:"health_check_command"`
	LimitsMemory                        string  `json:"limits_memory"`
	LimitsMemorySwap                    string  `json:"limits_memory_swap"`
	LimitsMemorySwappiness              int     `json:"limits_memory_swappiness"`
	LimitsMemoryReservation             string  `json:"limits_memory_reservation"`
	LimitsCpus                          string  `json:"limits_cpus"`
	LimitsCpuset                        *string `json:"limits_cpuset"`
	LimitsCPUShares                     int     `json:"limits_cpu_shares"`
	Status                              string  `json:"status"`
	PreviewURLTemplate                  string  `json:"preview_url_template"`
	DestinationType                     string  `json:"destination_type"`
	DestinationID                       int     `json:"destination_id"`
	SourceID                            *int    `json:"source_id"`
	PrivateKeyID                        *int    `json:"private_key_id"`
	EnvironmentID                       int     `json:"environment_id"`
	Dockerfile                          *string `json:"dockerfile"`
	DockerfileLocation                  string  `json:"dockerfile_location"`
	CustomLabels                        *string `json:"custom_labels"`
	DockerfileTargetBuild               *string `json:"dockerfile_target_build"`
	ManualWebhookSecretGithub           *string `json:"manual_webhook_secret_github"`
	ManualWebhookSecretGitlab           *string `json:"manual_webhook_secret_gitlab"`
	ManualWebhookSecretBitbucket        *string `json:"manual_webhook_secret_bitbucket"`
	ManualWebhookSecretGitea            *string `json:"manual_webhook_secret_gitea"`
	DockerComposeLocation               string  `json:"docker_compose_location"`
	DockerCompose                       *string `json:"docker_compose"`
	DockerComposeRaw                    *string `json:"docker_compose_raw"`
	DockerComposeDomains                *string `json:"docker_compose_domains"`
	DockerComposeCustomStartCommand     *string `json:"docker_compose_custom_start_command"`
	DockerComposeCustomBuildCommand     *string `json:"docker_compose_custom_build_command"`
	SwarmReplicas                       *int    `json:"swarm_replicas"`
	SwarmPlacementConstraints           *string `json:"swarm_placement_constraints"`
	CustomDockerRunOptions              *string `json:"custom_docker_run_options"`
	PostDeploymentCommand               *string `json:"post_deployment_command"`
	PostDeploymentCommandContainer      *string `json:"post_deployment_command_container"`
	PreDeploymentCommand                *string `json:"pre_deployment_command"`
	PreDeploymentCommandContainer       *string `json:"pre_deployment_command_container"`
	WatchPaths                          *string `json:"watch_paths"`
	CustomHealthcheckFound              bool    `json:"custom_healthcheck_found"`
	Redirect                            *string `json:"redirect"`
	CreatedAt                           string  `json:"created_at"`
	UpdatedAt                           string  `json:"updated_at"`
	DeletedAt                           *string `json:"deleted_at"`
	ComposeParsingVersion               string  `json:"compose_parsing_version"`
	CustomNginxConfiguration            *string `json:"custom_nginx_configuration"`
	IsHTTPBasicAuthEnabled              bool    `json:"is_http_basic_auth_enabled"`
	HTTPBasicAuthUsername               *string `json:"http_basic_auth_username"`
	HTTPBasicAuthPassword               *string `json:"http_basic_auth_password"`
}

// List retrieves all applications.
func (a *ApplicationInstance) List(ctx context.Context) (*[]Application, error) {
	body, err := a.client.HttpRequestWithContext(ctx, "applications", "GET")
	if err != nil {
		return nil, fmt.Errorf("failed to list applications: %w", err)
	}

	res, err := client.DecodeResponse(body, &[]Application{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode applications list: %w", err)
	}

	return res, nil
}

// Get retrieves a specific application by UUID.
func (a *ApplicationInstance) Get(ctx context.Context, uuid string) (*Application, error) {
	if uuid == "" {
		return nil, errors.New("UUID is required")
	}

	body, err := a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v", uuid), "GET")
	if err != nil {
		return nil, fmt.Errorf("failed to get application %s: %w", uuid, err)
	}

	res, err := client.DecodeResponse(body, &Application{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode application %s: %w", uuid, err)
	}

	return res, nil
}

// Delete removes an application.
type DeleteApplicationOptions struct {
	DeleteConfigurations    *bool
	DeleteVolumes           *bool
	DockerCleanup           *bool
	DeleteConnectedNetworks *bool
}

func (a *ApplicationInstance) Delete(ctx context.Context, uuid string, opts *DeleteApplicationOptions) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}

	path := fmt.Sprintf("applications/%v", uuid)
	if opts != nil {
		q := url.Values{}
		if opts.DeleteConfigurations != nil {
			q.Set("delete_configurations", fmt.Sprintf("%t", *opts.DeleteConfigurations))
		}
		if opts.DeleteVolumes != nil {
			q.Set("delete_volumes", fmt.Sprintf("%t", *opts.DeleteVolumes))
		}
		if opts.DockerCleanup != nil {
			q.Set("docker_cleanup", fmt.Sprintf("%t", *opts.DockerCleanup))
		}
		if opts.DeleteConnectedNetworks != nil {
			q.Set("delete_connected_networks", fmt.Sprintf("%t", *opts.DeleteConnectedNetworks))
		}
		if encoded := q.Encode(); encoded != "" {
			path = path + "?" + encoded
		}
	}

	_, err := a.client.HttpRequestWithContext(ctx, path, "DELETE", bytes.Buffer{})
	if err != nil {
		return fmt.Errorf("failed to delete application %s: %w", uuid, err)
	}

	return nil
}

// StartOptions controls Start behavior.
type StartOptions struct {
	Force         *bool
	InstantDeploy *bool
}

// StartResponse is the response from starting an application.
type StartResponse struct {
	Message        string `json:"message"`
	DeploymentUUID string `json:"deployment_uuid"`
}

// Start starts an application.
func (a *ApplicationInstance) Start(ctx context.Context, uuid string, opts *StartOptions) (*StartResponse, error) {
	if uuid == "" {
		return nil, errors.New("UUID is required")
	}

	path := fmt.Sprintf("applications/%v/start", uuid)
	if opts != nil {
		q := url.Values{}
		if opts.Force != nil {
			q.Set("force", fmt.Sprintf("%t", *opts.Force))
		}
		if opts.InstantDeploy != nil {
			q.Set("instant_deploy", fmt.Sprintf("%t", *opts.InstantDeploy))
		}
		if encoded := q.Encode(); encoded != "" {
			path = path + "?" + encoded
		}
	}

	body, err := a.client.HttpRequestWithContext(ctx, path, "GET")
	if err != nil {
		return nil, fmt.Errorf("failed to start application %s: %w", uuid, err)
	}

	res, err := client.DecodeResponse(body, &StartResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode start response for %s: %w", uuid, err)
	}

	return res, nil
}

// StopOptions controls Stop behavior.
type StopOptions struct {
	DockerCleanup *bool
}

// Stop stops an application.
func (a *ApplicationInstance) Stop(ctx context.Context, uuid string, opts *StopOptions) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}

	path := fmt.Sprintf("applications/%v/stop", uuid)
	if opts != nil {
		q := url.Values{}
		if opts.DockerCleanup != nil {
			q.Set("docker_cleanup", fmt.Sprintf("%t", *opts.DockerCleanup))
		}
		if encoded := q.Encode(); encoded != "" {
			path = path + "?" + encoded
		}
	}

	_, err := a.client.HttpRequestWithContext(ctx, path, "GET")
	if err != nil {
		return fmt.Errorf("failed to stop application %s: %w", uuid, err)
	}

	return nil
}

// RestartResponse is the response from restarting an application.
type RestartResponse struct {
	Message        string `json:"message"`
	DeploymentUUID string `json:"deployment_uuid"`
}

// Restart restarts an application.
func (a *ApplicationInstance) Restart(ctx context.Context, uuid string) (*RestartResponse, error) {
	if uuid == "" {
		return nil, errors.New("UUID is required")
	}

	body, err := a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v/restart", uuid), "GET")
	if err != nil {
		return nil, fmt.Errorf("failed to restart application %s: %w", uuid, err)
	}

	res, err := client.DecodeResponse(body, &RestartResponse{})
	if err != nil {
		return nil, fmt.Errorf("failed to decode restart response for %s: %w", uuid, err)
	}

	return res, nil
}

// UpdateApplicationDTO represents the data required to update an application.
type UpdateApplicationDTO struct {
	ProjectUUID                     *string `json:"project_uuid,omitempty"`
	ServerUUID                      *string `json:"server_uuid,omitempty"`
	EnvironmentName                 *string `json:"environment_name,omitempty"`
	DestinationUUID                 *string `json:"destination_uuid,omitempty"`
	Name                            *string `json:"name,omitempty"`
	Description                     *string `json:"description,omitempty"`
	Domains                         *string `json:"domains,omitempty"`
	DockerRegistryImageName         *string `json:"docker_registry_image_name,omitempty"`
	DockerRegistryImageTag          *string `json:"docker_registry_image_tag,omitempty"`
	PortsExposes                    *string `json:"ports_exposes,omitempty"`
	PortsMappings                   *string `json:"ports_mappings,omitempty"`

	HealthCheckEnabled      *bool   `json:"health_check_enabled,omitempty"`
	HealthCheckPath         *string `json:"health_check_path,omitempty"`
	HealthCheckPort         *string `json:"health_check_port,omitempty"`
	HealthCheckHost         *string `json:"health_check_host,omitempty"`
	HealthCheckMethod       *string `json:"health_check_method,omitempty"`
	HealthCheckReturnCode   *int    `json:"health_check_return_code,omitempty"`
	HealthCheckScheme       *string `json:"health_check_scheme,omitempty"`
	HealthCheckResponseText *string `json:"health_check_response_text,omitempty"`
	HealthCheckInterval     *int    `json:"health_check_interval,omitempty"`
	HealthCheckTimeout      *int    `json:"health_check_timeout,omitempty"`
	HealthCheckRetries      *int    `json:"health_check_retries,omitempty"`
	HealthCheckStartPeriod  *int    `json:"health_check_start_period,omitempty"`

	LimitsMemory            *string `json:"limits_memory,omitempty"`
	LimitsMemorySwap        *string `json:"limits_memory_swap,omitempty"`
	LimitsMemorySwappiness  *int    `json:"limits_memory_swappiness,omitempty"`
	LimitsMemoryReservation *string `json:"limits_memory_reservation,omitempty"`
	LimitsCpus              *string `json:"limits_cpus,omitempty"`
	LimitsCpuset            *string `json:"limits_cpuset,omitempty"`
	LimitsCPUShares         *int    `json:"limits_cpu_shares,omitempty"`

	CustomLabels                   *string `json:"custom_labels,omitempty"`
	CustomDockerRunOptions         *string `json:"custom_docker_run_options,omitempty"`
	PostDeploymentCommand          *string `json:"post_deployment_command,omitempty"`
	PostDeploymentCommandContainer *string `json:"post_deployment_command_container,omitempty"`
	PreDeploymentCommand           *string `json:"pre_deployment_command,omitempty"`
	PreDeploymentCommandContainer  *string `json:"pre_deployment_command_container,omitempty"`

	Redirect                  *string `json:"redirect,omitempty"`
	InstantDeploy             *bool   `json:"instant_deploy,omitempty"`
	IsForceHTTPSEnabled       *bool   `json:"is_force_https_enabled,omitempty"`
	IsHTTPBasicAuthEnabled    *bool   `json:"is_http_basic_auth_enabled,omitempty"`
	HTTPBasicAuthUsername     *string `json:"http_basic_auth_username,omitempty"`
	HTTPBasicAuthPassword     *string `json:"http_basic_auth_password,omitempty"`
	ConnectToDockerNetwork    *bool   `json:"connect_to_docker_network,omitempty"`
}

// Update updates an application.
func (a *ApplicationInstance) Update(ctx context.Context, uuid string, data *UpdateApplicationDTO) error {
	if uuid == "" {
		return errors.New("UUID is required")
	}

	buf, err := client.EncodeRequest(data)
	if err != nil {
		return fmt.Errorf("failed to encode update request: %w", err)
	}

	_, err = a.client.HttpRequestWithContext(ctx, fmt.Sprintf("applications/%v", uuid), "PATCH", *buf)
	if err != nil {
		return fmt.Errorf("failed to update application %s: %w", uuid, err)
	}

	return nil
}
