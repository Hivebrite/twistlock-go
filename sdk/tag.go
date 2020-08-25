package sdk

import (
	"fmt"
	"strings"
)

type Tag struct {
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Vulns []Vulns `json:"vulns,omitempty"`
}
type Vulns struct {
	ID          string `json:"id"`
	PackageName string `json:"packageName"`
	Comment     string `json:"comment"`
}

func (c *Client) GetTags() ([]Tag, error) {
	req, err := c.newRequest("GET", "tags", nil)
	if err != nil {
		return nil, err
	}

	tags := []Tag{}
	_, err = c.do(req, &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (c *Client) GetTag(tagName string) (*Tag, error) {
	resp, err := c.GetTags()
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

func (c *Client) CreateTag(spec *Tag) error {
	req, err := c.newRequest("POST", "tags", spec)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateTag(tagName string, spec *Tag) error {
	req, err := c.newRequest("PUT", fmt.Sprintf("tags/%s", tagName), spec)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteTag(tagName string) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("tags/%s", tagName), nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
