package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/healx/terraform-provider-clearml/internal/client"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api_url": {
					Type:     schema.TypeString,
					Optional: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"CLEARML_API_URL",
					}, "https://api.clear.ml"),
				},
				"access_key": {
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"CLEARML_ACCESS_KEY",
					}, nil),
				},
				"secret_key": {
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"CLEARML_SECRET_KEY",
					}, nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{

			},
			ResourcesMap: map[string]*schema.Resource{
				"clearml_queue": resourceQueue(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		userAgent := p.UserAgent("terraform-provider-clearml", version)
		c, err := client.NewClearMLClient(
			ctx, 
			userAgent,
			d.Get("access_key").(string), 
			d.Get("secret_key").(string),
			d.Get("api_url").(string))

		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, nil
	}
}
