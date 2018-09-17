package main

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hpidcock/terraform-provider-bugsnag/api"
)

func dataSourceBugsnagProject() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBugsnagProjectRead,
		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"html_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"errors_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"events_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceBugsnagProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	projectID := d.Get("project_id").(string)
	project, err := api.ReadProject(context.Background(), config.AuthToken, projectID)
	if err != nil {
		return err
	}
	d.SetId(project.ID)
	d.Set("name", project.Name)
	d.Set("slug", project.Slug)
	d.Set("api_key", project.APIKey)
	d.Set("url", project.URL)
	d.Set("html_url", project.HTMLURL)
	d.Set("errors_url", project.ErrorsURL)
	d.Set("events_url", project.EventsURL)
	return nil
}
