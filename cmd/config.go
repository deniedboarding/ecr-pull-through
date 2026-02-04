package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Registries     []string `yaml:"registries"`
	AwsAccountID   string   `yaml:"awsAccountId"`
	AwsRegion      string   `yaml:"awsRegion,omitempty"`
	ExcludedImages []string `yaml:"excludedImages"`

	excludedImageMap    map[string]bool
	ecrRegistryEndpoint string
}

func ReadConf(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := Config{}
	err = yaml.Unmarshal(buf, &content)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	content.excludedImageMap = make(map[string]bool)
	for _, image := range content.ExcludedImages {
		content.excludedImageMap[image] = true
	}

	content.ecrRegistryEndpoint = fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", content.AwsAccountID, content.AwsRegion)

	return &content, err
}

func (c *Config) RegistryList() []string {
	if len(c.Registries) == 0 {
		return []string{"docker.io"}
	}
	return c.Registries
}
