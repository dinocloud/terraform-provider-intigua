package intigua

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIntiguaCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntiguaCredentialCreate,
		Read:   resourceIntiguaCredentialRead,
		Update: resourceIntiguaCredentialUpdate,
		Delete: resourceIntiguaCredentialDelete,

		Schema: map[string]*schema.Schema{
			"accountname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ssh_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_key_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIntiguaCredentialCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaCredentialRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaCredentialUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaCredentialDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}