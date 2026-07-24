package main

import (
	"context"
	"flag"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/Paul-the-Wizord/terraform-repository-template/internal/provider"
)

// Version is the provider's version string, set at build time via -ldflags.
var Version = "dev"

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.New(Version), providerserver.ServeOpts{
		Address: "paul-the-wizord.github.io/example/hello",
		Debug:   debug,
	})
	if err != nil {
		panic(err)
	}
}
