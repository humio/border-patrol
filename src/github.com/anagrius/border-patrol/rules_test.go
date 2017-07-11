package main

import (
	"testing"
)

func TestReadingRestrictions(t *testing.T) {
	config, err := LoadConfig([]byte(`
    {
			"language": "elm",
      "restrictions": {
        "Api.Z": ["Core.Y", "Main"],
        "Core.Y": ["Main"]
      }
    }`))

	if err != nil {
		t.Error("Failed to Read Json", err)
	}

	expected := []string{"Core.Y", "Main"}
	expectListsEq(t, config.Restrictions["Api.Z"], expected)
}

func TestCheckingProject(t *testing.T) {
	config, err := LoadConfig([]byte(`
    {
      "restrictions": {
        "Api.Z": ["Core.Y", "Main"],
        "Core": ["Foo", "Main", "Api.Z"]
      }
    }`))

	project := map[string][]string{
		"Core.Y": []string{"Foo.Bar", "Api.Z", "Api.C"},
		"Api.U":  []string{"Basics", "Api.Z"},
	}

	if err != nil {
		t.Error("Failed to Read Json", err)
	}

	expected := []string{"import Foo.Bar", "import Api.Z"}
	result := Check(config, project)

	expectListsEq(t, result["Core.Y"], expected)
	if len(result) != 1 {
		t.Error("expected only one entry")
	}
}
