package main

import (
	"github.com/austinylin/terraform-provider-okta/okta"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return okta.Provider()
		},
	})
}
