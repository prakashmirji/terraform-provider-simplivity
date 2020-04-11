package simplivity

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"ovc_username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVC_USERNAME", ""),
			},
			"ovc_password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVC_PASSWORD", ""),
			},
			"ovc_host_address": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVC_HOST_ADDRESS", ""),
			},
			"ovc_cert_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVC_CERT_PATH", ""),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"svt_backup": dataSourceBackup(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"svt_backup": resourceBackup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:       d.Get("ovc_username").(string),
		Password:       d.Get("ovc_password").(string),
		OvcHostAddress: d.Get("ovc_host_address").(string),
		CertPath:       d.Get("ovc_cert_path").(string),
	}

	if err := conf.createClient(); err != nil {
		return nil, err
	}

	return &config, nil
}
