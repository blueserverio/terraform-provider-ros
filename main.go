package main

import (
	"flag"

	ros "github.com/blueserverio/ros/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name ros

var (
	version string = "dev"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug: debugMode,

		ProviderAddr: "blueserverio/ros",

		ProviderFunc: ros.New(version),
	}

	plugin.Serve(opts)
}
