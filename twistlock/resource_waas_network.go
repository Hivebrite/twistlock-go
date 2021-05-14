package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/waas"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceWaasNetwork() *schema.Resource {
	return &schema.Resource{
		Create: createWaasNetwork,
		Read:   readWaasNetwork,
		Update: createWaasNetwork,
		Delete: deleteWaasNetwork,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"container_enabled": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"host_enabled": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
		},
	}
}

func parseWaasNetwork(d *schema.ResourceData) *waas.NetworkWaas {
	return &waas.NetworkWaas{
		ContainerEnabled: d.Get("container_enabled").(bool),
		HostEnabled:      d.Get("host_enabled").(bool),
	}
}

func saveWaasNetwork(d *schema.ResourceData, waasObject *waas.NetworkWaas) error {
	d.SetId("networkFirewall")
	d.Set("host_enabled", waasObject.HostEnabled)
	d.Set("container_enabled", waasObject.ContainerEnabled)

	return nil
}

func createWaasNetwork(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := waas.SetNetworkWaas(*client, parseWaasNetwork(d))
	if err != nil {
		return err
	}

	return readWaasNetwork(d, meta)
}

func readWaasNetwork(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	networkWaas, err := waas.GetNetworkWaas(*client)
	if err != nil {
		return err
	}

	return saveWaasNetwork(d, networkWaas)
}

func deleteWaasNetwork(d *schema.ResourceData, meta interface{}) error {
	return nil
}
