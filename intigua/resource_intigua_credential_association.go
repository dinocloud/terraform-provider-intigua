package intigua

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIntiguaCredentialAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntiguaCredentialAssociationCreate,
		Read:   resourceIntiguaCredentialAssociationRead,
		Update: resourceIntiguaCredentialAssociationUpdate,
		Delete: resourceIntiguaCredentialAssociationDelete,

		Schema: map[string]*schema.Schema{
			"endpoint_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"credential_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIntiguaCredentialAssociationCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaCredentialAssociationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaCredentialAssociationUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaCredentialAssociationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}