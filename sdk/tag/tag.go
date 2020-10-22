package tag

import (
	"fmt"
	"strings"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Tag struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Vulns []Vuln `json:"vulns,omitempty"`
}
type Vuln struct {
	ID          string `json:"id"`
	PackageName string `json:"packageName"`
	Comment     string `json:"comment"`
}

func GetTags(c sdk.Client) ([]Tag, error) {
	req, err := c.NewRequest("GET", "tags", nil)
	if err != nil {
		return nil, err
	}

	tags := []Tag{}
	_, err = c.Do(req, &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func GetTag(c sdk.Client, tagName string) (*Tag, error) {
	resp, err := GetTags(c)
	if err != nil {
		return nil, err
	}

	for _, i := range resp {
		if strings.Compare(tagName, i.Name) == 0 {
			return &i, nil
		}
	}

	return nil, fmt.Errorf("tag: %s not found", tagName)
}

func CreateTag(c sdk.Client, spec *Tag) error {
	req, err := c.NewRequest("POST", "tags", spec)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTag(c sdk.Client, tagName string, spec *Tag) error {
	req, err := c.NewRequest("PUT", fmt.Sprintf("tags/%s", tagName), spec)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTag(c sdk.Client, tagName string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("tags/%s", tagName), nil)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
