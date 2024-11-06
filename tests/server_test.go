package coolify_sdk_test

import (
	"testing"

	sdk "github.com/marconneves/coolify-sdk-go"
	server "github.com/marconneves/coolify-sdk-go/server"
)

func TestListServer(t *testing.T) {
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
		"WithoutHost": {
			Host:   host,
			ApiKey: "",
			Error:  true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(testComponent.Host, testComponent.ApiKey)

			_, errors := client.Server.List()

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s) did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestGetServer(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		UUID   string
		Error  bool
	}{
		"ValidRequest": {
			Host:   host,
			ApiKey: apiKey,
			UUID:   "lo4sksgsks8kw8w0skog8c0s",
			Error:  false,
		},
		"WithInvalidId": {
			Host:   host,
			ApiKey: apiKey,
			UUID:   "",
			Error:  true,
		},
		"WithoutHost": {
			Host:   host,
			ApiKey: "",
			UUID:   "",
			Error:  true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(testComponent.Host, testComponent.ApiKey)

			_, errors := client.Server.Get(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s) did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestCreateServer(t *testing.T) {
	cases := map[string]struct {
		Server *server.CreateServerDTO
		Error  bool
	}{
		"ValidServer": {
			Server: &server.CreateServerDTO{
				Name:            "My Server",
				Description:     "My Server Description",
				IP:              "127.0.0.1",
				Port:            22,
				User:            "root",
				PrivateKeyUUID:  "fggkoowk084k8okc8wk4g4o4",
				IsBuildServer:   true,
				InstantValidate: true,
			},
			Error: false,
		},
		"MissingName": {
			Server: &server.CreateServerDTO{
				Description:     "No Name Server",
				IP:              "127.0.0.1",
				Port:            22,
				User:            "root",
				PrivateKeyUUID:  "fggkoowk084k8okc8wk4g4o4",
				IsBuildServer:   true,
				InstantValidate: true,
			},
			Error: true,
		},
		"InvalidPrivateKey": {
			Server: &server.CreateServerDTO{
				Name:            "No PrivateKey UUID Valid",
				Description:     "Invalid IP",
				IP:              "127.0.0.1",
				Port:            22,
				User:            "root",
				PrivateKeyUUID:  "asjdaksdhjaskljdha",
				IsBuildServer:   true,
				InstantValidate: true,
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			uuid, errors := client.Server.Create(testComponent.Server)

			if errors != nil && !testComponent.Error {
				t.Errorf("Server creation failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Server creation succeeded unexpectedly: %s", *uuid)
			}
		})
	}
}

func TestUpdateServer(t *testing.T) {
	cases := map[string]struct {
		UUID   string
		Server *server.UpdateServerDTO
		Error  bool
	}{
		"ValidUpdate": {
			UUID: "lcs8ggw8cos48kw0sc0sk0gc",
			Server: &server.UpdateServerDTO{
				Name:        "Updated Server",
				Description: "Updated Description",
				IP:          "192.168.1.1",
				Port:        22,
				User:        "admin",
			},
			Error: false,
		},
		"InvalidUUID": {
			UUID: "",
			Server: &server.UpdateServerDTO{
				Name:        "Updated Server",
				Description: "Updated Description",
				IP:          "192.168.1.1",
				Port:        22,
				User:        "admin",
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			errors := client.Server.Update(testComponent.UUID, testComponent.Server)

			if errors != nil && !testComponent.Error {
				t.Errorf("Server update failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Server update succeeded unexpectedly")
			}
		})
	}
}

func TestDeleteServer(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidDelete": {
			UUID:  "lcs8ggw8cos48kw0sc0sk0gc",
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

			errors := client.Server.Delete(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Server deletion failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Server deletion succeeded unexpectedly")
			}
		})
	}
}

func TestListDomains(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidRequest": {
			UUID:  "lo4sksgsks8kw8w0skog8c0s",
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

			_, errors := client.Server.Domains(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Domain listing failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Domain listing succeeded unexpectedly")
			}
		})
	}
}

func TestListResources(t *testing.T) {
	cases := map[string]struct {
		Host   string
		ApiKey string
		UUID   string
		Error  bool
	}{
		"ValidRequest": {
			UUID:  "lo4sksgsks8kw8w0skog8c0s",
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

			_, errors := client.Server.Resources(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Resource listing failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Resource listing succeeded unexpectedly")
			}
		})
	}
}

func TestValidateServer(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidUUID": {
			UUID:  "lo4sksgsks8kw8w0skog8c0s",
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

			errors := client.Server.Validate(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("Server validation failed unexpectedly: %v", errors)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Server validation succeeded unexpectedly")
			}
		})
	}
}
