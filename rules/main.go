package rules

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	rootDir := os.Args[1]

	config := readConfig(rootDir)
	extension := extensionForLanguage(config.Language)
	project := readProject(rootDir, extension)
	report := Check(config, project)

	printReport(report)

	if len(report) > 0 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func readConfig(rootDir string) Config {
	configFilePath := path.Join(rootDir, "boundaries.json")
	file, e := ioutil.ReadFile(configFilePath)

	if e != nil {
		log.Fatal(e)
	}

	config, err := LoadConfig(file)

	if err != nil {
		log.Fatal(err)
	}

	return config
}
