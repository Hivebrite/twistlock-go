package twistlock

import (
	"github.com/Hivebrite/twistlock-go/sdk/policies"
)

func blockThresholdFromRule(rule map[string]interface{}) *policies.BlockThreshold {
	blockThreshold := rule["block_threshold"].([]interface{})[0].(map[string]interface{})

	return &policies.BlockThreshold{
		Enabled: blockThreshold["enabled"].(bool),
		Value:   policies.AlertingLevelToInt(blockThreshold["value"].(string)),
	}
}

func alertThresholdFromRule(rule map[string]interface{}) *policies.AlertThreshold {
	alertThreshold := rule["alert_threshold"].([]interface{})[0].(map[string]interface{})

	return &policies.AlertThreshold{
		Disabled: alertThreshold["disabled"].(bool),
		Value:    policies.AlertingLevelToInt(alertThreshold["value"].(string)),
	}
}
