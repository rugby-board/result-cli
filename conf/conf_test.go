package conf

import (
	"testing"
)

const defaultConfFile = "conf.yaml"

func TestGetConf(t *testing.T) {
	c, err := GetConf(defaultConfFile)
	if err != nil {
		t.Error("Get conf file failed")
	}
	if c.BaseURL != "http://kratos.365.co.za:9001/getresultsbycompidanddaterange/%d/%s/%s" {
		t.Error("Base url test failed")
	}
}
