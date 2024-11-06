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
				PrivateKey:  "MIICWwIBAAKBgQC2B3b09WW8oKeVtYt7sqHs+1Ce7Kbgf60yj+DO6QU7okmpV8QU2jkjLWuAM4awKrBMuI2s5aQ62DaCXrP0k4xvdvSxhz5GrSHuvU4tn6a4TWK3nMqZE9yZYXRdpfPBf3+XHUAHY8FLJlMu49+uY7h6WxoKANDkJ1VMUBsFi7qJSQIDAQABAoGAAeLL6bfNKQolElkCK/Lq2JC2Ah+Djxnjin2RH7OsWTTSPI3rOygTpXin/3kJMTQQBYt39E2gyPdKgUlH5gXNU6nSfL3MZFMok8V9y3m73NETVQn489W7F0ykTqu5xBBt8YdbMhxjviEtzLVUlR2O6sL6aHihQV016N61Q/+ash0CQQDtidFHc20vYHGsES8fVsR5k8WtNfNkW5EDoQJHlkrF1M+2GlHAeZ0T47jfnzNwHtmv6qz+4uyhpyT995A7DJfTAkEAxC01WWE8cPlQjXFpwQh8tfee4CA4Fza/UwpI4cg9uIh8Qw6I7ZvwFLlmgnIUWQHxZD2FVbtyojshbr3Fq6hk8wJAT5rPKt3Q6n6suZhsrVj7sS7HoXuiHLDfEVNFG06PsmrWTVXWreVTsdWwICkPKPT9yQmhfi34VVhZek8b494dhwJAHYVID+kn9UAvNPmqFlg2nBNlHwj6J9QfOlnD2eYOE6TGPjkDte3PVO8JU/viv/og3xq648nPWEZZo5Z/FnPCrQJAZA3BGbng4Hx1ECUZlTYZ8q0xMaIoqjWDZd2Mf5ZwMmQ9uhyWIDkqrVgMNEwdK6AtI5twmqg/VHoKG9RSbkMSiw==",
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
