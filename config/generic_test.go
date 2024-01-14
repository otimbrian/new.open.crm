package config_test

import (
	"fmt"
	"general-utils/config"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	configu, err := config.Load("../config.json")

	if err != nil {
		t.FailNow()
	}

	fmt.Println(configu)
}

func TestDesirializeConfig(t *testing.T) {

	type DBConfig struct {
		DbName   string `json:"dbname"`
		Enforce  bool   `json:"enforce"`
		Host     string `json:"host"`
		Password string `json:"password"`
		Port     string `json:"port"`
		Sslmode  string `json:"sslmode"`
		TimeZone string `json:"timeZone"`
		User     string `json:"user"`
	}
	file, _ := config.Load("../config.json")

	requiredSection, found := file.GetSection("data")

	if !found {
		t.FailNow()
	}

	data := config.DeserilizeConfig[DBConfig](requiredSection, "connectionString")

	fmt.Println("Data ====>", data)
}

func TestGetString(t *testing.T) {
	configuration, err := config.Load("../tetsConfig.json")

	if err != nil {
		t.FailNow()
	}

	mainSection, found := configuration.GetSection("main")
	if !found {
		t.FailNow()
	}

	msg, found := mainSection.GetString("message")

	if !found || !strings.EqualFold(msg, "Hello This is the main section.") {
		t.FailNow()
	}

	fmt.Println(msg)
}

func TestGetSection(t *testing.T) {
	configuration, err := config.Load("../tetsConfig.json")

	if err != nil {
		t.FailNow()
	}
	mainSection, found := configuration.GetSection("main")
	if !found {
		t.FailNow()
	}

	fmt.Println(mainSection)
}

func TestGetConfig(t *testing.T) {
	data := config.GetConfig("../config.json")
	fmt.Println(data.Data())
}
