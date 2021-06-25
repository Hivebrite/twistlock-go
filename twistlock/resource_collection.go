package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/collections"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spf13/cast"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		Create: createCollection,
		Read:   readCollection,
		Update: updateCollection,
		Delete: deleteCollection,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"hosts": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"images": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"labels": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"containers": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"namespaces": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"account_ids": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"clusters": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				ForceNew:    true,
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name of the collection",
			},
			"description": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "",
			},
			"color": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "",
			},
		},
	}
}

func parseCollection(d *schema.ResourceData) *sdk.Collection {
	collection := sdk.Collection{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Hosts:       cast.ToStringSlice(d.Get("hosts")),
		Images:      cast.ToStringSlice(d.Get("images")),
		Labels:      cast.ToStringSlice(d.Get("labels")),
		Containers:  cast.ToStringSlice(d.Get("containers")),
		Namespaces:  cast.ToStringSlice(d.Get("namespaces")),
		AccountIDs:  cast.ToStringSlice(d.Get("account_ids")),
		Clusters:    cast.ToStringSlice(d.Get("clusters")),
		Color:       d.Get("color").(string),
	}

	return &collection
}

func saveCollection(d *schema.ResourceData, objectCollection *sdk.Collection) error {

	d.SetId(objectCollection.Name)

	err := d.Set("name", objectCollection.Name)
	if err != nil {
		log.Printf("[ERROR] name setting caused by: %s", err)
		return err
	}

	err = d.Set("description", objectCollection.Description)
	if err != nil {
		log.Printf("[ERROR] description setting caused by: %s", err)
		return err
	}

	err = d.Set("color", objectCollection.Color)
	if err != nil {
		log.Printf("[ERROR] color setting caused by: %s", err)
		return err
	}

	err = d.Set("hosts", objectCollection.Hosts)
	if err != nil {
		log.Printf("[ERROR] hosts setting caused by: %s", err)
		return err
	}
	err = d.Set("images", objectCollection.Images)
	if err != nil {
		log.Printf("[ERROR] images setting caused by: %s", err)
		return err
	}
	err = d.Set("labels", objectCollection.Labels)
	if err != nil {
		log.Printf("[ERROR] labels setting caused by: %s", err)
		return err
	}
	err = d.Set("containers", objectCollection.Containers)
	if err != nil {
		log.Printf("[ERROR] containers setting caused by: %s", err)
		return err
	}
	err = d.Set("namespaces", objectCollection.Namespaces)
	if err != nil {
		log.Printf("[ERROR] namespaces setting caused by: %s", err)
		return err
	}
	err = d.Set("account_ids", objectCollection.AccountIDs)
	if err != nil {
		log.Printf("[ERROR] account_ids setting caused by: %s", err)
		return err
	}
	err = d.Set("clusters", objectCollection.Clusters)
	if err != nil {
		log.Printf("[ERROR] clusters setting caused by: %s", err)
		return err
	}

	return nil
}

func createCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	err := collections.Create(*client, parseCollection(d))
	if err != nil {
		return err
	}

	if err := readCollection(d, meta); err != nil {
		log.Printf("[ERROR] readCollection func caused by: %s", err)
		return err
	}

	return nil
}

func readCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	collection, err := collections.Get(*client, d.Get("name").(string))
	if err != nil {
		return err
	}

	return saveCollection(d, collection)
}

func updateCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := collections.Update(*client, parseCollection(d))
	if err != nil {
		return err
	}

	return readCollection(d, meta)
}

func deleteCollection(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	return collections.Delete(*client, d.Id())
}
