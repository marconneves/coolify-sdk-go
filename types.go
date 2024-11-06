package coolify_sdk

import (
	database "github.com/marconneves/coolify-sdk-go/database"
	server "github.com/marconneves/coolify-sdk-go/server"
)

type CreateServerDTO = server.CreateServerDTO
type CreateServerResponse = server.CreateServerResponse
type UpdateServerDTO = server.UpdateServerDTO
type Server = server.Server
type Proxy = server.Proxy
type Settings = server.Settings
type Resource = server.Resource

type UpdateDatabaseDTO = database.UpdateDatabaseDTO
type CreateDatabasePostgresDTO = database.CreateDatabasePostgresDTO
type CreateDatabasePostgresResponse = database.CreateDatabasePostgresResponse
type CreateDatabaseRedisDTO = database.CreateDatabaseRedisDTO
type CreateDatabaseRedisResponse = database.CreateDatabaseRedisResponse
type Database = database.Database
type Destination = database.Destination
