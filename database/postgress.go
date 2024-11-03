package database

func (d *DatabaseInstance) CreatePostgreSQL(data *CreateDatabaseDTO) (*string, error) {
	buf, err := encodeRequest(data)
	if err != nil {
		return nil, err
	}
	body, err := d.client.HttpRequest("databases/postgresql", "POST", *buf)
	if err != nil {
		return nil, err
	}
	response, err := decodeResponse(body, &CreateDatabaseResponse{})
	if err != nil {
		return nil, err
	}
	return &response.UUID, nil
}

type CreateDatabaseDTO struct {
	ServerUUID  *string `json:"server_uuid"`
	ProjectUUID *string `json:"project_uuid"`
	Environment *string `json:"environment_name"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`
	PublicPort  *int    `json:"public_port,omitempty"`
}

type CreateDatabaseResponse struct {
	UUID string `json:"uuid"`
}
