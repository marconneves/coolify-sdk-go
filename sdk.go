package coolify_sdk

import (
	"bytes"
	"io"
	"net/http"

	client "github.com/marconneves/coolify-sdk-go/client"

	database "github.com/marconneves/coolify-sdk-go/database"
	server "github.com/marconneves/coolify-sdk-go/server"
)

type Sdk struct {
	Client     client.Client
	httpClient *http.Client

	Api        *ApiInstance
	Team       *TeamInstance
	Server     *server.ServerInstance
	PrivateKey *PrivateKeyInstance
	Project    *ProjectInstance
	Database   *database.DatabaseInstance
}

func Init(hostname string, apiToken string) *Sdk {
	sdk := &Sdk{
		httpClient: &http.Client{},
	}

	sdk.Client = *client.NewClient(hostname, apiToken)

	sdk.Api = &ApiInstance{client: &sdk.Client}
	sdk.Team = &TeamInstance{client: &sdk.Client}
	sdk.Server = server.NewServer(&sdk.Client)
	sdk.Database = database.NewDatabaseInstance(&sdk.Client)
	sdk.PrivateKey = &PrivateKeyInstance{client: &sdk.Client}
	sdk.Project = &ProjectInstance{client: &sdk.Client}

	return sdk
}

func (c *Sdk) HeathCheck() (*string, error) {
	body, err := c.Client.HttpRequest("healthcheck", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	defer body.Close()

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	var healCheckResponse string

	if string(bodyBytes) == "OK" {
		healCheckResponse = "success"
	} else {
		healCheckResponse = "failure"
	}

	return &healCheckResponse, nil
}

type CreateServerDTO = server.CreateServerDTO
type UpdateServerDTO = server.UpdateServerDTO

type UpdateDatabaseDTO = database.UpdateDatabaseDTO
type CreateDatabasePostgresDTO = database.CreateDatabasePostgresDTO
