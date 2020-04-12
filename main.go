package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/prakashmirji/terraform-provider-simplivity/simplivity"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: simplivity.Provider,
	})
}
