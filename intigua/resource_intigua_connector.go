package intigua

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIntiguaConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntiguaConectorCreate,
		Read:   resourceIntiguaConectorRead,
		Update: resourceIntiguaConectorUpdate,
		Delete: resourceIntiguaConectorDelete,

		Schema: map[string]*schema.Schema{
			"endpoint_connector_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIntiguaConectorCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaConectorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaConectorUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaConectorDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}