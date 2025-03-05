package main

import (
	"fmt"
	"sort"
)

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

	keys := make([]string, len(countMap))
	index := 0
	mapKeys := make([]string, 0, len(countMap))
	for k := range countMap {
		mapKeys = append(mapKeys, k)
	}

	for i := 0; i < len(mapKeys); i++ {
		keys[index] = mapKeys[i]
		index++
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

	fmt.Println("Input Arr 1:", arr1)
	fmt.Println("Input Arr 2:", arr2)
	fmt.Println("Output Union:", union)
	fmt.Println("Output Difference:", diff)
}
