package parser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Auth struct {
		Urlvalidreg string   `yaml:"urlvalidreg"`
		Contenttype []string `yaml:"contenttype"`
	} `yaml:"auth"`
}

func (c *Config) String() string {
	return fmt.Sprintf("Url %s and content %s", c.Auth.Urlvalidreg, c.Auth.Contenttype[0])
}

// что по безопасности??? норм вернуть указатель?
func ParseConfig(path string) (*Config, error) {
	var config Config

	filename, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
