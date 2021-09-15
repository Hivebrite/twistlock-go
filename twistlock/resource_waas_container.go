package twistlock

import (
	"log"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/waas"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/spf13/cast"
)

var waaf_effects = []string{Ban, Alert, Prevent, Disable}
var http_effects = []string{Alert, Prevent}

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
						"collections": collectionSchema(),
						"read_timeout_seconds": {
							Optional:    true,
							Default:     5,
							Type:        schema.TypeInt,
							Description: "",
						},
						"applications_spec": {
							Required:    true,
							Type:        schema.TypeList,
							Description: "firewalls rules to be applied to the resources",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ban_duration_minutes": {
										Optional:    true,
										Type:        schema.TypeInt,
										Description: "",
										Default:     5,
									},
									"app_id": {
										Required:    true,
										Type:        schema.TypeString,
										Description: "",
									},
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
									"dos_config": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "DoS Rules",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Optional:    true,
													Type:        schema.TypeBool,
													Description: "",
													Default:     false,
												},
												"alert": collectionDosConfigEffect(),
												"ban":   collectionDosConfigEffect(),
												"excluded_network_lists": {
													Optional:    true,
													Type:        schema.TypeList,
													Description: "",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"track_session": {
													Optional:    true,
													Type:        schema.TypeBool,
													Description: "",
												},
												"match_conditions": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"methods": {
																Optional:    true,
																Type:        schema.TypeList,
																Description: "",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"file_types": {
																Optional:    true,
																Type:        schema.TypeList,
																Description: "",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"response_code_ranges": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"start": {
																			Required:    true,
																			Type:        schema.TypeInt,
																			Description: "",
																		},
																		"end": {
																			Optional:    true,
																			Type:        schema.TypeInt,
																			Description: "",
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
												"skip_learning": {
													Optional:    true,
													Default:     false,
													Type:        schema.TypeBool,
													Description: "",
												},
												"description": {
													Optional:    true,
													Default:     "",
													Type:        schema.TypeString,
													Description: "",
												},
												"fallback_effect": {
													Optional:    true,
													Default:     Disable,
													Type:        schema.TypeString,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														waaf_effects,
														false),
												},
												"effect": {
													Optional:    true,
													Default:     Disable,
													Type:        schema.TypeString,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														waaf_effects,
														false),
												},
												"paths": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"path": {
																Required:    true,
																Type:        schema.TypeString,
																Description: "",
															},
															"methods": {
																Required:    true,
																Type:        schema.TypeList,
																Description: "",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"method": {
																			Optional:    true,
																			Type:        schema.TypeString,
																			Description: "",
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
									"bot_protection_spec": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "Bots description",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_defined_bots": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Optional:    true,
																Type:        schema.TypeString,
																Description: "",
															},
															"header_name": {
																Optional:    true,
																Type:        schema.TypeString,
																Description: "",
															},
															"header_values": {
																Optional:    true,
																Type:        schema.TypeList,
																Description: "",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"subnets": {
																Optional:    true,
																Type:        schema.TypeList,
																Description: "",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"effect": {
																Optional:    true,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	[]string{
																		"allow",
																		"alert",
																		"prevent",
																		"ban",
																	},
																	false,
																),
															},
														},
													},
												},
												"known_bot_protections_spec": {
													Type:        schema.TypeSet,
													Required:    true,
													Description: "",
													MinItems:    1,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"search_engine_crawlers": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"business_analytics": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"educational": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"news": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"financial": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"content_feed_clients": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"archiving": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"career_search": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"media_search": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
														},
													},
												},
												"unknown_bot_protections_spec": {
													Type:        schema.TypeSet,
													Required:    true,
													Description: "",
													MinItems:    1,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"generic": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"web_automation_tools": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"web_scrapers": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"api_libraries": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"http_libraries": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"bot_impersonation": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"browser_impersonation": {
																Optional:    true,
																Default:     Disable,
																Type:        schema.TypeString,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
															"request_anomalies": {
																Required:    true,
																Type:        schema.TypeSet,
																Description: "",
																MinItems:    1,
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"threshold": {
																			Optional:    true,
																			Type:        schema.TypeInt,
																			Description: "",
																			ValidateFunc: validation.IntInSlice(
																				[]int{
																					0, 3, 6, 9,
																				},
																			),
																		},
																		"effect": {
																			Optional:    true,
																			Type:        schema.TypeString,
																			Description: "",
																			Default:     Disable,
																			ValidateFunc: validation.StringInSlice(
																				waaf_effects, false,
																			),
																		},
																	},
																},
															},
														},
													},
												},
												"session_validation": {
													Optional:    true,
													Type:        schema.TypeString,
													Default:     Disable,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														waaf_effects,
														false,
													),
												},
												"interstitial_page": {
													Optional:    true,
													Type:        schema.TypeBool,
													Description: "",
												},
												"js_injection_spec": {
													Type:        schema.TypeSet,
													Required:    true,
													Description: "",
													MinItems:    1,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enabled": {
																Optional:    true,
																Default:     false,
																Type:        schema.TypeBool,
																Description: "",
															},
															"timeout_effect": {
																Optional:    true,
																Type:        schema.TypeString,
																Default:     Disable,
																Description: "",
																ValidateFunc: validation.StringInSlice(
																	waaf_effects,
																	false,
																),
															},
														},
													},
												},
												"recaptcha_spec": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "",
													MinItems:    1,
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"enabled": {
																Optional:    true,
																Default:     false,
																Type:        schema.TypeBool,
																Description: "",
															},
															"site_key": {
																Optional:    true,
																Type:        schema.TypeString,
																Default:     Disable,
																Description: "",
															},
															"type": {
																Optional:    true,
																Type:        schema.TypeString,
																Default:     Disable,
																Description: "",
															},
															"all_sessions": {
																Optional:    true,
																Default:     false,
																Type:        schema.TypeBool,
																Description: "",
															},
															"success_expiration_hours": {
																Optional:    true,
																Default:     24,
																Type:        schema.TypeInt,
																Description: "",
															},
															"secret_key": {
																Type:        schema.TypeSet,
																Optional:    true,
																Computed:    true,
																Description: "recaptcha key",
																MinItems:    1,
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"encrypted": {
																			Required: true,
																			Type:     schema.TypeString,
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
													Default:     Alert,
													ValidateFunc: validation.StringInSlice(
														waaf_effects,
														false,
													),
												},
												"subnets":   networkControlsEffect(),
												"countries": networkControlsEffect(),
												"exception_subnets": {
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
														http_effects,
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
									"intel_gathering": {
										Type:        schema.TypeSet,
										Required:    true,
										Description: "",
										MinItems:    1,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"info_leakage_effect": {
													Optional:    true,
													Default:     "alert",
													Type:        schema.TypeString,
													Description: "",
													ValidateFunc: validation.StringInSlice(
														waaf_effects,
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
														waaf_effects,
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
									"session_cookie_enabled": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
									"session_cookie_ban": {
										Optional:    true,
										Default:     false,
										Type:        schema.TypeBool,
										Description: "",
									},
									"sqli":           applicationSpecEffectsSchema(),
									"xss":            applicationSpecEffectsSchema(),
									"attack_tools":   applicationSpecEffectsSchema(),
									"shellshock":     applicationSpecEffectsSchema(),
									"malformed_req":  applicationSpecEffectsSchema(),
									"cmdi":           applicationSpecEffectsSchema(),
									"lfi":            applicationSpecEffectsSchema(),
									"code_injection": applicationSpecEffectsSchema(),
								},
							},
						},
					},
				},
			},
		},
	}
}

func parseWaasContainer(d *schema.ResourceData) *waas.Waas {
	waasSpec := waas.Waas{}
	waasSpec.MaxPort = d.Get("max_port").(int)
	waasSpec.MinPort = d.Get("min_port").(int)
	rules := d.Get("rules").([]interface{})

	for _, i := range rules {
		rule := i.(map[string]interface{})
		applicationsSpec := rule["applications_spec"].([]interface{})

		ruleObject := waas.Rule{
			Name:               rule["name"].(string),
			ReadTimeoutSeconds: rule["read_timeout_seconds"].(int),
		}

		ruleObject.Collections = parseCollections(rule["collections"].(*schema.Set).List())

		for _, j := range applicationsSpec {
			applicationSpec := j.(map[string]interface{})

			certificate := fetchOptionalMapFromSetParam(applicationSpec, "certificate")
			dosConfig := fetchOptionalMapFromSetParam(applicationSpec, "dos_config")
			apiSpec := fetchOptionalMapFromSetParam(applicationSpec, "api_spec")
			botProtectionSpec := fetchOptionalMapFromSetParam(applicationSpec, "bot_protection_spec")
			networkControls := fetchOptionalMapFromSetParam(applicationSpec, "network_controls")
			body := fetchOptionalMapFromSetParam(applicationSpec, "body")
			headers := fetchOptionalMapFromSetParam(applicationSpec, "header_specs")
			intelGathering := fetchOptionalMapFromSetParam(applicationSpec, "intel_gathering")
			maliciousUpload := fetchOptionalMapFromSetParam(applicationSpec, "malicious_upload")
			sqli := fetchOptionalMapFromSetParam(applicationSpec, "sqli")
			xss := fetchOptionalMapFromSetParam(applicationSpec, "xss")
			attackTools := fetchOptionalMapFromSetParam(applicationSpec, "attack_tools")
			shellshock := fetchOptionalMapFromSetParam(applicationSpec, "shellshock")
			malformedReq := fetchOptionalMapFromSetParam(applicationSpec, "malformed_req")
			cmdi := fetchOptionalMapFromSetParam(applicationSpec, "cmdi")
			lfi := fetchOptionalMapFromSetParam(applicationSpec, "lfi")
			codeInjection := fetchOptionalMapFromSetParam(applicationSpec, "code_injection")
			// remoteHostForwarding := fetchOptionalMapFromSetParam(applicationSpec, "remote_host_forwarding")

			var headersSpecs []waas.HeaderSpec
			for _, headerSpec := range headers {
				h := headerSpec.(map[string]interface{})
				headersSpecs = append(headersSpecs,
					waas.HeaderSpec{
						Allow:    h["allow"].(bool),
						Required: h["required"].(bool),
						Effect:   h["effect"].(string),
						Name:     h["name"].(string),
						Values:   cast.ToStringSlice(h["values"]),
					})
			}

			botProtectionObject := waas.BotProtectionSpec{}

			botProtectionObject.InterstitialPage = botProtectionSpec["interstitial_page"].(bool)
			botProtectionObject.SessionValidation = botProtectionSpec["session_validation"].(string)

			jsInjectionSpec := fetchOptionalMapFromSetParam(botProtectionSpec, "js_injection_spec")

			botProtectionObject.JsInjectionSpec = waas.JsInjectionSpec{
				Enabled:       jsInjectionSpec["enabled"].(bool),
				TimeoutEffect: jsInjectionSpec["timeout_effect"].(string),
			}

			knownBotProtection := fetchOptionalMapFromSetParam(botProtectionSpec, "known_bot_protections_spec")
			botProtectionObject.KnownBotProtectionsSpec = waas.KnownBotProtectionsSpec{
				SearchEngineCrawlers: knownBotProtection["search_engine_crawlers"].(string),
				BusinessAnalytics:    knownBotProtection["business_analytics"].(string),
				Educational:          knownBotProtection["educational"].(string),
				News:                 knownBotProtection["news"].(string),
				Financial:            knownBotProtection["financial"].(string),
				ContentFeedClients:   knownBotProtection["content_feed_clients"].(string),
				Archiving:            knownBotProtection["archiving"].(string),
				CareerSearch:         knownBotProtection["career_search"].(string),
				MediaSearch:          knownBotProtection["media_search"].(string),
			}

			unknownBotProtection := fetchOptionalMapFromSetParam(botProtectionSpec, "unknown_bot_protections_spec")
			requestAnomalies := fetchOptionalMapFromSetParam(unknownBotProtection, "request_anomalies")
			recaptchaSpec := fetchOptionalMapFromSetParam(botProtectionSpec, "recaptcha_spec")
			recaptchaSecret := fetchOptionalMapFromSetParam(recaptchaSpec, "secret_key")

			botProtectionObject.UnknownBotProtectionSpec = waas.UnknownBotProtectionSpec{
				Generic:              unknownBotProtection["generic"].(string),
				WebAutomationTools:   unknownBotProtection["web_automation_tools"].(string),
				WebScrapers:          unknownBotProtection["web_scrapers"].(string),
				APILibraries:         unknownBotProtection["api_libraries"].(string),
				HTTPLibraries:        unknownBotProtection["http_libraries"].(string),
				BotImpersonation:     unknownBotProtection["bot_impersonation"].(string),
				BrowserImpersonation: unknownBotProtection["browser_impersonation"].(string),
				RequestAnomalies: waas.RequestAnomalies{
					Threshold: requestAnomalies["threshold"].(int),
					Effect:    requestAnomalies["effect"].(string),
				}}

			var userDefinedBots []waas.UserDefinedBot
			for _, userDefinedBot := range botProtectionSpec["user_defined_bots"].([]interface{}) {
				botSpec := userDefinedBot.(map[string]interface{})
				userDefinedBots = append(userDefinedBots,
					waas.UserDefinedBot{
						Name:         botSpec["name"].(string),
						HeaderName:   botSpec["header_name"].(string),
						HeaderValues: botSpec["header_values"].([]string),
						Subnets:      botSpec["subnets"].([]string),
						Effect:       botSpec["effect"].(string),
					})
			}
			botProtectionObject.UserDefinedBots = userDefinedBots

			botProtectionObject.ReCAPTCHASpec = waas.ReCAPTCHASpec{
				Enabled:                recaptchaSpec["enabled"].(bool),
				SiteKey:                recaptchaSpec["site_key"].(string),
				Type:                   recaptchaSpec["type"].(string),
				AllSessions:            recaptchaSpec["all_sessions"].(bool),
				SuccessExpirationHours: recaptchaSpec["success_expiration_hours"].(int),
				SecretKey: sdk.Secret{
					Encrypted: func() string {
						if recaptchaSecret["encrypted"] == nil {
							return ""
						}

						return recaptchaSecret["encrypted"].(string)
					}(),
				},
			}

			var endpoints []waas.Endpoint
			for _, endpoint := range apiSpec["endpoints"].(*schema.Set).List() {
				e := endpoint.(map[string]interface{})
				endpoints = append(endpoints,
					waas.Endpoint{
						ExposedPort:  e["exposed_port"].(int),
						InternalPort: e["internal_port"].(int),
						Host:         e["host"].(string),
						BasePath:     e["base_path"].(string),
						TLS:          e["tls"].(bool),
						HTTP2:        e["http2"].(bool),
					})
			}

			var paths []waas.Path
			for _, path := range apiSpec["paths"].(*schema.Set).List() {
				p := path.(map[string]interface{})

				var methods []waas.Method
				for _, method := range p["methods"].([]interface{}) {
					m := method.(map[string]interface{})
					methods = append(methods,
						waas.Method{
							Method: m["method"].(string),
						})
				}
				paths = append(paths,
					waas.Path{
						Path:    p["path"].(string),
						Methods: methods,
					})
			}

			ruleObject.ApplicationsSpec = append(
				ruleObject.ApplicationsSpec,
				waas.ApplicationSpec{
					AppID: applicationSpec["app_id"].(string),
					Certificate: sdk.Secret{
						Encrypted: func() string {
							if certificate["encrypted"] == nil {
								return ""
							}

							return certificate["encrypted"].(string)
						}(),
					},
					APISpec: waas.APISpec{
						Endpoints:      endpoints,
						Paths:          paths,
						Effect:         apiSpec["effect"].(string),
						FallbackEffect: apiSpec["fallback_effect"].(string),
						Description:    apiSpec["description"].(string),
						SkipLearning:   apiSpec["skip_learning"].(bool),
					},
					NetworkControls: waas.NetworkControls{
						AdvancedProtectionEffect: networkControls["advanced_protection_effect"].(string),
						ExceptionSubnets:         cast.ToStringSlice(networkControls["exception_subnets"]),
						Subnets:                  parseNetworkControlsEffect(fetchOptionalMapFromSetParam(networkControls, "subnets")),
						Countries:                parseNetworkControlsEffect(fetchOptionalMapFromSetParam(networkControls, "countries")),
					},
					Sqli:          *applicationSpecEffectsFromInterface(sqli),
					XSS:           *applicationSpecEffectsFromInterface(xss),
					AttackTools:   *applicationSpecEffectsFromInterface(attackTools),
					Shellshock:    *applicationSpecEffectsFromInterface(shellshock),
					MalformedReq:  *applicationSpecEffectsFromInterface(malformedReq),
					Cmdi:          *applicationSpecEffectsFromInterface(cmdi),
					Lfi:           *applicationSpecEffectsFromInterface(lfi),
					CodeInjection: *applicationSpecEffectsFromInterface(codeInjection),
					Body: waas.Body{
						Skip:                body["skip"].(bool),
						InspectionSizeBytes: body["inspection_size_bytes"].(int),
					},
					IntelGathering: waas.IntelGathering{
						InfoLeakageEffect:         intelGathering["info_leakage_effect"].(string),
						RemoveFingerprintsEnabled: intelGathering["remove_fingerprints_enabled"].(bool),
					},
					MaliciousUpload: waas.MaliciousUpload{
						Effect:            maliciousUpload["effect"].(string),
						AllowedFileTypes:  cast.ToStringSlice(maliciousUpload["allowed_file_types"]),
						AllowedExtensions: cast.ToStringSlice(maliciousUpload["allowed_extensions"]),
					},
					CsrfEnabled:          applicationSpec["csrf_enabled"].(bool),
					BanDurationMinutes:   applicationSpec["ban_duration_minutes"].(int),
					ClickjackingEnabled:  applicationSpec["clickjacking_enabled"].(bool),
					SessionCookieEnabled: applicationSpec["session_cookie_enabled"].(bool),
					SessionCookieBan:     applicationSpec["session_cookie_ban"].(bool),
					HeaderSpecs:          headersSpecs,
					DosConfig: waas.DosConfig{
						Enabled:              dosConfig["enabled"].(bool),
						TrackSession:         dosConfig["track_session"].(bool),
						Alert:                parseDosConfigRate(fetchOptionalMapFromSetParam(dosConfig, "alert")),
						Ban:                  parseDosConfigRate(fetchOptionalMapFromSetParam(dosConfig, "ban")),
						ExcludedNetworkLists: cast.ToStringSlice(dosConfig["excluded_network_lists"]),
						MatchConditions:      parseDosConditions(dosConfig["match_conditions"].([]interface{})),
					},
					BotProtectionSpec: botProtectionObject,
					// remoteHostForwarding
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

func parseDosConditions(dosConfig []interface{}) []waas.MatchConditions {

	var dosConditions []waas.MatchConditions

	if len(dosConfig) > 0 {
		for _, matchConditionsInterface := range dosConfig {
			matchConditions := matchConditionsInterface.(map[string]interface{})
			var responseCodeRanges []waas.ResponseCodeRange
			for _, codeRangeInterface := range matchConditions["response_code_ranges"].([]interface{}) {
				codeRange := codeRangeInterface.(map[string]interface{})
				responseCodeRanges = append(responseCodeRanges,
					waas.ResponseCodeRange{
						Start: codeRange["start"].(int),
						End:   codeRange["end"].(int),
					})
			}

			var methods = cast.ToStringSlice(matchConditions["methods"])
			var fileTypes = cast.ToStringSlice(matchConditions["file_types"])

			dosConditions = append(dosConditions,
				waas.MatchConditions{
					Methods:            methods,
					FileTypes:          fileTypes,
					ResponseCodeRanges: responseCodeRanges,
				})
		}
	}

	return dosConditions
}

func parseDosConfigRate(dosConfigRate map[string]interface{}) waas.DosConfigRate {
	configRate := waas.DosConfigRate{}

	if dosConfigRate["burst"] != nil {
		configRate.Burst = dosConfigRate["burst"].(int)
	}

	if dosConfigRate["average"] != nil {
		configRate.Average = dosConfigRate["average"].(int)
	}

	return configRate
}

func parseNetworkControlsEffect(networkControlsEffect map[string]interface{}) waas.NetworkControlsEffect {

	return waas.NetworkControlsEffect{
		Enabled:        networkControlsEffect["enabled"].(bool),
		AllowMode:      networkControlsEffect["allow_mode"].(bool),
		FallbackEffect: networkControlsEffect["fallback_effect"].(string),
		Allow:          cast.ToStringSlice(networkControlsEffect["allow"]),
		Prevent:        cast.ToStringSlice(networkControlsEffect["prevent"]),
		Alert:          cast.ToStringSlice(networkControlsEffect["alert"]),
	}
}

func saveWaasContainer(d *schema.ResourceData, waasObject *waas.Waas) error {
	waasRules := make([]interface{}, 0, len(waasObject.Rules))

	for _, i := range waasObject.Rules {
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

			var dosConditions []map[string]interface{}
			for _, dosCondition := range applicationSpec.DosConfig.MatchConditions {

				var codeRanges []map[string]interface{}
				for _, codeRange := range dosCondition.ResponseCodeRanges {
					codeRanges = append(codeRanges, map[string]interface{}{
						"start": codeRange.Start,
						"end":   codeRange.End,
					})
				}
				dosConditions = append(dosConditions, map[string]interface{}{
					"methods":              dosCondition.Methods,
					"file_types":           dosCondition.FileTypes,
					"response_code_ranges": codeRanges,
				})
			}

			botProtectionSpec := applicationSpec.BotProtectionSpec

			userDefinedBotsObject := botProtectionSpec.UserDefinedBots
			var userDefinedBots []map[string]interface{}
			for _, userDefinedBotObject := range userDefinedBotsObject {
				userDefinedBots = append(userDefinedBots, map[string]interface{}{
					"name":          userDefinedBotObject.Name,
					"header_name":   userDefinedBotObject.HeaderName,
					"header_values": userDefinedBotObject.HeaderValues,
					"subnets":       userDefinedBotObject.Subnets,
					"effect":        userDefinedBotObject.Effect,
				})
			}

			knownBotProtectionsSpec := botProtectionSpec.KnownBotProtectionsSpec
			unknownBotProtectionSpec := botProtectionSpec.UnknownBotProtectionSpec
			recaptchaSpec := botProtectionSpec.ReCAPTCHASpec
			jsInjectionSpec := botProtectionSpec.JsInjectionSpec

			var paths []map[string]interface{}
			for _, path := range applicationSpec.APISpec.Paths {
				var methods []map[string]interface{}

				for _, method := range path.Methods {
					methods = append(methods, map[string]interface{}{
						"method": method.Method,
					})
				}

				paths = append(paths, map[string]interface{}{
					"path":    path.Path,
					"methods": methods,
				})
			}

			applicationsSpec = append(applicationsSpec, map[string]interface{}{
				"app_id": applicationSpec.AppID,
				"certificate": []map[string]interface{}{
					{
						"encrypted": applicationSpec.Certificate.Encrypted,
					},
				},
				"dos_config": []map[string]interface{}{
					{
						"enabled":       applicationSpec.DosConfig.Enabled,
						"track_session": applicationSpec.DosConfig.TrackSession,
						"alert": []map[string]interface{}{
							{
								"burst":   applicationSpec.DosConfig.Alert.Burst,
								"average": applicationSpec.DosConfig.Alert.Average,
							},
						},
						"ban": []map[string]interface{}{
							{
								"burst":   applicationSpec.DosConfig.Alert.Burst,
								"average": applicationSpec.DosConfig.Alert.Average,
							},
						},
						"excluded_network_lists": applicationSpec.DosConfig.ExcludedNetworkLists,
						"match_conditions":       dosConditions,
					},
				},
				"api_spec": []map[string]interface{}{
					{
						"endpoints":       endpoints,
						"paths":           paths,
						"effect":          applicationSpec.APISpec.Effect,
						"fallback_effect": applicationSpec.APISpec.FallbackEffect,
						"description":     applicationSpec.APISpec.Description,
						"skip_learning":   applicationSpec.APISpec.SkipLearning,
					},
				},
				"bot_protection_spec": []map[string]interface{}{
					{
						"user_defined_bots":  userDefinedBots,
						"session_validation": botProtectionSpec.SessionValidation,
						"interstitial_page":  botProtectionSpec.InterstitialPage,
						"known_bot_protections_spec": []map[string]interface{}{
							{
								"search_engine_crawlers": knownBotProtectionsSpec.SearchEngineCrawlers,
								"business_analytics":     knownBotProtectionsSpec.BusinessAnalytics,
								"educational":            knownBotProtectionsSpec.Educational,
								"news":                   knownBotProtectionsSpec.News,
								"financial":              knownBotProtectionsSpec.Financial,
								"content_feed_clients":   knownBotProtectionsSpec.ContentFeedClients,
								"archiving":              knownBotProtectionsSpec.Archiving,
								"career_search":          knownBotProtectionsSpec.CareerSearch,
								"media_search":           knownBotProtectionsSpec.MediaSearch,
							},
						},
						"unknown_bot_protections_spec": []map[string]interface{}{
							{
								"generic":               unknownBotProtectionSpec.Generic,
								"web_automation_tools":  unknownBotProtectionSpec.WebAutomationTools,
								"web_scrapers":          unknownBotProtectionSpec.WebScrapers,
								"api_libraries":         unknownBotProtectionSpec.APILibraries,
								"http_libraries":        unknownBotProtectionSpec.HTTPLibraries,
								"bot_impersonation":     unknownBotProtectionSpec.BotImpersonation,
								"browser_impersonation": unknownBotProtectionSpec.BrowserImpersonation,
								"request_anomalies": []map[string]interface{}{
									{
										"threshold": unknownBotProtectionSpec.RequestAnomalies.Threshold,
										"effect":    unknownBotProtectionSpec.RequestAnomalies.Effect,
									},
								},
							},
						},
						"js_injection_spec": []map[string]interface{}{
							{
								"enabled":        jsInjectionSpec.Enabled,
								"timeout_effect": jsInjectionSpec.TimeoutEffect,
							},
						},
						"recaptcha_spec": []map[string]interface{}{
							{
								"enabled":                  recaptchaSpec.Enabled,
								"site_key":                 recaptchaSpec.SiteKey,
								"type":                     recaptchaSpec.Type,
								"all_sessions":             recaptchaSpec.AllSessions,
								"success_expiration_hours": recaptchaSpec.SuccessExpirationHours,
								"secret_key": []map[string]interface{}{
									{
										"encrypted": recaptchaSpec.SecretKey.Encrypted,
									},
								},
							},
						},
					},
				},
				"network_controls": []map[string]interface{}{
					{
						"advanced_protection_effect": applicationSpec.NetworkControls.AdvancedProtectionEffect,
						"exception_subnets":          applicationSpec.NetworkControls.ExceptionSubnets,
						"subnets":                    *networkControlsEffectFromObject(applicationSpec.NetworkControls.Subnets),
						"countries":                  *networkControlsEffectFromObject(applicationSpec.NetworkControls.Countries),
					},
				},
				"body": []map[string]interface{}{
					{
						"skip":                  applicationSpec.Body.Skip,
						"inspection_size_bytes": applicationSpec.Body.InspectionSizeBytes,
					},
				},
				"header_specs": headerSpecs,
				"intel_gathering": []map[string]interface{}{
					{
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

				"sqli":           *applicationSpecEffectsFromObject(applicationSpec.Sqli),
				"xss":            *applicationSpecEffectsFromObject(applicationSpec.XSS),
				"attack_tools":   *applicationSpecEffectsFromObject(applicationSpec.AttackTools),
				"shellshock":     *applicationSpecEffectsFromObject(applicationSpec.Shellshock),
				"malformed_req":  *applicationSpecEffectsFromObject(applicationSpec.MalformedReq),
				"cmdi":           *applicationSpecEffectsFromObject(applicationSpec.Cmdi),
				"lfi":            *applicationSpecEffectsFromObject(applicationSpec.Lfi),
				"code_injection": *applicationSpecEffectsFromObject(applicationSpec.CodeInjection),

				"csrf_enabled":           applicationSpec.CsrfEnabled,
				"clickjacking_enabled":   applicationSpec.ClickjackingEnabled,
				"session_cookie_enabled": applicationSpec.SessionCookieEnabled,
				"session_cookie_ban":     applicationSpec.SessionCookieBan,
				"ban_duration_minutes":   applicationSpec.BanDurationMinutes,
				// "remote_host_forwarding":,

			})
		}

		collections := collectionSliceToInterface(i.Collections)

		waasRules = append(
			waasRules,
			map[string]interface{}{
				"name":                 i.Name,
				"read_timeout_seconds": i.ReadTimeoutSeconds,
				"collections":          collections,
				"applications_spec":    applicationsSpec,
			})
	}

	d.SetId("cnaf")

	err := d.Set("rules", waasRules)
	if err != nil {
		log.Printf("[ERROR] rules caused by: %s", err)
		return err
	}

	errMaxPort := d.Set("max_port", waasObject.MaxPort)
	if errMaxPort != nil {
		log.Printf("[ERROR] max_port caused by: %s", errMaxPort)
		return errMaxPort
	}

	errMinPort := d.Set("min_port", waasObject.MinPort)
	if errMinPort != nil {
		log.Printf("[ERROR] min_port caused by: %s", errMinPort)
		return errMinPort
	}

	return nil
}

func createWaasContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := waas.SetContainerWaas(*client, parseWaasContainer(d))
	if err != nil {
		return err
	}

	return readWaasContainer(d, meta)
}

func readWaasContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	waasRules, err := waas.GetContainerWaas(*client)
	if err != nil {
		return err
	}

	return saveWaasContainer(d, waasRules)
}

func deleteWaasContainer(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := waas.SetContainerWaas(*client, &waas.Waas{})
	if err != nil {
		return err
	}

	return nil
}

func applicationSpecEffectsFromInterface(applicationSpecEffects map[string]interface{}) *waas.ApplicationSpecEffects {

	applicationSpecObject := waas.ApplicationSpecEffects{}
	if applicationSpecEffects["effect"] != nil {
		applicationSpecObject.Effect = applicationSpecEffects["effect"].(string)
	}

	var exceptionFields []waas.ExceptionFields

	if applicationSpecEffects["exception_fields"] != nil {
		for _, exceptionField := range applicationSpecEffects["exception_fields"].([]interface{}) {
			field := exceptionField.(map[string]interface{})
			exceptionFields = append(exceptionFields,
				waas.ExceptionFields{
					Location: field["effect"].(string),
					Key:      field["key"].(string),
				})
		}
	}

	applicationSpecObject.ExceptionFields = exceptionFields

	return &applicationSpecObject
}

func applicationSpecEffectsFromObject(appSpec waas.ApplicationSpecEffects) *[]map[string]interface{} {

	var exceptionFields []map[string]interface{}
	for _, i := range appSpec.ExceptionFields {
		exceptionFields = append(exceptionFields,
			map[string]interface{}{
				"location": i.Location,
				"key":      i.Key,
			})
	}

	return &[]map[string]interface{}{
		{
			"effect":           appSpec.Effect,
			"exception_fields": exceptionFields,
		},
	}
}
func networkControlsEffectFromObject(ncEffects waas.NetworkControlsEffect) *[]map[string]interface{} {
	return &[]map[string]interface{}{
		{
			"enabled":         ncEffects.Enabled,
			"allow_mode":      ncEffects.AllowMode,
			"fallback_effect": ncEffects.FallbackEffect,
			"allow":           ncEffects.Allow,
			"alert":           ncEffects.Alert,
			"prevent":         ncEffects.Prevent,
		},
	}
}

func applicationSpecEffectsSchema() *schema.Schema {
	model := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"effect": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					waaf_effects,
					false),
			},
			"exception_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location": {
							Optional:    true,
							Type:        schema.TypeString,
							Description: "",
							ValidateFunc: validation.StringInSlice(
								[]string{
									"query",
									"body",
									"cookie",
									"header",
									"JSONPath",
									"XMLPath",
								},
								false,
							),
						},
						"key": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "",
						},
					},
				},
			},
		},
	}

	return &schema.Schema{
		Optional:    true,
		Computed:    true,
		Type:        schema.TypeSet,
		Description: "application spec effects",
		Elem:        model,
	}
}
