package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

var scalaImportRegexp = regexp.MustCompile(`^\s*import (.+)`)
var scalaPackageRegexp = regexp.MustCompile(`^\s*package\s+(\S+)`)

func parseScalaFile(filepath string) map[string][]string {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return parseScala(file)
}

func parseScala(content []byte) map[string][]string {
	scanner := bufio.NewScanner(strings.NewReader(string(content)))

	imports := []string{}
	module := ""

	for scanner.Scan() {
		line := scanner.Text()

		// If we are reading the module line set the FROM field to that.

		if module == "" {
			moduleMatch := scalaPackageRegexp.FindStringSubmatch(line)

			if moduleMatch != nil {
				module = moduleMatch[1]
			}

			continue
		}

		match := scalaImportRegexp.FindStringSubmatch(line)

		if match != nil && len(match) == 2 {
			imports = append(imports, match[1])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if module == "" {
		// No module found in the file. Skip it.
		return make(map[string][]string, 0)
	}

	return map[string][]string{module: imports}
}
