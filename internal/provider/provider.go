package ros

import (
	"context"

	"github.com/blueserverio/ros/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			// DataSourcesMap: map[string]*schema.Resource{
			// 	"ros_data_source": dataSourceScaffolding(),
			// },
			ResourcesMap: map[string]*schema.Resource{
				"ros_system_identity": resourceSystemIdentity(),
			},
			Schema: map[string]*schema.Schema{
				"hosturl": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("ROS_HOSTURL", nil),
					Description: "URL of the ROS router. Include the scheme (http/https)",
				},
				"username": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("ROS_USERNAME", nil),
					Description: "Username for the ROS user",
				},
				"password": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("ROS_PASSWORD", nil),
					Description: "Password for the ROS user",
					Sensitive:   true,
				},
				"insecure": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     false,
					DefaultFunc: schema.EnvDefaultFunc("ROS_INSECURE", false),
					Description: "Whether to verify the SSL certificate or not",
				},
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		userAgent := p.UserAgent("terraform-provider-ros", version)
		hostUrl := d.Get("hosturl").(string)
		username := d.Get("username").(string)
		password := d.Get("password").(string)
		insecure := d.Get("insecure").(bool)
		return client.NewClient(hostUrl, username, password, insecure, userAgent), nil
		//return &apiClient{}, nil
	}
}
