package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var importRegexp = regexp.MustCompile(`^import (\S+)`)
var moduleRegexp = regexp.MustCompile(`^(?:port|effect)?\s*module\s+(\S+)`)

func readProject(rootDir string, extension string) map[string][]string {
	entries, err := ioutil.ReadDir(rootDir)

	if err != nil {
		log.Fatal(err)
	}

	project := make(map[string][]string)

	for _, e := range entries {

		ePath := filepath.Join(rootDir, e.Name())
		if e.IsDir() {
			for k, v := range readProject(ePath, extension) {
				project[k] = v
			}
		} else {
			if !strings.HasSuffix(e.Name(), "."+extension) {
				continue
			}

			if extension == "scala" {
				project = mergeInto(project, parseScalaFile(ePath))
			} else {
				project = mergeInto(project, parseFile(ePath))
			}
		}
	}

	return project
}

func mergeInto(existing map[string][]string, newStuff map[string][]string) map[string][]string {
	for k, v := range newStuff {
		existing[k] = append(existing[k], v...)
	}

	return existing
}

func parseFile(filepath string) map[string][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	imports := []string{}
	module := ""

	for scanner.Scan() {
		line := scanner.Text()

		// If we are reading the module line set the FROM field to that.

		if module == "" {
			moduleMatch := moduleRegexp.FindStringSubmatch(line)

			if moduleMatch != nil {
				module = moduleMatch[1]
			}

			continue
		}
		//

		match := importRegexp.FindStringSubmatch(line)

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
