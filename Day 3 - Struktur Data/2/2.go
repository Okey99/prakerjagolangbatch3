package main

import (
	"fmt"
)

func countOccurrences(slice []string, str string) int {
	count := 0
	for _, item := range slice {
		if item == str {
			count++
		}
	}
	return count
}

func main() {
	stringSlice := []string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}
	targetString1 := "asd"
	targetString2 := "qwe"
	targetString3 := "adi"

	occurrences := countOccurrences(stringSlice, targetString1, targetString2, targetString3)
	fmt.Printf("Jumlah kemunculan dari string '%s' adalah %d\n", targetString1, targetString2, targetString3, occurrences)
}
