package coolify_sdk_test

import (
	"testing"

	sdk "github.com/marconneves/coolify-sdk-go"
)

func TestListProjects(t *testing.T) {
	var client = sdk.Init(host, apiKey)

	_, err := client.Project.List()
	if err != nil {
		t.Errorf("Project listing failed unexpectedly: %v", err)
	}
}

func TestGetProject(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidUUID": {
			UUID:  "v8ckogcwgo0sgsogwooww84c",
			Error: false,
		},
		"InvalidUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			_, err := client.Project.Get(testComponent.UUID)

			if err != nil && !testComponent.Error {
				t.Errorf("Project retrieval failed unexpectedly: %v", err)
			} else if err == nil && testComponent.Error {
				t.Errorf("Project retrieval succeeded unexpectedly")
			}
		})
	}
}

func TestCreateProject(t *testing.T) {
	cases := map[string]struct {
		Project *sdk.CreateProjectDTO
		Error   bool
	}{
		"ValidProject": {
			Project: &sdk.CreateProjectDTO{
				Name:        stringPtr("New Project"),
				Description: stringPtr("Project Description"),
			},
			Error: false,
		},
		"MissingName": {
			Project: &sdk.CreateProjectDTO{
				Description: stringPtr("No Name Project"),
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			uuid, err := client.Project.Create(testComponent.Project)

			if err != nil && !testComponent.Error {
				t.Errorf("Project creation failed unexpectedly: %v", err)
			} else if err == nil && testComponent.Error {
				t.Errorf("Project creation succeeded unexpectedly: %s", *uuid)

				client.Project.Delete(*uuid)
			}
		})
	}
}

func TestUpdateProject(t *testing.T) {
	cases := map[string]struct {
		UUID    string
		Project *sdk.UpdateProjectDTO
		Error   bool
	}{
		"ValidUpdate": {
			UUID: "v8ckogcwgo0sgsogwooww84c",
			Project: &sdk.UpdateProjectDTO{
				Name:        stringPtr("Updated Project"),
				Description: stringPtr("Updated Description"),
			},
			Error: false,
		},
		"InvalidUUID": {
			UUID: "",
			Project: &sdk.UpdateProjectDTO{
				Name:        stringPtr("Updated Project"),
				Description: stringPtr("Updated Description"),
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			err := client.Project.Update(testComponent.UUID, testComponent.Project)

			if err != nil && !testComponent.Error {
				t.Errorf("Project update failed unexpectedly: %v", err)
			} else if err == nil && testComponent.Error {
				t.Errorf("Project update succeeded unexpectedly")
			}
		})
	}
}

func TestDeleteProject(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidDelete": {
			UUID:  "v8ckogcwgo0sgsogwooww84c",
			Error: false,
		},
		"InvalidUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			err := client.Project.Delete(testComponent.UUID)

			if err != nil && !testComponent.Error {
				t.Errorf("Project deletion failed unexpectedly: %v", err)
			} else if err == nil && testComponent.Error {
				t.Errorf("Project deletion succeeded unexpectedly")
			}
		})
	}
}

func TestGetEnvironment(t *testing.T) {
	cases := map[string]struct {
		UUID        string
		Environment string
		Error       bool
	}{
		"ValidRequest": {
			UUID:        "v8ckogcwgo0sgsogwooww84c",
			Environment: "dev",
			Error:       false,
		},
		"InvalidUUID": {
			UUID:        "",
			Environment: "dev",
			Error:       true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			_, err := client.Project.Environment(testComponent.UUID, testComponent.Environment)

			if err != nil && !testComponent.Error {
				t.Errorf("Environment retrieval failed unexpectedly: %v", err)
			} else if err == nil && testComponent.Error {
				t.Errorf("Environment retrieval succeeded unexpectedly")
			}
		})
	}
}
