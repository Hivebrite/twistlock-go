package twistlock

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func fetchOptionalMapFromSetParam(rule map[string]interface{}, fieldName string) map[string]interface{} {
	schemaField := rule[fieldName].(*schema.Set).List()
	if len(schemaField) > 0 {
		return schemaField[0].(map[string]interface{})
	}
	return map[string]interface{}{}
}
