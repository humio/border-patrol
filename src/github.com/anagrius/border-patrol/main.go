package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
)

var importRegexp = regexp.MustCompile(`^import (\S+)`)
var moduleRegexp = regexp.MustCompile(`^(?:port|effect)?\s*module\s+(\S+)`)

func main() {
	rootDir := os.Args[1]

	config := readConfig(rootDir)
	project := readProject(rootDir)
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
