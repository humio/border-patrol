package main

import "testing"

func TestParsingModules(t *testing.T) {
	project := readProject("/Users/thb/Code/border-patrol/test/elm")
	expected := []string{"Api.Z", "Core.Y", "Util.X", "Main"}
	expectListsEq(t, sources(project), expected)
}

func TestParsingImports(t *testing.T) {
	project := readProject("/Users/thb/Code/border-patrol/test/elm")
	imports := importsBySource("Core.Y", project)
	expected := []string{"Html", "Html.Attributes", "Api.Z", "Util.X"}
	expectListsEq(t, imports, expected)
}
