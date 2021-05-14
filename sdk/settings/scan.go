package settings

import (
	"github.com/Hivebrite/twistlock-go/sdk"
)

type ScanSettings struct {
	ScanRunningImages             bool `json:"scanRunningImages"`
	ImagesScanPeriodMs            int  `json:"imagesScanPeriodMs"`
	ContainersScanPeriodMs        int  `json:"containersScanPeriodMs"`
	SystemScanPeriodMs            int  `json:"systemScanPeriodMs"`
	RegistryScanPeriodMs          int  `json:"registryScanPeriodMs"`
	TasDropletsScanPeriodMs       int  `json:"tasDropletsScanPeriodMs"`
	ServerlessScanPeriodMs        int  `json:"serverlessScanPeriodMs"`
	CloudPlatformsScanPeriodMs    int  `json:"cloudPlatformsScanPeriodMs"`
	VMScanPeriodMs                int  `json:"vmScanPeriodMs"`
	CodeReposScanPeriodMs         int  `json:"codeReposScanPeriodMs"`
	RegistryScanRetentionDays     int  `json:"registryScanRetentionDays"`
	ExtractArchive                bool `json:"extractArchive"`
	ShowNegligibleVulnerabilities bool `json:"showNegligibleVulnerabilities"`
	ShowInfraContainers           bool `json:"showInfraContainers"`
	DisableCLIScanResults         bool `json:"disableCLIScanResults"`
	IncludeJsDependencies         bool `json:"includeJsDependencies"`
}

func GetScanSettings(c sdk.Client) (*ScanSettings, error) {
	req, err := c.NewRequest("GET", "settings/scan", nil)

	if err != nil {
		return nil, err
	}

	var settings ScanSettings
	_, err = c.Do(req, &settings)

	if err != nil {
		return nil, err
	}

	return &settings, nil
}

func UpdateScanSettings(c sdk.Client, settings *ScanSettings) error {
	req, err := c.NewRequest("POST", "settings/scan", settings)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
