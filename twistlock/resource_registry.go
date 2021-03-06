package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/registry"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRegistrySettings() *schema.Resource {
	return &schema.Resource{
		Create: createRegistrySettings,
		Read:   readRegistrySettings,
		Update: createRegistrySettings,
		Delete: deleteRegistrySettings,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Model for the registry settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version": {
							Type:        schema.TypeString,
							Description: "",
							Default:     "2",
							Optional:    true,
						},
						"registry": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "Contains either the registry name (e.g., gcr.io) or url (e.g., https://gcr.io)",
						},
						"repository": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "The repository name to scan",
						},
						"tag": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "tag to scan, wildcard is supported",
						},
						"os": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "indicates the registry images base OS. Range of acceptable values: linux, windows",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"linux",
									"windows",
								},
								false,
							),
						},
						"cap": {
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Indicates only the last k images should be fetched",
							Default:     5,
						},
						"hostname": {
							Type:        schema.TypeString,
							Description: "The hostname of the defender that is used as registry scanner",
							Optional:    true,
							Default:     "",
						},
						"scanners": {
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Indicates the amount of defenders assigned to scan this registry, this applies only for registries with auto-selected defenders",
							Default:     2,
						},
						"namespace": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "",
							Default:     "",
						},
						"use_aws_role": {
							Optional:    true,
							Type:        schema.TypeBool,
							Description: "",
							Default:     false,
						},
						"credential": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The credential id",
							Default:     "",
						},
						"role_arn": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "",
							Default:     "",
						},
						"version_pattern": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "",
							Default:     "",
						},
					},
				},
			},
		},
	}
}

func parseRegistrySettings(d *schema.ResourceData, client *sdk.Client) (*registry.Specifications, error) {
	spec := registry.Specifications{}
	settings := d.Get("registry").(*schema.Set)
	for _, i := range settings.List() {
		setting := i.(map[string]interface{})

		spec.Settings = append(
			spec.Settings,
			registry.Setting{
				Version:        setting["version"].(string),
				Registry:       setting["registry"].(string),
				Repository:     setting["repository"].(string),
				Tag:            setting["tag"].(string),
				Os:             setting["os"].(string),
				Cap:            setting["cap"].(int),
				Hostname:       setting["hostname"].(string),
				Scanners:       setting["scanners"].(int),
				UseAWSRole:     setting["use_aws_role"].(bool),
				CredentialID:   setting["credential"].(string),
				RoleArn:        setting["role_arn"].(string),
				VersionPattern: setting["version_pattern"].(string),
			})
	}

	return &spec, nil
}

func saveRegistrySettings(d *schema.ResourceData, spec *registry.Specifications) error {
	specRegistryTf := make([]interface{}, 0, len(spec.Settings))

	for _, i := range spec.Settings {
		specRegistryTf = append(
			specRegistryTf,
			map[string]interface{}{
				"version":    i.Version,
				"registry":   i.Registry,
				"repository": i.Repository,
				"tag":        i.Tag,
				"os":         i.Os,
				"cap":        i.Cap,
				"hostname":   i.Hostname,
				"scanners":   i.Scanners,
				"credential": i.CredentialID,
			})
	}

	d.SetId("registry")

	err := d.Set("registry", specRegistryTf)
	if err != nil {
		log.Printf("[ERROR] registry setting caused by: %s", err)
	}

	return err
}

func createRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	settings, err := parseRegistrySettings(d, client)
	if err != nil {
		return err
	}

	err = registry.Set(*client, settings)
	if err != nil {
		return err
	}

	return readRegistrySettings(d, meta)
}

func readRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	registries, err := registry.Index(*client)
	if err != nil {
		return err
	}

	return saveRegistrySettings(d, registries)
}

func deleteRegistrySettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := registry.Set(*client, &registry.Specifications{})
	if err != nil {
		return err
	}

	return nil
}
