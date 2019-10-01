package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/client"
	"terraform-provider/pkg/resource"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_PORT", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_TOKEN", ""),
			},
		},
		ResourcesMap:  map[string]*schema.Resource{
			"example_item": resource.Item(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	port := d.Get("port").(int)
	token := d.Get("token").(string)

	transport := transport.Endpoint{
		Password: token,
		Host:     address,
		Port:     port,
	}

	trspt, err := client.NewClient(&transport)
	if err != nil {
		return nil, nil
	}

	return trspt, nil
}
