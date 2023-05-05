package main

import "fmt"

func findMinMax(numbers *[6]int) (int, int) {
	var max, min int

	// Inisialisasi max dan min dengan nilai dari array pertama
	max = numbers[0]
	min = numbers[0]

	// Looping array untuk mencari nilai max dan min
	for _, number := range numbers {
		if number > max {
			max = number
		}

		if number < min {
			min = number
		}
	}

	// Mengembalikan nilai max dan min
	return max, min
}

func main() {
	// Deklarasi array 6 angka
	var numbers [6]int

	// Mengisi array dengan input dari user
	fmt.Println("Masukkan 6 angka:")
	for i := 0; i < len(numbers); i++ {
		fmt.Scan(&numbers[i])
	}

	// Memanggil fungsi findMinMax dengan pointer referencing
	max, min := findMinMax(&numbers)

	// Menampilkan nilai max dan min
	fmt.Printf("Nilai maksimum: %d\n", max)
	fmt.Printf("Nilai minimum: %d\n", min)
}
