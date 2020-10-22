package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/spf13/cast"
)

// TODO(This resource create update correctly on prismacloud but at each run the change is trigger, we will investigate later on that)
func resourceWaasContainer() *schema.Resource {
	return &schema.Resource{
		Create: createWaasContainer,
		Read:   readWaasContainer,
		Update: createWaasContainer,
		Delete: deleteWaasContainer,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_port": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     31000,
			},
			"min_port": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     30000,
			},
			"rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"resources": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
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
								},
							},
						},
						"applications_spec": {
							Required:    true,
							Type:        schema.TypeList,
							Description: "firewalls rules to be applied to the resources",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate": {
										Type:        schema.TypeSet,
										Optional:    true,
										Computed:    true,
										Description: "X509 Certificate",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"encrypted": {
													Required:    true,
													Type:        schema.TypeString,
													Description: "Day",
												},
											},
										},
									},
									"api_spec": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "list of URLs to protect",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"endpoints": {
													Type:        schema.TypeSet,
													Required:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"host": {
																Required:    true,
																Type:        schema.TypeString,
																Description: "",
															},
															"base_path": {
																Required:    true,
																Type:        schema.TypeString,
																Description: "",
															},
															"exposed_port": {
																Optional:    true,
																Default:     0,
																Type:        schema.TypeInt,
																Description: "",
															},
															"internal_port": {
																Required:    true,
																Type:        schema.TypeInt,
																Description: "",
															},
															"tls": {
																Required:    true,
																Type:        schema.TypeBool,
																Description: "",
															},
															"http2": {
																Required:    true,
																Type:        schema.TypeBool,
																Description: "",
															},
														},
													},
												},
											},
										},
									},
									"network_controls": {
										Required:    true,
										Type:        schema.TypeSet,
										Description: "",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"advanced_protection_effect": {
													Optional:    true,
													Type:        schema.TypeString,
													Description: "",
													Default:     "alert",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
												"denied_subnets_effect": {
													Optional:    true,
													Type:        schema.TypeString,
													Description: "",
													Default:     "alert",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
												"denied_countries_effect": {
													Optional:    true,
													Type:        schema.TypeString,
													Description: "",
													Default:     "alert",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
												"allowed_countries_effect": {
													Optional:    true,
													Type:        schema.TypeString,
													Description: "",
													Default:     "alert",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
												"denied_subnets": {
													Required:    true,
													Type:        schema.TypeList,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_subnets": {
													Required:    true,
													Type:        schema.TypeList,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"denied_countries": {
													Required:    true,
													Type:        schema.TypeList,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_countries": {
													Required:    true,
													Type:        schema.TypeList,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"libinject": {
										Required:    true,
										Type:        schema.TypeSet,
										Description: "",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sqli_effect": {
													Required:    true,
													Type:        schema.TypeString,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
												"xss_effect": {
													Required:    true,
													Type:        schema.TypeString,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
											},
										},
									},
									"body": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"inspection_size_bytes": {
													Optional:    true,
													Default:     131072,
													Type:        schema.TypeInt,
													Description: "",
												},
												"skip": {
													Optional:    true,
													Default:     false,
													Type:        schema.TypeBool,
													Description: "",
												},
											},
										},
									},
									"intel_gathering": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bruteforce_enabled": {
													Optional:    true,
													Default:     true,
													Type:        schema.TypeBool,
													Description: "",
												},
												"track_errors_enabled": {
													Optional:    true,
													Default:     true,
													Type:        schema.TypeBool,
													Description: "",
												},
												"info_leakage_effect": {
													Optional:    true,
													Default:     "alert",
													Type:        schema.TypeString,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
												"remove_fingerprints_enabled": {
													Optional:    true,
													Default:     true,
													Type:        schema.TypeBool,
													Description: "",
												},
											},
										},
									},
									"malicious_upload": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"effect": {
													Optional:    true,
													Default:     "disable",
													Type:        schema.TypeString,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"ban",
															"alert",
															"prevent",
															"disable",
														},
														false,
													),
												},
												"allowed_file_types": {
													Required:    true,
													Type:        schema.TypeList,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_extensions": {
													Required:    true,
													Type:        schema.TypeList,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"attack_tools_effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"ban",
												"alert",
												"prevent",
												"disable",
											},
											false,
										),
									},
									"shellshock_effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"ban",
												"alert",
												"prevent",
												"disable",
											},
											false,
										),
									},
									"malformed_req_effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"ban",
												"alert",
												"prevent",
												"disable",
											},
											false,
										),
									},
									"cmdi_effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"ban",
												"alert",
												"prevent",
												"disable",
											},
											false,
										),
									},
									"lfi_effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"ban",
												"alert",
												"prevent",
												"disable",
											},
											false,
										),
									},
									"code_injection_effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"ban",
												"alert",
												"prevent",
												"disable",
											},
											false,
										),
									},
									"csrf_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
									"clickjacking_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
									"header_specs": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allow": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "",
												},
												"required": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "",
												},
												"effect": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														[]string{
															"alert",
															"prevent",
														},
														false,
													),
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "",
												},
												"values": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func parseWaasContainer(d *schema.ResourceData) *sdk.Waas {
	waasSpec := sdk.Waas{}
	waasSpec.MaxPort = d.Get("max_port").(int)
	waasSpec.MinPort = d.Get("min_port").(int)
	rules := d.Get("rules").([]interface{})

	for _, i := range rules {
		rule := i.(map[string]interface{})
		resources := rule["resources"].(*schema.Set).List()[0].(map[string]interface{})
		applicationsSpec := rule["applications_spec"].([]interface{})

		ruleObject := sdk.Rule{
			Name: rule["name"].(string),
			Resources: sdk.Resources{
				Hosts:      cast.ToStringSlice(resources["hosts"]),
				Images:     cast.ToStringSlice(resources["images"]),
				Labels:     cast.ToStringSlice(resources["labels"]),
				Containers: cast.ToStringSlice(resources["containers"]),
				Namespaces: cast.ToStringSlice(resources["namespaces"]),
				AccountIDs: cast.ToStringSlice(resources["account_ids"]),
			},
		}

		for _, j := range applicationsSpec {
			applicationSpec := j.(map[string]interface{})

			certificate := fetchOptionalMapFromSetParam(applicationSpec, "certificate")
			libInject := fetchOptionalMapFromSetParam(applicationSpec, "libinject")
			intelGathering := fetchOptionalMapFromSetParam(applicationSpec, "intel_gathering")
			networkControls := fetchOptionalMapFromSetParam(applicationSpec, "network_controls")
			apiSpec := fetchOptionalMapFromSetParam(applicationSpec, "api_spec")
			maliciousUpload := fetchOptionalMapFromSetParam(applicationSpec, "malicious_upload")
			body := fetchOptionalMapFromSetParam(applicationSpec, "body")
			headers := fetchOptionalMapFromSetParam(applicationSpec, "header_specs")

			var headersSpecs []sdk.HeaderSpec
			for _, headerSpec := range headers {
				h := headerSpec.(map[string]interface{})
				headersSpecs = append(headersSpecs,
					sdk.HeaderSpec{
						Allow:    h["allow"].(bool),
						Required: h["required"].(bool),
						Effect:   h["effect"].(string),
						Name:     h["name"].(string),
						Values:   cast.ToStringSlice(h["values"]),
					})
			}

			var endpoints []sdk.Endpoint
			for _, endpoint := range apiSpec["endpoints"].(*schema.Set).List() {
				e := endpoint.(map[string]interface{})
				endpoints = append(endpoints,
					sdk.Endpoint{
						ExposedPort:  e["exposed_port"].(int),
						InternalPort: e["internal_port"].(int),
						Host:         e["host"].(string),
						BasePath:     e["base_path"].(string),
						TLS:          e["tls"].(bool),
						HTTP2:        e["http2"].(bool),
					})
			}

			ruleObject.ApplicationsSpec = append(
				ruleObject.ApplicationsSpec,
				sdk.ApplicationSpec{
					Certificate: sdk.Certificate{
						Encrypted: func() string {
							if certificate["encrypted"] == nil {
								return ""
							}

							return certificate["encrypted"].(string)
						}(),
					},

					APISpec: sdk.ApiSpec{
						Endpoints: endpoints,
						Effect:    "disable",
					},
					NetworkControls: sdk.NetworkControls{
						AllowedSubnets:           cast.ToStringSlice(networkControls["allowed_subnets"]),
						AdvancedProtectionEffect: networkControls["advanced_protection_effect"].(string),
						DeniedSubnetsEffect:      networkControls["denied_subnets_effect"].(string),
						DeniedSubnets:            cast.ToStringSlice(networkControls["denied_subnets"]),
						DeniedCountriesEffect:    networkControls["denied_countries_effect"].(string),
						DeniedCountries:          cast.ToStringSlice(networkControls["denied_countries"]),
						AllowedCountriesEffect:   networkControls["allowed_countries_effect"].(string),
						AllowedCountries:         cast.ToStringSlice(networkControls["allowed_countries"]),
					},
					Libinject: sdk.LibInject{
						SqliEffect: libInject["sqli_effect"].(string),
						XSSEffect:  libInject["xss_effect"].(string),
					},
					Body: sdk.Body{
						Skip:                body["skip"].(bool),
						InspectionSizeBytes: body["inspection_size_bytes"].(int),
					},
					IntelGathering: sdk.IntelGathering{
						BruteforceEnabled:         intelGathering["bruteforce_enabled"].(bool),
						TrackErrorsEnabled:        intelGathering["track_errors_enabled"].(bool),
						InfoLeakageEffect:         intelGathering["info_leakage_effect"].(string),
						RemoveFingerprintsEnabled: intelGathering["remove_fingerprints_enabled"].(bool),
					},
					MaliciousUpload: sdk.MaliciousUpload{
						Effect:            maliciousUpload["effect"].(string),
						AllowedFileTypes:  cast.ToStringSlice(maliciousUpload["allowed_file_types"]),
						AllowedExtensions: cast.ToStringSlice(maliciousUpload["allowed_extensions"]),
					},
					AttackToolsEffect:   applicationSpec["attack_tools_effect"].(string),
					ShellshockEffect:    applicationSpec["shellshock_effect"].(string),
					MalformedReqEffect:  applicationSpec["malformed_req_effect"].(string),
					CmdiEffect:          applicationSpec["cmdi_effect"].(string),
					LfiEffect:           applicationSpec["lfi_effect"].(string),
					CodeInjectionEffect: applicationSpec["code_injection_effect"].(string),
					CsrfEnabled:         applicationSpec["csrf_enabled"].(bool),
					ClickjackingEnabled: applicationSpec["clickjacking_enabled"].(bool),
					HeaderSpecs:         headersSpecs,
				},
			)

		}

		waasSpec.Rules = append(
			waasSpec.Rules,
			ruleObject,
		)
	}

	return &waasSpec
}

func saveWaasContainer(d *schema.ResourceData, waas *sdk.Waas) error {
	waasRules := make([]interface{}, 0, len(waas.Rules))

	for _, i := range waas.Rules {
		var applicationsSpec []map[string]interface{}

		for _, applicationSpec := range i.ApplicationsSpec {

			var endpoints []map[string]interface{}
			for _, endpoint := range applicationSpec.APISpec.Endpoints {
				endpoints = append(endpoints, map[string]interface{}{
					"exposed_port":  endpoint.ExposedPort,
					"internal_port": endpoint.InternalPort,
					"host":          endpoint.Host,
					"base_path":     endpoint.BasePath,
					"tls":           endpoint.TLS,
					"http2":         endpoint.HTTP2,
				})
			}

			var headerSpecs []map[string]interface{}

			for _, headerSpec := range applicationSpec.HeaderSpecs {
				headerSpecs = append(headerSpecs, map[string]interface{}{
					"allow":    headerSpec.Allow,
					"name":     headerSpec.Name,
					"values":   headerSpec.Values,
					"required": headerSpec.Required,
					"effect":   headerSpec.Effect,
				})
			}

			applicationsSpec = append(applicationsSpec, map[string]interface{}{
				"certificate": []map[string]interface{}{
					{
						"encrypted": applicationSpec.Certificate.Encrypted,
					},
				},

				"api_spec": []map[string]interface{}{
					{
						"endpoints": endpoints,
					},
				},
				"network_controls": []map[string]interface{}{
					{
						"advanced_protection_effect": applicationSpec.NetworkControls.AdvancedProtectionEffect,
						"denied_subnets_effect":      applicationSpec.NetworkControls.DeniedSubnetsEffect,
						"denied_countries_effect":    applicationSpec.NetworkControls.DeniedCountriesEffect,
						"allowed_countries_effect":   applicationSpec.NetworkControls.AllowedCountriesEffect,
						"denied_subnets":             applicationSpec.NetworkControls.DeniedSubnets,
						"allowed_subnets":            applicationSpec.NetworkControls.AllowedSubnets,
						"denied_countries":           applicationSpec.NetworkControls.DeniedCountries,
						"allowed_countries":          applicationSpec.NetworkControls.AllowedCountries,
					},
				},
				"libinject": []map[string]interface{}{
					{
						"sqli_effect": applicationSpec.Libinject.SqliEffect,
						"xss_effect":  applicationSpec.Libinject.XSSEffect,
					},
				},
				"body": []map[string]interface{}{
					{
						"skip":                  applicationSpec.Body.Skip,
						"inspection_size_bytes": applicationSpec.Body.InspectionSizeBytes,
					},
				},
				"intel_gathering": []map[string]interface{}{
					{
						"bruteforce_enabled":          applicationSpec.IntelGathering.BruteforceEnabled,
						"track_errors_enabled":        applicationSpec.IntelGathering.TrackErrorsEnabled,
						"info_leakage_effect":         applicationSpec.IntelGathering.InfoLeakageEffect,
						"remove_fingerprints_enabled": applicationSpec.IntelGathering.RemoveFingerprintsEnabled,
					},
				},
				"malicious_upload": []map[string]interface{}{
					{
						"effect":             applicationSpec.MaliciousUpload.Effect,
						"allowed_file_types": applicationSpec.MaliciousUpload.AllowedFileTypes,
						"allowed_extensions": applicationSpec.MaliciousUpload.AllowedExtensions,
					},
				},
				"attack_tools_effect":   applicationSpec.AttackToolsEffect,
				"shellshock_effect":     applicationSpec.ShellshockEffect,
				"malformed_req_effect":  applicationSpec.MalformedReqEffect,
				"cmdi_effect":           applicationSpec.CmdiEffect,
				"lfi_effect":            applicationSpec.LfiEffect,
				"code_injection_effect": applicationSpec.CodeInjectionEffect,
				"csrf_enabled":          applicationSpec.CsrfEnabled,
				"clickjacking_enabled":  applicationSpec.ClickjackingEnabled,
				"header_specs":          headerSpecs,
			})
		}

		waasRules = append(
			waasRules,
			map[string]interface{}{
				"name": i.Name,
				"resources": []map[string]interface{}{
					{
						"hosts":       i.Resources.Hosts,
						"images":      i.Resources.Images,
						"labels":      i.Resources.Labels,
						"containers":  i.Resources.Containers,
						"namespaces":  i.Resources.Namespaces,
						"account_ids": i.Resources.AccountIDs,
					},
				},
				"applications_spec": applicationsSpec,
			})
	}

	d.SetId("cnaf")

	err := d.Set("rules", waasRules)
	if err != nil {
		log.Printf("[ERROR] rules caused by: %s", err)
		return err
	}

	errMaxPort := d.Set("max_port", waas.MaxPort)
	if errMaxPort != nil {
		log.Printf("[ERROR] max_port caused by: %s", errMaxPort)
		return errMaxPort
	}

	errMinPort := d.Set("min_port", waas.MinPort)
	if errMinPort != nil {
		log.Printf("[ERROR] min_port caused by: %s", errMinPort)
		return errMinPort
	}

	return nil
}

func createWaasContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.SetWaasRules(parseWaasContainer(d))
	if err != nil {
		return err
	}

	return readWaasContainer(d, meta)
}

func readWaasContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	waasRules, err := client.GetWaasRules()
	if err != nil {
		return err
	}

	return saveWaasContainer(d, waasRules)
}

func deleteWaasContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.SetRegistries(&sdk.RegistrySpecifications{})
	if err != nil {
		return err
	}

	return nil
}
