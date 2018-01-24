package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/dinocloud/terraform-provider-intigua/intigua"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: intigua.Provider})
}