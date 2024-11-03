package main

import (
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func main() {

	sdk := coolify_sdk.Init("", "")

	sdk.Server.Create(&coolify_sdk.CreateServerDTO{})
}
