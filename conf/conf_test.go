package conf

import (
	"testing"
)

const defaultConfFile = "conf/conf.yaml"

func TestGetConf(t *testing.T) {
	c, err := GetConf(defaultConfFile)
	if err != nil {
		t.Error("Get conf file failed")
	}
	if c.BaseURL != "http://kratos.365.co.za:9001/getresultsbycompidanddaterange/%d/%s/%s" {
		t.Error("Base url test failed")
	}
}

func TestLoadYamlFail(t *testing.T) {
	_, err := GetConf("conf/not_existed.yaml")
	if err == nil {
		t.Error("Should load yaml failed")
	}
}
