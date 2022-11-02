package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-provider-scaffolding/internal/provider"
)

var (
	version string = "dev"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug: debugMode,

		// TODO: update this string with the full name of your provider as used in your configs
		ProviderAddr: "registry.terraform.io/hashicorp/scaffolding",

		ProviderFunc: ros.New(version),
	}

	plugin.Serve(opts)
}
