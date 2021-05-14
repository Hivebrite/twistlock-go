package twistlock

import (
	"time"

	"github.com/Hivebrite/twistlock-go/sdk"
	"github.com/Hivebrite/twistlock-go/sdk/policies"
)

const (
	Ban     = "ban"
	Allow   = "allow"
	Block   = "block"
	Alert   = "alert"
	Prevent = "prevent"
	Disable = "disable"
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

func tagObjectFromInterface(tag map[string]interface{}) *policies.Tags {
	expiration := fetchOptionalMapFromSetParam(tag, "expiration")
	tagObject := policies.Tags{
		Effect:      tag["effect"].(string),
		Description: tag["description"].(string),
		Name:        tag["name"].(string),
	}

	if len(expiration) != 0 {
		if expiration["enabled"] == nil {
			tagObject.Expiration.Enabled = false
		} else {
			tagObject.Expiration.Enabled = expiration["disabled"].(bool)
		}

		if expiration["date"] != nil && expiration["enabled"] == true {
			tagObject.Expiration.Date = *stringToTime(expiration["date"].(string))
		}
	}

	return &tagObject
}
func cveRuleObjectFromInterface(cveRule map[string]interface{}) *policies.CveRules {
	expiration := fetchOptionalMapFromSetParam(cveRule, "expiration")
	cveRuleObject := policies.CveRules{
		Effect:      cveRule["effect"].(string),
		Description: cveRule["description"].(string),
		ID:          cveRule["id"].(string),
	}

	if len(expiration) != 0 {
		if expiration["enabled"] == nil {
			cveRuleObject.Expiration.Enabled = false
		} else {
			cveRuleObject.Expiration.Enabled = expiration["disabled"].(bool)
		}

		if expiration["date"] != nil && expiration["enabled"] == true {
			cveRuleObject.Expiration.Date = *stringToTime(expiration["date"].(string))
		}
	}

	return &cveRuleObject
}
func stringToTime(stringifiedTime string) *time.Time {
	layout := "2006-01-02T15:04:05.999Z00:00"
	expirationDate, _ := time.Parse(layout, stringifiedTime)

	return &expirationDate
}

func timeToString(timeObject time.Time) string {
	layout := "2006-01-02T15:04:05.999Z00:00"
	return timeObject.Format(layout)
}

func expirationMapFromObject(expiration policies.Expiration) *map[string]interface{} {
	schema := map[string]interface{}{}

	if expiration.Enabled {
		schema["enabled"] = true
		schema["date"] = timeToString(expiration.Date)
	}
	return &schema
}

func parseCollections(collectionsList []interface{}) []sdk.Collection {
	var collectionsSlice []sdk.Collection

	for _, collection := range collectionsList {
		c := collection.(map[string]interface{})
		collectionsSlice = append(collectionsSlice,
			sdk.Collection{
				Name: c["name"].(string),
			})
	}

	return collectionsSlice
}

func collectionSliceToInterface(collectionsSlice []sdk.Collection) []map[string]interface{} {
	var collections []map[string]interface{}

	for _, collection := range collectionsSlice {
		collections = append(collections,
			map[string]interface{}{
				"name": collection.Name,
			},
		)
	}

	return collections
}
