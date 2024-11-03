package coolify_sdk

import (
	"encoding/json"

	client "github.com/marconneves/coolify-sdk-go/client"
)

type ApiInstance struct {
	client *client.Client
}

type EnableApiResponse struct {
	Message string `json:"message"`
}

func (a *ApiInstance) Enable() (*string, error) {
	body, err := a.client.HttpRequest("enable", "GET")
	if err != nil {
		return nil, err
	}

	response := &EnableApiResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	var apiEnabled string
	if response.Message == "API Enabled." {
		apiEnabled = "success"
	} else {
		apiEnabled = "failure"
	}

	return &apiEnabled, nil
}

func (a *ApiInstance) Disable() (*string, error) {
	body, err := a.client.HttpRequest("disable", "GET")
	if err != nil {
		return nil, err
	}

	response := &EnableApiResponse{}
	err = json.NewDecoder(body).Decode(response)
	if err != nil {
		return nil, err
	}

	var apiDisabled string
	if response.Message == "API disabled." {
		apiDisabled = "success"
	} else {
		apiDisabled = "failure"
	}

	return &apiDisabled, nil
}
