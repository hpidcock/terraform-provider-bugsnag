package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Config provides configuration to resource/data source CRUD methods.
type Config struct {
	AuthToken string
}

func provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credentials": &schema.Schema{
				Type:     schema.TypeString,
				Optional: false,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"BUGSNAG_ACCESS_TOKEN",
				}, nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bugsnag_project": dataSourceBugsnagProject(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := Config{
		AuthToken: d.Get("credentials").(string),
	}
	return &c, nil
}
