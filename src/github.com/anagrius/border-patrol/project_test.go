package main

import "testing"

func TestParsingModules(t *testing.T) {
	project := readProject("/Users/thb/Code/border-patrol/test/elm", "elm")
	expected := []string{"Api.Z", "Core.Y", "Util.X", "Main"}
	expectListsEq(t, sources(project), expected)
}

func TestParsingImports(t *testing.T) {
	project := readProject("/Users/thb/Code/border-patrol/test/elm", "elm")
	imports := importsBySource("Core.Y", project)
	expected := []string{"Html", "Html.Attributes", "Api.Z", "Util.X"}
	expectListsEq(t, imports, expected)
}

func TestParsingScala(t *testing.T) {
	rootDir := "/Users/thb/Code/border-patrol/test/scala"
	project := readProject(rootDir, "scala")
	imports := importsBySource("com.humio.core", project)
	expected := []string{"bar.xi.moo.X", "java.lang._", "com.humio.kafka._"}
	expectListsEq(t, imports, expected)

	config := readConfig(rootDir)
	report := Check(config, project)

	if len(report) != 2 {
		t.Error("Unexpected number of files")
	}
}
