package full_outer_join

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	got1, err := os.ReadFile(f1Path)
	if err != nil {
		fmt.Printf("could not read first file: %s", err)
		os.Exit(1)
	}
	got2, err := os.ReadFile(f2Path)
	if err != nil {
		fmt.Printf("could not read second file: %s", err)
		os.Exit(1)
	}
	words1 := strings.Split(string(got1), "\n")
	words2 := strings.Split(string(got2), "\n")
	mappedWords1 := make(map[string]bool)
	mappedWords2 := make(map[string]bool)
	for _, w := range words1 {
		mappedWords1[w] = true
	}
	for _, w := range words2 {
		mappedWords2[w] = true
	}
	result := make([]byte, 0)
	for _, w := range words2 {
		if _, ok := mappedWords1[w]; !ok {
			result = append(result, w...)
			result = append(result, "\n"...)
		}
	}
	for _, w := range words1 {
		if _, ok := mappedWords2[w]; !ok {
			result = append(result, w...)
			result = append(result, "\n"...)
		}
	}
	result = result[:len(result)-1]
	sorted := strings.Split(string(result), "\n")
	sort.Strings(sorted)
	result = nil
	for _, w := range sorted {
		result = append(result, w...)
		result = append(result, "\n"...)
	}
	result = result[:len(result)-1]
	if err := os.WriteFile(resultPath, []byte(result), os.ModePerm); err != nil {
		fmt.Printf("could not write result: %s", err)
	}
}
