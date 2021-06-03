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

func Index(c sdk.Client) ([]Tag, error) {
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

func Get(c sdk.Client, tagName string) (*Tag, error) {
	resp, err := Index(c)
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

func Create(c sdk.Client, spec *Tag) error {
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

func Update(c sdk.Client, tagName string, spec *Tag) error {
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

func Delete(c sdk.Client, tagName string) error {
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
