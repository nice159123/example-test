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

	mapKeys := make([]string, len(countMap))
	keysIndex := 0
	keysArray := make([]string, 0, len(countMap)) // Temporary array

	keyList := make([]string, 0, len(countMap))
	for key := range countMap {
		keyList = append(keyList, key)
	}

	for i := 0; i < len(keyList); i++ {
		keysArray = append(keysArray, keyList[i])
	}

	for i := 0; i < len(keysArray); i++ {
		mapKeys[keysIndex] = keysArray[i]
		keysIndex++
	}

	sort.Strings(mapKeys)

	for i := 0; i < len(mapKeys); i++ {
		key := mapKeys[i]
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
