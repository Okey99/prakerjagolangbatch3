package main

import (
	"fmt"
	"strings"
)

func munculSekali(input string) []int {
	uniqueMap := make(map[int]int)
	numbers := strings.Fields(input)
	for _, numberStr := range numbers {
		var number int
		fmt.Sscan(numberStr, &number)
		uniqueMap[number]++
	}
	var result []int
	for number, count := range uniqueMap {
		if count == 1 {
			result = append(result, number)
		}
	}
	return result
}

// func main() {
// 	input := "1 2 3 4 1 2 3"
// 	uniqueNumbers := munculSekali(input)
// 	fmt.Println("Angka-angka yang hanya muncul 1 kali:", uniqueNumbers)
// }

// func main() {
// 	input := "7 6 5 2 3 7 5 2"
// 	uniqueNumbers := munculSekali(input)
// 	fmt.Println("Angka-angka yang hanya muncul 1 kali:", uniqueNumbers)
// }

// func main() {
// 	input := "1 2 3 4 5"
// 	uniqueNumbers := munculSekali(input)
// 	fmt.Println("Angka-angka yang hanya muncul 1 kali:", uniqueNumbers)
// }

// func main() {
// 	input := "1 1 2 2 3 3 4 4 5 5"
// 	uniqueNumbers := munculSekali(input)
// 	fmt.Println("Angka-angka yang hanya muncul 1 kali:", uniqueNumbers)
// }

func main() {
	input := "0 8 7 2 5 0 4"
	uniqueNumbers := munculSekali(input)
	fmt.Println("Angka-angka yang hanya muncul 1 kali:", uniqueNumbers)
}
