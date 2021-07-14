package main

import "fmt"

func main() {
	input := []int{9, 7, 6, 1, 3, 5, 4, 2, 0}
	result := mergeSort(input)
	fmt.Println(result)
}

func merge(arr1 []int, arr2 []int) []int {

	i, j := 0, 0

	fmt.Println("Left", arr1)
	fmt.Println("Right", arr2)
	result := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i = i + 1
		} else {
			result = append(result, arr2[j])
			j = j + 1
		}
	}
	if i == len(arr1) {
		result = append(result, arr2[j:]...)
	}
	if j == len(arr2) {
		result = append(result, arr1[i:]...)
	}

	return result

}

var final []int = make([]int, 0)

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	l := len(arr)
	first := mergeSort(arr[0 : l/2])
	second := mergeSort(arr[l/2:])
	return merge(first, second)

}
