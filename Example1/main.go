package main

import "fmt"

func uniqueArray(arr1, arr2 []string) []string {
	uniqueMap := make(map[string]bool)
	var result []string
	for _, val := range arr1 {
		uniqueMap[val] = true
	}
	for _, val := range arr2 {
		uniqueMap[val] = true
	}
	for key := range uniqueMap {
		result = append(result, key)
	}
	return result
}

func uniqueArrayDifference(arr1, arr2 []string) []string {
	countMap := make(map[string]int)
	var result []string
	for _, val := range arr1 {
		countMap[val]++
	}
	for _, val := range arr2 {
		countMap[val]++
	}
	for key, count := range countMap {
		if count == 1 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	arr1 := []string{"a", "b", "c"}
	arr2 := []string{"b", "c", "d"}

	unionResult := uniqueArray(arr1, arr2)
	diffResult := uniqueArrayDifference(arr1, arr2)

	fmt.Println("Input Arr 1", arr1)
	fmt.Println("Input Arr 2", arr2)
	fmt.Println("Output Union:", unionResult)
	fmt.Println("Output Difference:", diffResult)
}
