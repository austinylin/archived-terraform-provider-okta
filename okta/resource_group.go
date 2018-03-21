package okta

import (
	"log"

	"github.com/austinylin/go-okta/okta"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_membership_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*okta.Client)

	profile := &okta.GroupProfile{
		Name:        d.Get("name").(string),
		Description: "[Managed by TF] " + d.Get("description").(string),
	}

	log.Printf("[DEBUG] Creating Okta group...")

	ctx, cancel := contextWithTimeout()
	defer cancel()
	group, _, err := client.Groups.Add(ctx, profile)
	if err != nil {
		return errors.Wrap(err, "failed to create group")
	}

	d.SetId(group.ID)
	log.Printf("[INFO] Group ID: %s", d.Id())

	return resourceGroupRead(d, meta)
}

func resourceGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*okta.Client)
	ctx, cancel := contextWithTimeout()
	defer cancel()
	group, _, err := client.Groups.GetByID(ctx, d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to read group")
	}

	d.Set("name", group.Profile.Name)
	d.Set("description", group.Profile.Description)
	d.Set("type", group.Type)
	d.Set("created", group.Created)
	d.Set("last_updated", group.LastUpdated.Time.String())
	d.Set("last_membership_updated", group.LastMembershipUpdated.Time.String())

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*okta.Client)

	profile := &okta.GroupProfile{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	ctx, cancel := contextWithTimeout()
	defer cancel()
	group, _, err := client.Groups.UpdateWithProfile(ctx, d.Id(), profile)
	if err != nil {
		return errors.Wrap(err, "failed to update group")
	}

	d.SetId(group.ID)

	return resourceGroupRead(d, meta)
}

func resourceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*okta.Client)
	ctx, cancel := contextWithTimeout()
	defer cancel()
	_, err := client.Groups.Remove(ctx, d.Id())
	if err != nil {
		return errors.Wrap(err, "failed to delete group")
	}

	d.SetId("")

	return nil
}
