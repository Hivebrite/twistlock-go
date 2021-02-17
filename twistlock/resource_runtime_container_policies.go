package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/policies"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/spf13/cast"
)

func resourceRuntimeContainerPolicies() *schema.Resource {
	return &schema.Resource{
		Create: createRuntimeContainerPolicies,
		Read:   readRuntimeContainerPolicies,
		Update: createRuntimeContainerPolicies,
		Delete: deleteRuntimeContainerPolicies,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"policy_type": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "",
			},
			"learning_disabled": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
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
						"notes": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "",
						},
						"collections": collectionSchema(),
						"advanced_protection": {
							Optional:    true,
							Default:     false,
							Type:        schema.TypeBool,
							Description: "",
						},
						"kubernetes_enforcement": {
							Optional:    true,
							Default:     false,
							Type:        schema.TypeBool,
							Description: "",
						},
						"cloud_metadata_enforcement": {
							Optional:    true,
							Default:     false,
							Type:        schema.TypeBool,
							Description: "",
						},
						"processes": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "list of URLs to protect",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												policies.EffectBlock,
												policies.EffectAlert,
												policies.EffectDisable,
												policies.EffectPrevent,
											},
											false,
										),
									},
									"blacklist": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelist": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"check_crypto_miners": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
									"skip_modified": {
										Optional:    true,
										Default:     true,
										Type:        schema.TypeBool,
										Description: "",
									},
									"check_lateral_movement": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
									"check_parent_child": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
								},
							},
						},
						"network": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "list of URLs to protect",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												policies.EffectBlock,
												policies.EffectAlert,
												policies.EffectDisable,
											},
											false,
										),
									},
									"blacklist_listening_ports": listOfPortSchema(),
									"whitelist_listening_ports": listOfPortSchema(),
									"blacklist_outbound_ports":  listOfPortSchema(),
									"whitelist_outbound_ports":  listOfPortSchema(),
									"blacklist_ips": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelist_ips": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"skip_modified_proc": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
									"detect_port_scan": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
									"skip_raw_sockets": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
								},
							},
						},
						"dns": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "list of URLs to protect",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												policies.EffectBlock,
												policies.EffectAlert,
												policies.EffectDisable,
												policies.EffectPrevent,
											},
											false,
										),
									},
									"blacklist": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelist": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"filesystem": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "list of URLs to protect",
							MinItems:    1,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												policies.EffectBlock,
												policies.EffectAlert,
												policies.EffectDisable,
												policies.EffectPrevent,
											},
											false,
										),
									},
									"blacklist": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"whitelist": {
										Optional:    true,
										Type:        schema.TypeList,
										Description: "",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"check_new_files": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
									"backdoor_files": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
								},
							},
						},
						"custom_rules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "list of URLs to protect",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"identifier": {
										Required:    true,
										Type:        schema.TypeInt,
										Description: "",
									},
									"action": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"",
												policies.ActionIncident,
												policies.ActionAudit,
											},
											false,
										),
									},
									"effect": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
										ValidateFunc: validation.StringInSlice(
											[]string{
												policies.EffectAllow,
												policies.EffectAlert,
												policies.EffectPrevent,
												policies.EffectBlock,
											},
											false,
										),
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

func parseRuntimeContainerPolicies(d *schema.ResourceData) *policies.Runtime {
	policiesObject := policies.Runtime{
		ID: "containerRuntime",
	}

	rules := d.Get("rules").([]interface{})

	for _, i := range rules {
		rule := i.(map[string]interface{})
		processes := fetchOptionalMapFromSetParam(rule, "processes")
		dns := fetchOptionalMapFromSetParam(rule, "dns")
		filesystem := fetchOptionalMapFromSetParam(rule, "filesystem")

		var customRules []policies.CustomRules

		for _, i := range rule["custom_rules"].([]interface{}) {
			customRule := i.(map[string]interface{})

			customRules = append(customRules, policies.CustomRules{
				ID:     customRule["identifier"].(int),
				Action: customRule["action"].(string),
				Effect: customRule["effect"].(string),
			})
		}

		ruleObject := policies.RuntimeRules{
			Name:                     rule["name"].(string),
			Notes:                    rule["notes"].(string),
			AdvancedProtection:       rule["advanced_protection"].(bool),
			KubernetesEnforcement:    rule["kubernetes_enforcement"].(bool),
			CloudMetadataEnforcement: rule["cloud_metadata_enforcement"].(bool),
			Collections:              parseCollections(rule["collections"].(*schema.Set).List()),
			CustomRules:              customRules,
			Network:                  *networkFromRule(rule),
			Processes: policies.Processes{
				Effect:               processes["effect"].(string),
				Blacklist:            cast.ToStringSlice(processes["blacklist"]),
				Whitelist:            cast.ToStringSlice(processes["whitelist"]),
				SkipModified:         processes["skip_modified"].(bool),
				CheckCryptoMiners:    processes["check_crypto_miners"].(bool),
				CheckLateralMovement: processes["check_lateral_movement"].(bool),
				CheckParentChild:     processes["check_parent_child"].(bool),
			},
			DNS: policies.DNS{
				Effect:    dns["effect"].(string),
				Blacklist: cast.ToStringSlice(dns["blacklist"]),
				Whitelist: cast.ToStringSlice(dns["whitelist"]),
			},
			Filesystem: policies.Filesystem{
				Effect:        filesystem["effect"].(string),
				Blacklist:     cast.ToStringSlice(filesystem["blacklist"]),
				Whitelist:     cast.ToStringSlice(filesystem["whitelist"]),
				CheckNewFiles: filesystem["check_new_files"].(bool),
				BackdoorFiles: filesystem["backdoor_files"].(bool),
			},
		}

		policiesObject.Rules = append(
			policiesObject.Rules,
			ruleObject,
		)
	}

	return &policiesObject
}

func saveRuntimeContainerPolicies(d *schema.ResourceData, policiesObject *policies.Runtime) error {
	rules := make([]interface{}, 0, len(policiesObject.Rules))

	for _, i := range policiesObject.Rules {

		var customRules []map[string]interface{}

		for _, j := range i.CustomRules {
			customRules = append(customRules,
				map[string]interface{}{
					"effect":     j.Effect,
					"identifier": j.ID,
					"action":     j.Action,
				},
			)
		}

		rules = append(
			rules,
			map[string]interface{}{
				"name":                       i.Name,
				"notes":                      i.Notes,
				"advanced_protection":        i.AdvancedProtection,
				"kubernetes_enforcement":     i.KubernetesEnforcement,
				"cloud_metadata_enforcement": i.CloudMetadataEnforcement,
				"collections":                collectionSliceToInterface(i.Collections),
				"custom_rules":               customRules,
				"processes": []map[string]interface{}{
					{
						"effect":                 i.Processes.Effect,
						"blacklist":              i.Processes.Blacklist,
						"whitelist":              i.Processes.Whitelist,
						"skip_modified":          i.Processes.SkipModified,
						"check_crypto_miners":    i.Processes.CheckCryptoMiners,
						"check_lateral_movement": i.Processes.CheckLateralMovement,
						"check_parent_child":     i.Processes.CheckParentChild,
					},
				},
				"network": []map[string]interface{}{
					*networkFromObject(i.Network),
				},
				"dns": []map[string]interface{}{
					{
						"effect":    i.DNS.Effect,
						"blacklist": i.DNS.Blacklist,
						"whitelist": i.DNS.Whitelist,
					},
				},
				"filesystem": []map[string]interface{}{
					{
						"effect":          i.Filesystem.Effect,
						"blacklist":       i.Filesystem.Blacklist,
						"whitelist":       i.Filesystem.Whitelist,
						"check_new_files": i.Filesystem.CheckNewFiles,
						"backdoor_files":  i.Filesystem.BackdoorFiles,
					},
				},
			},
		)
	}

	d.SetId("containerRuntime")

	err := d.Set("rules", rules)
	if err != nil {
		log.Printf("[ERROR] rules caused by: %s", err)
		return err
	}

	err = d.Set("learning_disabled", policiesObject.LearningDisabled)
	if err != nil {
		log.Printf("[ERROR] learning_disabled caused by: %s", err)
		return err
	}
	return nil
}

func createRuntimeContainerPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := policies.SetRuntime(*client, parseRuntimeContainerPolicies(d))
	if err != nil {
		return err
	}

	return readRuntimeContainerPolicies(d, meta)
}

func readRuntimeContainerPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	policies, err := policies.GetRuntime(*client)
	if err != nil {
		return err
	}

	return saveRuntimeContainerPolicies(d, policies)
}

func deleteRuntimeContainerPolicies(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := policies.SetImages(*client, &policies.Images{
		ID: "containerRuntime",
	})
	if err != nil {
		return err
	}

	return nil
}

func networkFromRule(rule map[string]interface{}) *policies.Network {
	network := rule["network"].(*schema.Set).List()[0].(map[string]interface{})
	var blacklistListeningPorts []policies.Ports
	var blacklistOutboundPorts []policies.Ports
	var whitelistListeningPorts []policies.Ports
	var whitelistOutboundPorts []policies.Ports

	for _, i := range network["blacklist_listening_ports"].([]interface{}) {
		blacklistListeningPorts = append(blacklistListeningPorts, *portFromNetwork(i.(map[string]interface{})))
	}

	for _, i := range network["blacklist_outbound_ports"].([]interface{}) {
		blacklistOutboundPorts = append(blacklistOutboundPorts, *portFromNetwork(i.(map[string]interface{})))
	}

	for _, i := range network["whitelist_listening_ports"].([]interface{}) {
		whitelistListeningPorts = append(whitelistListeningPorts, *portFromNetwork(i.(map[string]interface{})))
	}

	for _, i := range network["whitelist_outbound_ports"].([]interface{}) {
		whitelistOutboundPorts = append(whitelistOutboundPorts, *portFromNetwork(i.(map[string]interface{})))
	}

	return &policies.Network{
		Effect:                  network["effect"].(string),
		SkipModifiedProc:        network["skip_modified_proc"].(bool),
		DetectPortScan:          network["detect_port_scan"].(bool),
		SkipRawSockets:          network["skip_raw_sockets"].(bool),
		BlacklistIPs:            cast.ToStringSlice(network["blacklist_ips"]),
		WhitelistIPs:            cast.ToStringSlice(network["whitelist_ips"]),
		BlacklistListeningPorts: blacklistListeningPorts,
		BlacklistOutboundPorts:  blacklistOutboundPorts,
		WhitelistListeningPorts: whitelistListeningPorts,
		WhitelistOutboundPorts:  whitelistOutboundPorts,
	}
}

func networkFromObject(networkObject policies.Network) *map[string]interface{} {
	var blacklistListeningPorts []map[string]interface{}
	var blacklistOutboundPorts []map[string]interface{}
	var whitelistListeningPorts []map[string]interface{}
	var whitelistOutboundPorts []map[string]interface{}

	for _, i := range networkObject.BlacklistListeningPorts {
		blacklistListeningPorts = append(blacklistListeningPorts, *portFromObject(i))
	}

	for _, i := range networkObject.BlacklistOutboundPorts {
		blacklistOutboundPorts = append(blacklistOutboundPorts, *portFromObject(i))
	}

	for _, i := range networkObject.WhitelistListeningPorts {
		whitelistListeningPorts = append(whitelistListeningPorts, *portFromObject(i))
	}

	for _, i := range networkObject.WhitelistOutboundPorts {
		whitelistOutboundPorts = append(whitelistOutboundPorts, *portFromObject(i))
	}

	return &map[string]interface{}{
		"effect":                    networkObject.Effect,
		"blacklist_ips":             networkObject.BlacklistIPs,
		"whitelist_ips":             networkObject.WhitelistIPs,
		"skip_modified_proc":        networkObject.SkipModifiedProc,
		"detect_port_scan":          networkObject.DetectPortScan,
		"skip_raw_sockets":          networkObject.SkipRawSockets,
		"blacklist_listening_ports": blacklistListeningPorts,
		"whitelist_listening_ports": whitelistListeningPorts,
		"blacklist_outbound_ports":  blacklistOutboundPorts,
		"whitelist_outbound_ports":  whitelistOutboundPorts,
	}
}

func portFromNetwork(port map[string]interface{}) *policies.Ports {
	return &policies.Ports{
		Start: port["start"].(int),
		End:   port["end"].(int),
	}
}

func portFromObject(port policies.Ports) *map[string]interface{} {
	return &map[string]interface{}{
		"start": port.Start,
		"end":   port.End,
	}
}
