package main

import (
	"github.com/Hivebrite/twistlock-go/twistlock"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: twistlock.Provider,
	})
}
