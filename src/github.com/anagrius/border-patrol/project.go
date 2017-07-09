package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func readProject(rootDir string) map[string][]string {
	entries, err := ioutil.ReadDir(rootDir)

	if err != nil {
		log.Fatal(err)
	}

	project := make(map[string][]string)

	for _, e := range entries {

		ePath := filepath.Join(rootDir, e.Name())
		if e.IsDir() {
			for k, v := range readProject(ePath) {
				project[k] = v
			}
		} else {
			// TODO: Consider supporting 'elmx' extension.
			if !strings.HasSuffix(e.Name(), ".elm") {
				continue
			}

			for k, v := range parseFile(ePath) {
				project[k] = v
			}
		}
	}

	return project
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
