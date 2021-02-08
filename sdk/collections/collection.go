package collections

import (
	"fmt"
	"strings"

	"github.com/Hivebrite/twistlock-go/sdk"
)

type Collection struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Containers  []string `json:"containers"`
	Hosts       []string `json:"hosts"`
	Images      []string `json:"images"`
	Labels      []string `json:"labels"`
	AppIDs      []string `json:"appIDs"`
	Functions   []string `json:"functions"`
	Namespaces  []string `json:"namespaces"`
	AccountIDs  []string `json:"accountIDs"`
	CodeRepos   []string `json:"codeRepos"`
	Clusters    []string `json:"clusters"`
	Color       string   `json:"color"`
}

func Index(c sdk.Client) ([]Collection, error) {
	req, err := c.NewRequest("GET", "collections", nil)
	if err != nil {
		return nil, err
	}

	collections := []Collection{}
	_, err = c.Do(req, &collections)
	if err != nil {
		return nil, err
	}

	return collections, nil
}

func Get(c sdk.Client, collectionName string) (*Collection, error) {
	resp, err := Index(c)
	if err != nil {
		return nil, err
	}

	for _, i := range resp {
		if strings.Compare(collectionName, i.Name) == 0 {
			return &i, nil
		}
	}

	return nil, fmt.Errorf("collection: %s not found", collectionName)
}

func Update(c sdk.Client, collection *Collection) error {
	req, err := c.NewRequest("PUT", fmt.Sprintf("collections/%s", collection.Name), collection)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func Create(c sdk.Client, collection *Collection) error {
	req, err := c.NewRequest("POST", "collections", collection)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func Delete(c sdk.Client, collectionName string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("collections/%s", collectionName), nil)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
