package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIntiguaEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceIntiguaEndpointCreate,
		Read:   resourceIntiguaEndpointRead,
		Update: resourceIntiguaEndpointUpdate,
		Delete: resourceIntiguaEndpointDelete,

		Schema: map[string]*schema.Schema{
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dnsname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"OS": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIntiguaEndpointCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaEndpointRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaEndpointUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceIntiguaEndpointDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}