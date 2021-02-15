package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/credentials"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCredentialProvider() *schema.Resource {
	return &schema.Resource{
		Create: SetProviderCredential,
		Read:   readProviderCredential,
		Update: SetProviderCredential,
		Delete: deleteProviderCredential,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the credentials",
				Required:    true,
				ForceNew:    true,
			},
			"account_id": {
				Type:        schema.TypeString,
				Description: "ID of the onboarded account (if using onboarded PC account)",
				Optional:    true,
			},
			"secret": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "GCP service account token",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plain": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "plain value of the token",
						},
					},
				},
			},
			"external": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of credentials",
				Required:    true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						"aws", "azure", "gcp", "ibmCloud", "apiToken", "basic", "dtr", "kubeconfig", "certificate",
					},
					false,
				),
			},
		},
	}
}

func parseProviderCredential(d *schema.ResourceData) *credentials.ProviderCredential {
	secret := d.Get("secret").(map[string]interface{})
	secretObject := sdk.Secret{}
	if secret["plain"] != nil {
		secretObject.Plain = secret["plain"].(string)
	}
	credential := credentials.ProviderCredential{
		Type:      d.Get("type").(string),
		ID:        d.Get("name").(string),
		AccountID: d.Get("account_id").(string),
		External:  d.Get("external").(bool),
		Secret:    secretObject,
	}

	return &credential
}

func saveCredentialProvider(d *schema.ResourceData, credential *credentials.ProviderCredential) error {
	d.SetId(credential.ID)

	err := d.Set("name", credential.ID)
	if err != nil {
		log.Printf("[ERROR] id setting caused by: %s", err)
		return err
	}

	err = d.Set("type", credential.Type)
	if err != nil {
		log.Printf("[ERROR] type setting caused by: %s", err)
		return err
	}

	err = d.Set("external", credential.External)
	if err != nil {
		log.Printf("[ERROR] external setting caused by: %s", err)
		return err
	}

	err = d.Set("account_id", credential.AccountID)
	if err != nil {
		log.Printf("[ERROR] account_id setting caused by: %s", err)
		return err
	}

	credentialSecret := map[string]interface{}{
		"plain": credential.Secret.Plain,
	}
	err = d.Set("secret", credentialSecret)
	if err != nil {
		log.Printf("[ERROR] secret setting caused by: %s", err)
		return err
	}

	return nil
}

func SetProviderCredential(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	err := credentials.Set(*client, parseProviderCredential(d))
	if err != nil {
		return err
	}

	if err := readProviderCredential(d, meta); err != nil {
		log.Printf("[ERROR] readProviderCredential func caused by: %s", err)
		return err
	}

	return nil
}

func readProviderCredential(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	credential, err := credentials.Get(*client, d.Get("name").(string))
	if err != nil {
		return err
	}

	return saveCredentialProvider(d, credential)
}

func deleteProviderCredential(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	return credentials.Delete(*client, d.Id())
}
