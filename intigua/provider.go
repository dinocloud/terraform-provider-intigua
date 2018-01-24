package intigua

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

func Provider() terraform.ResourceProvider  {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {

				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTIGUA_HOST", nil),
				Description: descriptions["intigua_host"],
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTIGUA_USER", nil),
				Description: descriptions["intigua_username"],
			},

			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTIGUA_API_KEY", nil),
				Description: descriptions["intigua_api_key"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"intigua_endpoint": resourceIntiguaEndpoint(),
		},
		ConfigureFunc: providerConfigure,

	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"intigua_host": "The host where intigua is deployed (aka. the IP or the DNS name)",
		"intigua_username": "The username with permissions to access the intigua Rest API",
		"intigua_api_key": "The API Key to authenticate",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error){
	config := Config{
		Host:		d.Get("host").(string),
		Username:	d.Get("username").(string),
		ApiKey:		d.Get("api_key").(string),
	}
	log.Println("[INFO] Initializing Intigua client")
	return config.Client()
}
