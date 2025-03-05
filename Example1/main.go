package main

import "fmt"

func processArrays(arr1, arr2 []string) ([]string, []string) {
	countMap := make(map[string]int)
	var uniqueResult []string
	var diffResult []string

	for _, val := range arr1 {
		countMap[val]++
	}
	for _, val := range arr2 {
		countMap[val]++
	}

	for key, count := range countMap {
		if count == 1 {
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
