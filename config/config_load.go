package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/Jeffail/gabs"
)

func Load(fileName string) (config Configuration, err error) {
	var data []byte

	data, err = os.ReadFile(fileName)

	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		m := map[string]interface{}{}
		err := decoder.Decode(&m)

		if err == nil {
			config = &GenericConfig{configData: m}
		}
	}
	return
}

func GetConfig(fileName string) (consumedData *gabs.Container) {
	data, err := os.ReadFile(fileName)

	if err == nil {
		consumedData, err = gabs.ParseJSON(data)
		if err == nil {
			return consumedData
		}
		return
	}
	return

}
