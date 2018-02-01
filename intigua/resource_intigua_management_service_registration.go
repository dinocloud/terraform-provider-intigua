package intigua

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIntiguaManagementServiceRegistration() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntiguaManagementServiceRegistrationCreate,
		Read:   resourceIntiguaManagementServiceRegistrationRead,
		Update: resourceIntiguaManagementServiceRegistrationUpdate,
		Delete: resourceIntiguaManagementServiceRegistrationDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"configpackage": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"endpoint_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIntiguaManagementServiceRegistrationCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaManagementServiceRegistrationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaManagementServiceRegistrationUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaManagementServiceRegistrationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}