package okta

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

var (
	// contextTimeout is the global context timeout for requests to complete.
	contextTimeout = 15 * time.Second
)

// Provider for Okta
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},

			"base_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"okta_group": resourceGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		BaseURL:  d.Get("base_url").(string),
		APIToken: d.Get("api_token").(string),
	}

	log.Println("[INFO] Initializing Okta client")
	return config.Client()
}

func contextWithTimeout() (context.Context, func()) {
	return context.WithTimeout(context.Background(), contextTimeout)
}
