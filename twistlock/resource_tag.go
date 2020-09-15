package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceTag() *schema.Resource {
	return &schema.Resource{
		Create: createTag,
		Read:   readTag,
		Update: updateTag,
		Delete: deleteTag,

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
			"color": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The hex code color",
			},
			"vulns": {
				Optional:    true,
				Type:        schema.TypeSet,
				Description: "Vuln link to the tag",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "The vulnerability id",
							Required:    true,
						},
						"package_name": {
							Type:        schema.TypeString,
							Description: "The package name",
							Required:    true,
						},
						"comment": {
							Type:        schema.TypeString,
							Description: "comment about why the vuln have this tag",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func parseTag(d *schema.ResourceData) *sdk.Tag {
	tag := sdk.Tag{
		Name:  d.Get("name").(string),
		Color: d.Get("color").(string),
		Vulns: []sdk.Vulns{},
	}

	vulns := d.Get("vulns").(*schema.Set)
	for _, i := range vulns.List() {
		vuln := i.(map[string]interface{})
		tag.Vulns = append(
			tag.Vulns,
			sdk.Vulns{
				ID:          vuln["id"].(string),
				PackageName: vuln["package_name"].(string),
				Comment:     vuln["comment"].(string),
			})
	}

	return &tag
}

func saveTag(d *schema.ResourceData, tag *sdk.Tag) error {
	vulnTagTf := make([]interface{}, 0, len(tag.Vulns))

	d.SetId(tag.Name)

	err := d.Set("name", tag.Name)
	if err != nil {
		log.Printf("[ERROR] name setting caused by: %s", err)
		return err
	}

	err = d.Set("color", tag.Color)
	if err != nil {
		log.Printf("[ERROR] color setting caused by: %s", err)
		return err
	}

	for _, i := range tag.Vulns {
		vulnTagTf = append(
			vulnTagTf,
			map[string]interface{}{
				"id":           i.ID,
				"package_name": i.PackageName,
				"comment":      i.Comment,
			},
		)
	}

	err = d.Set("vulns", tag.Vulns)
	if err != nil {
		log.Printf("[ERROR] vuln setting caused by: %s", err)
		return err
	}

	return nil
}

func createTag(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	err := client.CreateTag(parseTag(d))
	if err != nil {
		return err
	}

	if err := readTag(d, meta); err != nil {
		log.Printf("[ERROR] readTag func caused by: %s", err)
		return err
	}

	return nil
}

func readTag(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	tag, err := client.GetTag(d.Get("name").(string))
	if err != nil {
		return err
	}

	return saveTag(d, tag)
}

func updateTag(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.UpdateTag(d.Id(), parseTag(d))
	if err != nil {
		return err
	}

	return readTag(d, meta)
}

func deleteTag(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	return client.DeleteTag(d.Id())
}
