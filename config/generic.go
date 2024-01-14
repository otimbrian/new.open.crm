package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ohler55/ojg/oj"
)

type GenericConfig struct {
	configData map[string]interface{}
}

func (config *GenericConfig) get(name string) (result interface{}, found bool) {
	data := config.configData

	for _, key := range strings.Split(name, ":") {
		result, found = data[key]

		if newSection, ok := result.(map[string]interface{}); ok && found {
			data = newSection
		} else {
			return
		}
	}
	return
}

func (config *GenericConfig) GetString(name string) (configValue string, found bool) {
	value, found := config.get(name)

	if found {
		configValue = value.(string)
		return
	}
	return
}

func (config *GenericConfig) GetSection(name string) (section Configuration, found bool) {
	value, found := config.get(name)

	if found {
		if sectionData, ok := value.(map[string]interface{}); ok {
			section = &GenericConfig{configData: sectionData}
			return
		}
	}
	return
}

// func deserializeConfig[model interface{}](fileName string, selector string) (modelInstance model) {

// 	configFile, err := Load(fileName)

// 	if err == nil {
// 		requiredSection, found := configFile.GetSection(selector)

// 		if found {
// 			jsonString := oj.JSON(requiredSection)
// 			fmt.Println("Json String ====>", jsonString)
// 			objectByte, err := oj.ParseString(jsonString)

// 			if err == nil {
// 				fmt.Println("Object Byte ====>", objectByte)
// 			}
// 		}
// 	}

// }

func DeserilizeConfig[model interface{}](configFile Configuration, selector string) (modelInstance model) {
	// configFile, err := Load(name)

	requiredSection, found := configFile.GetSection(selector)

	stringValue := oj.JSON(requiredSection)
	fmt.Println("requiredSection ====>", requiredSection)
	fmt.Println("String Value =====>", stringValue)
	// var modelInstance model

	if found {
		byteData, err := json.Marshal(requiredSection)
		fmt.Println("Byte Data ====>", byteData)

		if err == nil {
			err := json.Unmarshal(byteData, &modelInstance)

			if err == nil {
				return modelInstance
			}

		}

	}
	return
}
