package policies

import (
	"fmt"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type CustomRuleRuntime struct {
	ID          int    `json:"_id"`
	Type        string `json:"type"`
	Message     string `json:"message"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

const (
	TypeProcess         = "processes"
	TypeFileSystem      = "filesystem"
	TypeNetworkOutgoing = "network-outgoing"
	TypeKubernetesAudit = "kubernetes-audit"
)

func GetCustomRuleRuntime(c sdk.Client, id int) (*CustomRuleRuntime, error) {
	req, err := c.NewRequest("GET", "policies/runtime/custom-rules", nil)
	if err != nil {
		return nil, err
	}
	var customRules []CustomRuleRuntime
	_, err = c.Do(req, &customRules)
	if err != nil {
		return nil, err
	}

	for _, i := range customRules {
		if i.ID == id {
			return &i, nil
		}
	}

	return &CustomRuleRuntime{}, sdk.ObjectNotFoundError

}

func SetCustomRuleRuntime(c sdk.Client, customRule *CustomRuleRuntime) error {
	req, err := c.NewRequest("PUT", fmt.Sprintf("policies/runtime/custom-rules/%d", customRule.ID), customRule)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCustomRuleRuntime(c sdk.Client, customRule *CustomRuleRuntime) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("policies/runtime/custom-rules/%d", customRule.ID), nil)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
