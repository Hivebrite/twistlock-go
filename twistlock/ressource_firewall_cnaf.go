package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/spf13/cast"
	"log"
)

// TODO(This resource create update correctly on prismacloud but at each run the change is trigger, we will investigate later on that)
func resourceFirewallCnaf() *schema.Resource {
	return &schema.Resource{
		Create: createFirewallCnaf,
		Read:   readFirewallCnaf,
		Update: createFirewallCnaf,
		Delete: deleteFirewallCnaf,

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
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"effect": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "Contains either the registry name (e.g., gcr.io) or url (e.g., https://gcr.io)",
						},
						"black_list": {
							Required:    true,
							Type:        schema.TypeSet,
							Description: "Days of week",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"advanced_protection": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "Day",
									},
									"subnets": {
										Required:    true,
										Type:        schema.TypeList,
										Description: "Offset",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"white_list_subnets": {
							Required:    true,
							Type:        schema.TypeList,
							Description: "Offset",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"libinject": {
							Required:    true,
							Type:        schema.TypeSet,
							Description: "Days of week",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sqli_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "Day",
									},
									"xss_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "Offset",
									},
								},
							},
						},
						"headers": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Days of week",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"spec": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Day",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allow": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Day",
												},
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Day",
												},
												"values": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Day",
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
						"resources": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "Days of week",
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
						"certificate": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Days of week",
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
						"csrf_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"click_jacking_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"attack_tools_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
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
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
									"dir_traversal_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
									"track_errors_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
									"info_leakage_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
									"remove_fingerprints_enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
								},
							},
						},
						"shellshock_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"malformed_req_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"malicious_upload": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
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
						"port_maps": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exposed": {
										Required:    true,
										Type:        schema.TypeInt,
										Description: "",
									},
									"internal": {
										Required:    true,
										Type:        schema.TypeInt,
										Description: "",
									},
									"tls": {
										Required:    true,
										Type:        schema.TypeBool,
										Description: "",
									},
								},
							},
						},
						"cmdi_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"lfi_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
						"code_injection_enabled": {
							Required:    true,
							Type:        schema.TypeBool,
							Description: "",
						},
					},
				},
			},
		},
	}
}

func parseFirewallCnaf(d *schema.ResourceData) *sdk.Cnaf {
	cnafSpec := sdk.Cnaf{}
	cnafSpec.MaxPort = d.Get("max_port").(int)
	cnafSpec.MinPort = d.Get("min_port").(int)
	rules := d.Get("rules").(*schema.Set).List()

	for _, i := range rules {
		rule := i.(map[string]interface{})
		blackList := rule["black_list"].(*schema.Set).List()[0].(map[string]interface{})
		libInject := rule["libinject"].(*schema.Set).List()[0].(map[string]interface{})
		resources := rule["resources"].(*schema.Set).List()[0].(map[string]interface{})
		intelGathering := rule["intel_gathering"].(*schema.Set).List()[0].(map[string]interface{})
		headers := fetchOptionalMapFRomSetParam(rule, "headers")
		certificate := fetchOptionalMapFRomSetParam(rule, "certificate")
		maliciousUpload := fetchOptionalMapFRomSetParam(rule, "malicious_upload")

		var headersSpecs []sdk.Specs
		for _, headerSpec := range cast.ToSlice(headers["spec"]) {
			h := headerSpec.(map[string]interface{})
			headersSpecs = append(headersSpecs,
				sdk.Specs{
					Allow:  h["allow"].(bool),
					Name:   h["name"].(string),
					Values: cast.ToStringSlice(h["values"]),
				})
		}

		var portMaps []sdk.PortMaps
		for _, portMap := range rule["port_maps"].(*schema.Set).List() {
			p := portMap.(map[string]interface{})
			portMaps = append(portMaps,
				sdk.PortMaps{
					Exposed:  p["exposed"].(int),
					Internal: p["internal"].(int),
					TLS:      p["tls"].(bool),
				})
		}

		cnafSpec.Rules = append(
			cnafSpec.Rules,
			sdk.Rules{
				Name:   rule["name"].(string),
				Effect: rule["effect"].(string),
				Blacklist: sdk.Blacklist{
					AdvancedProtection: blackList["advanced_protection"].(bool),
					Subnets:            cast.ToStringSlice(blackList["subnets"]),
				},
				WhitelistSubnets: cast.ToStringSlice(rule["subnets"]),
				Libinject: sdk.Libinject{
					SqliEnabled: libInject["sqli_enabled"].(bool),
					XSSEnabled:  libInject["xss_enabled"].(bool),
				},
				Headers: sdk.Headers{
					Specs: headersSpecs,
				},
				Resources: sdk.Resources{
					Hosts:      cast.ToStringSlice(resources["hosts"]),
					Images:     cast.ToStringSlice(resources["images"]),
					Labels:     cast.ToStringSlice(resources["labels"]),
					Containers: cast.ToStringSlice(resources["containers"]),
					Namespaces: cast.ToStringSlice(resources["namespaces"]),
					AccountIDs: cast.ToStringSlice(resources["account_ids"]),
				},
				Certificate: sdk.Certificate{
					Encrypted: func() string {
						if certificate["encrypted"] == nil {
							return ""
						}

						return certificate["encrypted"].(string)
					}(),
				},
				CsrfEnabled:         rule["csrf_enabled"].(bool),
				ClickjackingEnabled: rule["click_jacking_enabled"].(bool),
				AttackToolsEnabled:  rule["attack_tools_enabled"].(bool),
				IntelGathering: sdk.IntelGathering{
					BruteforceEnabled:         intelGathering["bruteforce_enabled"].(bool),
					DirTraversalEnabled:       intelGathering["dir_traversal_enabled"].(bool),
					TrackErrorsEnabled:        intelGathering["track_errors_enabled"].(bool),
					InfoLeakageEnabled:        intelGathering["info_leakage_enabled"].(bool),
					RemoveFingerprintsEnabled: intelGathering["remove_fingerprints_enabled"].(bool),
				},
				ShellshockEnabled:   rule["shellshock_enabled"].(bool),
				MalformedReqEnabled: rule["malformed_req_enabled"].(bool),
				MaliciousUpload: sdk.MaliciousUpload{
					Enabled: func() bool {
						if maliciousUpload["enabled"] == nil {
							return false
						}

						return maliciousUpload["enabled"].(bool)
					}(),
					AllowedFileTypes:  cast.ToStringSlice(maliciousUpload["allowed_file_types"]),
					AllowedExtensions: cast.ToStringSlice(maliciousUpload["allowed_extensions"]),
				},
				PortMaps:             portMaps,
				CmdiEnabled:          rule["cmdi_enabled"].(bool),
				LfiEnabled:           rule["lfi_enabled"].(bool),
				CodeInjectionEnabled: rule["code_injection_enabled"].(bool),
			})
	}

	return &cnafSpec
}

func saveFirewallCnaf(d *schema.ResourceData, cnaf *sdk.Cnaf) error {
	cnafRules := make([]interface{}, 0, len(cnaf.Rules))

	for _, i := range cnaf.Rules {
		var headerSpecs []map[string]interface{}
		var portMappings []map[string]interface{}

		for _, headerSpec := range i.Headers.Specs {
			headerSpecs = append(headerSpecs, map[string]interface{}{
				"allow":  headerSpec.Allow,
				"name":   headerSpec.Name,
				"values": headerSpec.Values,
			})
		}

		for _, portMapping := range i.PortMaps {
			portMappings = append(portMappings, map[string]interface{}{
				"exposed":  portMapping.Exposed,
				"internal": portMapping.Internal,
				"tls":      portMapping.TLS,
			})
		}

		cnafRules = append(
			cnafRules,
			map[string]interface{}{
				"name":   i.Name,
				"effect": i.Effect,
				"black_list": []map[string]interface{}{
					{
						"advanced_protection": i.Blacklist.AdvancedProtection,
						"subnets":             i.Blacklist.Subnets,
					},
				},
				"white_list_subnets": i.WhitelistSubnets,
				"libinject": []map[string]interface{}{
					{
						"sqli_enabled": i.Libinject.SqliEnabled,
						"xss_enabled":  i.Libinject.XSSEnabled,
					},
				},
				"headers": []map[string]interface{}{
					{
						"spec": headerSpecs,
					},
				},
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
				"certificate": []map[string]interface{}{
					{
						"encrypted": i.Certificate.Encrypted,
					},
				},
				"csrf_enabled":          i.CsrfEnabled,
				"click_jacking_enabled": i.ClickjackingEnabled,
				"attack_tools_enabled":  i.AttackToolsEnabled,
				"intel_gathering": []map[string]interface{}{
					{
						"bruteforce_enabled":          i.IntelGathering.BruteforceEnabled,
						"dir_traversal_enabled":       i.IntelGathering.DirTraversalEnabled,
						"track_errors_enabled":        i.IntelGathering.TrackErrorsEnabled,
						"info_leakage_enabled":        i.IntelGathering.InfoLeakageEnabled,
						"remove_fingerprints_enabled": i.IntelGathering.RemoveFingerprintsEnabled,
					},
				},
				"shellshock_enabled":    i.ShellshockEnabled,
				"malformed_req_enabled": i.MalformedReqEnabled,
				"malicious_upload": []map[string]interface{}{
					{
						"enabled":            i.MaliciousUpload.Enabled,
						"allowed_file_types": i.MaliciousUpload.AllowedFileTypes,
						"allowed_extensions": i.MaliciousUpload.AllowedExtensions,
					},
				},
				"port_maps":              portMappings,
				"cmdi_enabled":           i.CmdiEnabled,
				"lfi_enabled":            i.LfiEnabled,
				"code_injection_enabled": i.CodeInjectionEnabled,
			})
	}

	d.SetId("cnaf")

	err := d.Set("rules", cnafRules)
	if err != nil {
		log.Printf("[ERROR] rules caused by: %s", err)
		return err
	}

	errMaxPort := d.Set("max_port", cnaf.MaxPort)
	if errMaxPort != nil {
		log.Printf("[ERROR] max_port caused by: %s", errMaxPort)
		return errMaxPort
	}

	errMinPort := d.Set("min_port", cnaf.MinPort)
	if errMinPort != nil {
		log.Printf("[ERROR] min_port caused by: %s", errMinPort)
		return errMinPort
	}

	return nil
}

func createFirewallCnaf(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.SetCnafRules(parseFirewallCnaf(d))
	if err != nil {
		return err
	}

	return readFirewallCnaf(d, meta)
}

func readFirewallCnaf(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	cnafRules, err := client.GetCnafRules()
	if err != nil {
		return err
	}

	return saveFirewallCnaf(d, cnafRules)
}

func deleteFirewallCnaf(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := client.SetRegistries(&sdk.RegistrySpecifications{})
	if err != nil {
		return err
	}

	return nil
}
