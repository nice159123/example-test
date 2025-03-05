package main

import "fmt"

func processArrays(arr1, arr2 []string) ([]string, []string) {
	countMap := make(map[string]int)
	var uniqueResult []string
	var diffResult []string

	for i := 0; i < len(arr1); i++ {
		countMap[arr1[i]]++
	}
	
	for i := 0; i < len(arr2); i++ {
		countMap[arr2[i]]++
	}

	keys := make([]string, 0, len(countMap))
	for key := range countMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if countMap[key] == 1 {
			diffResult = append(diffResult, key)
		}
		uniqueResult = append(uniqueResult, key)
	}

	return uniqueResult, diffResult
}

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}

	union, diff := processArrays(arr1, arr2)

	fmt.Println("Input Arr 1", arr1)
	fmt.Println("Input Arr 2", arr2)
	fmt.Println("Output Union:", union)
	fmt.Println("Output Difference:", diff)
}
