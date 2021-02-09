package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	subnet "github.com/Hivebrite/twistlock-go/sdk/subnets"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		Create: createSubnet,
		Read:   readSubnet,
		Update: updateSubnet,
		Delete: deleteSubnet,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "",
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Description of the Subnet",
			},
			"subnets": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func parseSubnet(d *schema.ResourceData) *subnet.Subnet {
	objectSubnet := subnet.Subnet{
		ID:          d.Get("name").(string),
		Description: d.Get("description").(string),
		Subnets:     []string{},
	}

	subnets := d.Get("subnets").([]interface{})
	for _, i := range subnets {
		subnet := i.(string)
		objectSubnet.Subnets = append(
			objectSubnet.Subnets,
			subnet,
		)
	}

	return &objectSubnet
}

func saveSubnet(d *schema.ResourceData, objectSubnet *subnet.Subnet) error {
	d.SetId(objectSubnet.ID)

	err := d.Set("name", objectSubnet.ID)
	if err != nil {
		log.Printf("[ERROR] name setting caused by: %s", err)
		return err
	}

	err = d.Set("description", objectSubnet.Description)
	if err != nil {
		log.Printf("[ERROR] description setting caused by: %s", err)
		return err
	}

	err = d.Set("subnets", objectSubnet.Subnets)
	if err != nil {
		log.Printf("[ERROR] subnets setting caused by: %s", err)
		return err
	}

	return nil
}

func createSubnet(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	err := subnet.Create(*client, parseSubnet(d))
	if err != nil {
		return err
	}

	if err := readSubnet(d, meta); err != nil {
		log.Printf("[ERROR] readSubnet func caused by: %s", err)
		return err
	}

	return nil
}

func readSubnet(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	subnet, err := subnet.Get(*client, d.Get("name").(string))
	if err != nil {
		return err
	}

	return saveSubnet(d, subnet)
}

func updateSubnet(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := subnet.Update(*client, parseSubnet(d))
	if err != nil {
		return err
	}

	return readSubnet(d, meta)
}

func deleteSubnet(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	return subnet.Delete(*client, d.Id())
}
