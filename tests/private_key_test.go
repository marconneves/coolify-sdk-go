package coolify_sdk_test

import (
	"testing"

	sdk "github.com/marconneves/coolify-sdk-go"
)

func TestListPrivateKey(t *testing.T) {
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

			_, errors := client.PrivateKey.List()

			if errors != nil && !testComponent.Error {
				t.Errorf("Host (%s), Key (%s) produced an unexpected error", testComponent.Host, testComponent.ApiKey)
			} else if errors == nil && testComponent.Error {
				t.Errorf("Host (%s), Key (%s) did not error", testComponent.Host, testComponent.ApiKey)
			}
		})
	}
}

func TestGetPrivateKey(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidRequest": {
			UUID:  "fggkoowk084k8okc8wk4g4o4",
			Error: false,
		},
		"WithoutHost": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			_, errors := client.PrivateKey.Get(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("UUID (%s)produced an unexpected error", testComponent.UUID)
			} else if errors == nil && testComponent.Error {
				t.Errorf("UUID (%s)did not error", testComponent.UUID)
			}
		})
	}
}

func TestCreatePrivateKey(t *testing.T) {
	cases := map[string]struct {
		DTO   sdk.CreatePrivateKeyDTO
		Error bool
	}{
		"ValidRequest": {
			DTO: sdk.CreatePrivateKeyDTO{
				Name:        "Test Key",
				Description: stringPtr("A test private key"),
				PrivateKey:  "private-key-content",
			},
			Error: false,
		},
		"MissingPrivateKey": {
			DTO: sdk.CreatePrivateKeyDTO{
				Name:        "Test Key",
				Description: stringPtr("A test private key"),
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			_, errors := client.PrivateKey.Create(&testComponent.DTO)

			if errors != nil && !testComponent.Error {
				t.Errorf("DTO (%v) produced an unexpected error", testComponent.DTO)
			} else if errors == nil && testComponent.Error {
				t.Errorf("DTO (%v) did not error", testComponent.DTO)
			}
		})
	}
}

func TestUpdatePrivateKey(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		DTO   sdk.UpdatePrivateKeyDTO
		Error bool
	}{
		"ValidRequest": {
			UUID: "fggkoowk084k8okc8wk4g4o4",
			DTO: sdk.UpdatePrivateKeyDTO{
				Name:        stringPtr("Updated Key"),
				Description: stringPtr("An updated private key"),
				PrivateKey:  stringPtr("updated-private-key-content"),
			},
			Error: false,
		},
		"MissingUUID": {
			UUID: "",
			DTO: sdk.UpdatePrivateKeyDTO{
				Name:        stringPtr("Updated Key"),
				Description: stringPtr("An updated private key"),
				PrivateKey:  stringPtr("updated-private-key-content"),
			},
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			errors := client.PrivateKey.Update(testComponent.UUID, &testComponent.DTO)

			if errors != nil && !testComponent.Error {
				t.Errorf("UUID (%s), DTO (%v) produced an unexpected error", testComponent.UUID, testComponent.DTO)
			} else if errors == nil && testComponent.Error {
				t.Errorf("UUID (%s), DTO (%v) did not error", testComponent.UUID, testComponent.DTO)
			}
		})
	}
}

func TestDeletePrivateKey(t *testing.T) {
	cases := map[string]struct {
		UUID  string
		Error bool
	}{
		"ValidRequest": {
			UUID:  "fggkoowk084k8okc8wk4g4o4",
			Error: false,
		},
		"MissingUUID": {
			UUID:  "",
			Error: true,
		},
	}

	for testName, testComponent := range cases {
		t.Run(testName, func(t *testing.T) {
			var client = sdk.Init(host, apiKey)

			errors := client.PrivateKey.Delete(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("UUID (%s) produced an unexpected error", testComponent.UUID)
			} else if errors == nil && testComponent.Error {
				t.Errorf("UUID (%s) did not error", testComponent.UUID)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}
