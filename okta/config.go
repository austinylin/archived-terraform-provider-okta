package okta

import (
	"log"

	"github.com/austinylin/go-okta/okta"
)

// Config is a struct holding configuration information for the okta terraform provider.
type Config struct {
	APIToken string
	BaseURL  string
}

// Client returns a new Service for accessing Okta.
func (c *Config) Client() (*okta.Client, error) {
	client, err := okta.NewClient(c.APIToken, c.BaseURL, nil)
	if err != nil {
		log.Fatalf("[FATAL] Unable to configure Okta Client: %s", err)
	}
	log.Printf("[INFO] Okta Client configured")

	return client, nil
}
