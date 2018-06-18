package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Body ...
type Body struct {
	BaseURL string `yaml:"base_url"`
}

const (
	// RepoPath Relative path of this repo
	RepoPath = "github.com/rugby-board/result-cli"
)

// LoadEnvConfPath ...
func LoadEnvConfPath(confPath string) (string, error) {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		return "", errors.New("GOPATH is not set")
	}

	path := fmt.Sprintf("%s/src/%s/%s", gopath, RepoPath, confPath)
	return path, nil
}

// GetConf for result-cli
func GetConf(confPath string) (*Body, error) {
	c := &Body{}
	realConfPath, err := LoadEnvConfPath(confPath)
	if err != nil {
		log.Fatalf("Get real conf path failed: %#v", err)
	}
	yamlFile, err := ioutil.ReadFile(realConfPath)
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
