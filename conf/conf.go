package conf

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Body ...
type Body struct {
	BaseURL string `yaml:"base_url"`
}

// GetConf for result-cli
func GetConf(confPath string) (*Body, error) {
	c := &Body{}
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Read file failed: %#v", err)
		return c, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal failed: %#v", err)
		return c, err
	}
	return c, err
}
