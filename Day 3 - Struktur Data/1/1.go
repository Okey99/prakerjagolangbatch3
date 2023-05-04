package main

func mergeArrays(arr1 []string, arr2 []string) []string {
	mergedArr := append(arr1, arr2...)
	uniqueNames := make(map[string]bool)
	result := []string{}

	for _, name := range mergedArr {
		if _, value := uniqueNames[name]; !value {
			uniqueNames[name] = true
			result = append(result, name)
		}
	}

	return result
}

// func main() {
// 	array1 := []string{"king", "devil jin", "akuma"}
// 	array2 := []string{"eddie", "steve", "geese"}
// 	mergedArray := mergeArrays(array1, array2)
// 	fmt.Println(mergedArray)
// }

// func main() {
// 	array1 := []string{"sergei", "jin"}
// 	array2 := []string{"jin", "steve", "bryan"}
// 	mergedArray := mergeArrays(array1, array2)
// 	fmt.Println(mergedArray)
// }

// func main() {
// 	array1 := []string{"alisa", "yoshimitsu"}
// 	array2 := []string{"devil jin", "yoshimitsu", "alisa", "law"}

// 	mergedArray := mergeArrays(array1, array2)
// 	fmt.Println(mergedArray)
// }

// func main() {
// 	array1 := []string{}
// 	array2 := []string{"devil jin", "sergei"}

// 	mergedArray := mergeArrays(array1, array2)
// 	fmt.Println(mergedArray)
// }

// func main() {
// 	array1 := []string{"hwoarang"}
// 	array2 := []string{}

// 	mergedArray := mergeArrays(array1, array2)
// 	fmt.Println(mergedArray)
// }

// func main() {
// 	array1 := []string{}
// 	array2 := []string{}

// 	mergedArray := mergeArrays(array1, array2)
// 	fmt.Println(mergedArray)
// }
