package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/settings"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceScanSettings() *schema.Resource {
	return &schema.Resource{
		Create: createScanSettings,
		Read:   readScanSettings,
		Update: createScanSettings,
		Delete: deleteScanSettings,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scan_running_images": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     true,
			},
			"extract_archive": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"show_negligible_vulnerabilities": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"show_infra_containers": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     true,
			},
			"disable_cli_scan_results": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"include_js_dependencies": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "",
				Default:     false,
			},
			"images_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"containers_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"system_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"registry_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"tas_droplets_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"serverless_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"cloud_platforms_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"vm_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"code_repos_scan_period_ms": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     86400000,
			},
			"registry_scan_retention_days": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "",
				Default:     1,
			},
		},
	}
}

func parseScanSettings(d *schema.ResourceData) *settings.ScanSettings {
	return &settings.ScanSettings{
		ScanRunningImages:             d.Get("scan_running_images").(bool),
		ImagesScanPeriodMs:            d.Get("images_scan_period_ms").(int),
		ContainersScanPeriodMs:        d.Get("containers_scan_period_ms").(int),
		SystemScanPeriodMs:            d.Get("system_scan_period_ms").(int),
		RegistryScanPeriodMs:          d.Get("registry_scan_period_ms").(int),
		TasDropletsScanPeriodMs:       d.Get("tas_droplets_scan_period_ms").(int),
		ServerlessScanPeriodMs:        d.Get("serverless_scan_period_ms").(int),
		CloudPlatformsScanPeriodMs:    d.Get("cloud_platforms_scan_period_ms").(int),
		VMScanPeriodMs:                d.Get("vm_scan_period_ms").(int),
		CodeReposScanPeriodMs:         d.Get("code_repos_scan_period_ms").(int),
		RegistryScanRetentionDays:     d.Get("registry_scan_retention_days").(int),
		ExtractArchive:                d.Get("extract_archive").(bool),
		ShowNegligibleVulnerabilities: d.Get("show_negligible_vulnerabilities").(bool),
		ShowInfraContainers:           d.Get("show_infra_containers").(bool),
		DisableCLIScanResults:         d.Get("disable_cli_scan_results").(bool),
		IncludeJsDependencies:         d.Get("include_js_dependencies").(bool),
	}
}

func saveScanSettings(d *schema.ResourceData, settings *settings.ScanSettings) error {
	d.SetId("ScanSettings")
	d.Set("scan_running_images", settings.ScanRunningImages)
	d.Set("images_scan_period_ms", settings.ImagesScanPeriodMs)
	d.Set("containers_scan_period_ms", settings.ContainersScanPeriodMs)
	d.Set("system_scan_period_ms", settings.SystemScanPeriodMs)
	d.Set("registry_scan_period_ms", settings.RegistryScanPeriodMs)
	d.Set("tas_droplets_scan_period_ms", settings.TasDropletsScanPeriodMs)
	d.Set("serverless_scan_period_ms", settings.ServerlessScanPeriodMs)
	d.Set("cloud_platforms_scan_period_ms", settings.CloudPlatformsScanPeriodMs)
	d.Set("vm_scan_period_ms", settings.VMScanPeriodMs)
	d.Set("code_repos_scan_period_ms", settings.CodeReposScanPeriodMs)
	d.Set("registry_scan_retention_days", settings.RegistryScanRetentionDays)
	d.Set("extract_archive", settings.ExtractArchive)
	d.Set("show_negligible_vulnerabilities", settings.ShowNegligibleVulnerabilities)
	d.Set("show_infra_containers", settings.ShowInfraContainers)
	d.Set("disable_cli_scan_results", settings.DisableCLIScanResults)
	d.Set("include_js_dependencies", settings.IncludeJsDependencies)

	return nil
}

func createScanSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	err := settings.UpdateScanSettings(*client, parseScanSettings(d))
	if err != nil {
		return err
	}

	return readScanSettings(d, meta)
}

func readScanSettings(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	settings, err := settings.GetScanSettings(*client)
	if err != nil {
		return err
	}

	return saveScanSettings(d, settings)
}

func deleteScanSettings(d *schema.ResourceData, meta interface{}) error {
	return nil
}
