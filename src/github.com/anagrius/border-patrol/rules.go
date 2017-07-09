package main

import (
	"encoding/json"
	"strings"
)

// Config ...
type Config struct {
	Restrictions map[string]([]string) `json:"restrictions"`
}

// LoadConfig ...
func LoadConfig(data []byte) (Config, error) {

	var config Config
	json.Unmarshal(data, &config)

	return config, nil
}

func Check(config Config, project map[string][]string) map[string][]string {
	errors := make(map[string][]string)
	for modulePrefix, disallowedImports := range config.Restrictions {
		candicates := matchingModules(sources(project), modulePrefix)
		for _, candicate := range candicates {
			for _, i := range disallowedImports {
				errorImport := checkImport(project[candicate], i)
				if errorImport != "" {
					if errors[candicate] == nil {
						errors[candicate] = make([]string, 0)
					}
					errors[candicate] = append(errors[candicate], "Restricted import "+errorImport)
				}
			}
		}
	}
	return errors
}

func checkImport(imports []string, anImport string) string {
	for _, x := range imports {
		if x == anImport || strings.HasPrefix(x, anImport+".") {
			return x
		}
	}
	return ""
}

func matchingModules(candicates []string, prefix string) []string {
	result := []string{}
	for _, v := range candicates {
		if v == prefix || strings.HasPrefix(v, prefix+".") {
			result = append(result, v)
		}
	}
	return result
}
