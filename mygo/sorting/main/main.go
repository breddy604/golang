package main

import "fmt"

func main() {
	input := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}

	heap(input)

	fmt.Println(input)

}

func heap(arr []int) {

	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr)-1)
	}

	lastIndex := len(arr) - 1

	for lastIndex != 0 {
		arr[0], arr[lastIndex] = arr[lastIndex], arr[0]
		lastIndex = lastIndex - 1
		heapify(arr, 0, lastIndex)
	}

}

func heapify(arr []int, index int, endIndex int) {

	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left <= endIndex && arr[largest] < arr[left] {
		largest = left
	}
	if right <= endIndex && arr[largest] < arr[right] {
		largest = right
	}

	if index != largest {
		arr[index], arr[largest] = arr[largest], arr[index]
		heapify(arr, largest, endIndex)
	}

}

func selection(arr []int) {

	l := len(arr)

	for i := 0; i < l; i++ { // 1 2 3 4 5 6 7 8
		minIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

}

func insertion(arr []int) {
	for i := 1; i < len(arr); i++ {
		for i > 0 && arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i = i - 1
		}
	}
}

func quicksort(arr []int) {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return
	}
	left, right := partition(arr)
	fmt.Println(left, right)
	quicksort(left)
	quicksort(right)
}

func partition(arr []int) ([]int, []int) {

	pivot := len(arr) - 1

	i := -1
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < arr[pivot] {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i = i + 1
	arr[i], arr[pivot] = arr[pivot], arr[i]
	return arr[:i], arr[i+1:]

}

func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	result := 0
	size := int32(len(expenditure))
	if size <= d {
		return 0
	}

	part := expenditure[0:d]
	sorted := mergeSort(part)

	for i := 0; int32(i) < size-d; i++ {

		if len(sorted)%2 == 0 {
			index := (len(sorted) - 1) / 2
			median := float64((sorted[index] + sorted[index+1])) / float64(2.0)
			fmt.Println(median, expenditure[int32(i)+d])
			if float64(expenditure[int32(i)+d]) >= 2*median {
				result = result + 1
			}
		} else {
			index := (len(sorted) - 1) / 2
			median := sorted[index]
			fmt.Println(median, expenditure[int32(i)+d])
			if expenditure[int32(i)+d] >= 2*median {
				result = result + 1
			}
		}

		sorted[0] = expenditure[d+int32(i)]
		sorted = merge(sorted[0:1], sorted[1:d-1])
	}

	return int32(result)

}

func mergeSort(arr []int32) []int32 {
	if len(arr) == 1 {
		return arr
	}

	first := mergeSort(arr[0 : len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

func merge(arr1 []int32, arr2 []int32) []int32 {
	i, j, k := 0, 0, 0
	tmp := make([]int32, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			tmp[k] = arr1[i]
			i = i + 1
		} else {
			tmp[k] = arr2[j]
			j = j + 1
		}
		k = k + 1
	}

	if i == len(arr1) {
		for j < len(arr2) {
			tmp[k] = arr2[j]
			k = k + 1
			j = j + 1
		}
	}
	if j == len(arr2) {
		for i < len(arr1) {
			tmp[k] = arr1[i]
			k = k + 1
			i = i + 1
		}
	}
	return tmp
}
