package main

// Pair ...

func sources(group map[string][]string) []string {
	sources := []string(nil)
	for k := range group {
		sources = append(sources, k)
	}
	return removeDuplicatesUnordered(sources)
}

func importsBySource(source string, group map[string][]string) []string {
	imports := []string(nil)
	for k, v := range group {
		if k == source {
			imports = append(imports, v...)
		}
	}
	return removeDuplicatesUnordered(imports)
}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
