package ros

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
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

type apiClient struct {
	hostUrl string
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent

		return &apiClient{}, nil
	}
}
