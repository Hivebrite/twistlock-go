package sdk

import "time"

type Secret struct {
	Encrypted string `json:"encrypted"`
	Plain     string `json:"plain"`
}

type Collection struct {
	Hosts       []string  `json:"hosts"`
	Images      []string  `json:"images"`
	Labels      []string  `json:"labels"`
	Containers  []string  `json:"containers"`
	Functions   []string  `json:"functions"`
	Namespaces  []string  `json:"namespaces"`
	AppIDs      []string  `json:"appIDs"`
	AccountIDs  []string  `json:"accountIDs"`
	CodeRepos   []string  `json:"codeRepos"`
	Clusters    []string  `json:"clusters"`
	Name        string    `json:"name"`
	Owner       string    `json:"owner"`
	Modified    time.Time `json:"modified"`
	Color       string    `json:"color"`
	Description string    `json:"description"`
	System      bool      `json:"system"`
}
