package coolify_sdk_test

import (
	"context"
	"testing"

	sdk "github.com/marconneves/coolify-sdk-go"
	"github.com/marconneves/coolify-sdk-go/application"
)

// Replace these placeholders with values from your Coolify instance before
// running. The "ValidRequest" cases require resources that actually exist.
const (
	applicationProjectUUID     = "REPLACE_WITH_PROJECT_UUID"
	applicationServerUUID      = "REPLACE_WITH_SERVER_UUID"
	applicationEnvironmentName = "production"

	// An existing application UUID used by Get/Start/Stop/Restart/Update/envs tests.
	// CreateDockerImage will populate this in a real run, but the test suite
	// runs each case independently, so set it here too.
	existingApplicationUUID = "REPLACE_WITH_APPLICATION_UUID"

	// An env UUID belonging to existingApplicationUUID, used by DeleteEnv.
	existingApplicationEnvUUID = "REPLACE_WITH_ENV_UUID"
)

func TestListApplications(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		Error  bool
	}{
		"ValidRequest": {
			Host:   host,
			ApiKey: apiKey,
			Error:  false,
		},
		"WithoutToken": {
			Host:   host,
			ApiKey: "",
			Error:  true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(testComponent.Host, testComponent.ApiKey)

			_, err := client.Application.List(context.Background())

			if err != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error: %v", testComponent.Host, testComponent.ApiKey, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s) did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestGetApplication(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidUUID": {
			UUID:  existingApplicationUUID,
			Error: false,
		},
		"EmptyUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.Get(context.Background(), testComponent.UUID)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func TestCreateDockerImageApplication(t *testing.T) {
	cases := map[string]struct {
		DTO   *application.CreateDockerImageDTO
		Error bool
	}{
		"ValidRequest": {
			DTO: &application.CreateDockerImageDTO{
				ProjectUUID:             applicationProjectUUID,
				ServerUUID:              applicationServerUUID,
				EnvironmentName:         applicationEnvironmentName,
				DockerRegistryImageName: "nginx",
				DockerRegistryImageTag:  stringPtr("alpine"),
				PortsExposes:            "80",
				Name:                    stringPtr("sdk-test-app"),
			},
			Error: false,
		},
		"MissingProject": {
			DTO: &application.CreateDockerImageDTO{
				ServerUUID:              applicationServerUUID,
				EnvironmentName:         applicationEnvironmentName,
				DockerRegistryImageName: "nginx",
				PortsExposes:            "80",
			},
			Error: true,
		},
		"MissingImage": {
			DTO: &application.CreateDockerImageDTO{
				ProjectUUID:     applicationProjectUUID,
				ServerUUID:      applicationServerUUID,
				EnvironmentName: applicationEnvironmentName,
				PortsExposes:    "80",
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.CreateDockerImage(context.Background(), testComponent.DTO)

			if err != nil && !testComponent.Error {
				t.Errorf("DTO (%+v) produced an unexpected error: %v", testComponent.DTO, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("DTO (%+v) did not error", testComponent.DTO)
			}
		})
	}
}

func TestUpdateApplication(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		DTO   *application.UpdateApplicationDTO
		Error bool
	}{
		"ValidUpdate": {
			UUID: existingApplicationUUID,
			DTO: &application.UpdateApplicationDTO{
				Description: stringPtr("Updated by SDK test"),
			},
			Error: false,
		},
		"EmptyUUID": {
			UUID: "",
			DTO: &application.UpdateApplicationDTO{
				Description: stringPtr("noop"),
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			err := client.Application.Update(context.Background(), testComponent.UUID, testComponent.DTO)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func TestStartApplication(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Opts  *application.StartOptions
		Error bool
	}{
		"ValidStart": {
			UUID:  existingApplicationUUID,
			Opts:  nil,
			Error: false,
		},
		"InstantDeploy": {
			UUID: existingApplicationUUID,
			Opts: &application.StartOptions{
				InstantDeploy: boolPtr(true),
			},
			Error: false,
		},
		"EmptyUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.Start(context.Background(), testComponent.UUID, testComponent.Opts)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func TestStopApplication(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidStop": {
			UUID:  existingApplicationUUID,
			Error: false,
		},
		"EmptyUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			err := client.Application.Stop(context.Background(), testComponent.UUID, nil)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func TestRestartApplication(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidRestart": {
			UUID:  existingApplicationUUID,
			Error: false,
		},
		"EmptyUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.Restart(context.Background(), testComponent.UUID)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func TestListApplicationEnvs(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidUUID": {
			UUID:  existingApplicationUUID,
			Error: false,
		},
		"EmptyUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.ListEnvs(context.Background(), testComponent.UUID)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func TestCreateApplicationEnv(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Env   *application.EnvDTO
		Error bool
	}{
		"ValidEnv": {
			UUID: existingApplicationUUID,
			Env: &application.EnvDTO{
				Key:   "SDK_TEST_KEY",
				Value: "sdk-test-value",
			},
			Error: false,
		},
		"EmptyUUID": {
			UUID: "",
			Env: &application.EnvDTO{
				Key:   "SDK_TEST_KEY",
				Value: "sdk-test-value",
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.CreateEnv(context.Background(), testComponent.UUID, testComponent.Env)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s), Env (%+v) produced an unexpected error: %v", testComponent.UUID, testComponent.Env, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s), Env (%+v) did not error", testComponent.UUID, testComponent.Env)
			}
		})
	}
}

func TestUpdateApplicationEnv(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Env   *application.EnvDTO
		Error bool
	}{
		"ValidUpdate": {
			UUID: existingApplicationUUID,
			Env: &application.EnvDTO{
				Key:   "SDK_TEST_KEY",
				Value: "sdk-test-value-updated",
			},
			Error: false,
		},
		"EmptyUUID": {
			UUID: "",
			Env: &application.EnvDTO{
				Key:   "SDK_TEST_KEY",
				Value: "sdk-test-value-updated",
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.UpdateEnv(context.Background(), testComponent.UUID, testComponent.Env)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s), Env (%+v) produced an unexpected error: %v", testComponent.UUID, testComponent.Env, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s), Env (%+v) did not error", testComponent.UUID, testComponent.Env)
			}
		})
	}
}

func TestUpdateApplicationEnvsBulk(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Envs  []application.EnvDTO
		Error bool
	}{
		"ValidBulk": {
			UUID: existingApplicationUUID,
			Envs: []application.EnvDTO{
				{Key: "SDK_BULK_A", Value: "value-a"},
				{Key: "SDK_BULK_B", Value: "value-b"},
			},
			Error: false,
		},
		"EmptyList": {
			UUID:  existingApplicationUUID,
			Envs:  nil,
			Error: false,
		},
		"EmptyUUID": {
			UUID: "",
			Envs: []application.EnvDTO{
				{Key: "SDK_BULK_A", Value: "value-a"},
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			_, err := client.Application.UpdateEnvsBulk(context.Background(), testComponent.UUID, testComponent.Envs)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func TestDeleteApplicationEnv(t *testing.T) {
	cases := map[string]struct {
		AppUUID string
		EnvUUID string
		Error   bool
	}{
		"ValidDelete": {
			AppUUID: existingApplicationUUID,
			EnvUUID: existingApplicationEnvUUID,
			Error:   false,
		},
		"EmptyAppUUID": {
			AppUUID: "",
			EnvUUID: existingApplicationEnvUUID,
			Error:   true,
		},
		"EmptyEnvUUID": {
			AppUUID: existingApplicationUUID,
			EnvUUID: "",
			Error:   true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			err := client.Application.DeleteEnv(context.Background(), testComponent.AppUUID, testComponent.EnvUUID)

			if err != nil && !testComponent.Error {
				t.Errorf("AppUUID (%s), EnvUUID (%s) produced an unexpected error: %v", testComponent.AppUUID, testComponent.EnvUUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("AppUUID (%s), EnvUUID (%s) did not error", testComponent.AppUUID, testComponent.EnvUUID)
			}
		})
	}
}

func TestDeleteApplication(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidDelete": {
			UUID:  existingApplicationUUID,
			Error: false,
		},
		"EmptyUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			client := sdk.Init(host, apiKey)

			err := client.Application.Delete(context.Background(), testComponent.UUID, nil)

			if err != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error: %v", testComponent.UUID, err)
			} else if err == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func boolPtr(b bool) *bool {
	return &b
}
