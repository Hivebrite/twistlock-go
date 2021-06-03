package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/tag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func parseTag(d *schema.ResourceData) *tag.Tag {
	objectTag := tag.Tag{
		Name:  d.Get("name").(string),
		Color: d.Get("color").(string),
		Vulns: []tag.Vuln{},
	}

	vulns := d.Get("vulns").(*schema.Set)
	for _, i := range vulns.List() {
		vuln := i.(map[string]interface{})
		objectTag.Vulns = append(
			objectTag.Vulns,
			tag.Vuln{
				ID:          vuln["id"].(string),
				PackageName: vuln["package_name"].(string),
				Comment:     vuln["comment"].(string),
			})
	}

	return &objectTag
}

func saveTag(d *schema.ResourceData, objectTag *tag.Tag) error {
	var vulnTag []map[string]interface{}

	d.SetId(objectTag.Name)

	err := d.Set("name", objectTag.Name)
	if err != nil {
		log.Printf("[ERROR] name setting caused by: %s", err)
		return err
	}

	err = d.Set("color", objectTag.Color)
	if err != nil {
		log.Printf("[ERROR] color setting caused by: %s", err)
		return err
	}

	for _, i := range objectTag.Vulns {
		vulnTag = append(
			vulnTag,
			map[string]interface{}{
				"id":           i.ID,
				"package_name": i.PackageName,
				"comment":      i.Comment,
			},
		)
	}

	err = d.Set("vulns", vulnTag)
	if err != nil {
		log.Printf("[ERROR] vuln setting caused by: %s", err)
		return err
	}

	return nil
}

func createTag(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	err := tag.Create(*client, parseTag(d))
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
	tag, err := tag.Get(*client, d.Get("name").(string))
	if err != nil {
		return err
	}

	return saveTag(d, tag)
}

func updateTag(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := tag.Update(*client, d.Id(), parseTag(d))
	if err != nil {
		return err
	}

	return readTag(d, meta)
}

func deleteTag(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	return tag.Delete(*client, d.Id())
}
