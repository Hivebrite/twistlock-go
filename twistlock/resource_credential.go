package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceCredentialProvider() *schema.Resource {
	return &schema.Resource{
		Create: createCredentialProvider,
		Read:   readCredentialProvider,
		Update: updateCredentialProvider,
		Delete: deleteCredentialProvider,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "Name of the credentials",
				Required:    true,
				ForceNew:    true,			
			},
			"secret": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "GCP service account token",
				MinItems:    1,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plain": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "plain value of the token",
						},
					},
				},	
			}
			"type": {
				Type:        schema.TypeString,
				Description: "Type of credentials",
				Required:    true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						"aws", "azure", "gcp", "ibmCloud", "apiToken", "basic", "dtr", "kubeconfig", "certificate"
					},
					false,
				),
				
			},
		},
	}
}

func parseCredentialprovider(d *schema.ResourceData) *sdk.ProviderCredential {
	credential := sdk.ProviderCredential{
		Type:  d.Get("type").(string),
		ID: d.Get("id").(string),
		Secret: []sdk.Secret{},
	}


	secret := d.Get("secret").(*schema.Set).(map[string]interface{})
	credential.Secret = append(credential.Secret, sdk.Secret{
		Plain: secret["plain"].(string),
	})
	
	return &credential
}

func saveCredentialProvider(d *schema.ResourceData, credential *sdk.ProviderCredential) error {
	credentialSecret := make([]interface{}, 0, len(credential.Secret))

	d.SetId(credential.ID)

	err := d.Set("id", credential.ID)
	if err != nil {
		log.Printf("[ERROR] id setting caused by: %s", err)
		return err
	}

	err = d.Set("type", credential.Type)
	if err != nil {
		log.Printf("[ERROR] type setting caused by: %s", err)
		return err
	}

	for _, i := range credential.Secret {
		credentialSecret = append(
			credentialSecret,
			map[string]interface{}{
				"plain":           i.Plain,
			},
		)
	}

	err = d.Set("secret", credential.Secret)
	if err != nil {
		log.Printf("[ERROR] secret setting caused by: %s", err)
		return err
	}

	return nil
}

func createCredentialrovider(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)

	err := client.CreateCredentialProvider(parseCredentialprovider(d))
	if err != nil {
		return err
	}

	if err := readCredentialProvider(d, meta); err != nil {
		log.Printf("[ERROR] readCredentialprovider func caused by: %s", err)
		return err
	}

	return nil
}

func readCredentialProvider(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	credential, err := client.GetCredentialProvider(d.Get("id").(string))
	if err != nil {
		return err
	}

	return saveCredentialprovider(d, credential)
}

func updateCredentialProvider(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.UpdateCredentialProvider(d.Id(), parseCredentialprovider(d))
	if err != nil {
		return err
	}

	return readCredentialprovider(d, meta)
}

func deleteCredentialProvider(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	return client.DeleteCredentialprovider(d.Id())
}
