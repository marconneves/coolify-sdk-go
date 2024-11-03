package coolify_sdk_test

import (
	sdk "coolify-sdk"
	"testing"
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
			var client = sdk.NewClient(testComponent.Host, testComponent.ApiKey)

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
			var client = sdk.NewClient(host, apiKey)

			_, errors := client.PrivateKey.Get(testComponent.UUID)

			if errors != nil && !testComponent.Error {
				t.Errorf("UUID (%s)produced an unexpected error", testComponent.UUID)
			} else if errors == nil && testComponent.Error {
				t.Errorf("UUID (%s)did not error", testComponent.UUID)
			}
		})
	}
}
