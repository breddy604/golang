package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks int
}

type byAge []Student
type byName []Student

func (l byAge) Less(i int, j int) bool {
	return l[i].age < l[j].age
}

func (l byAge) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byAge) Len() int {
	return len(l)
}

func (l byName) Less(i int, j int) bool {
	return l[i].name < l[j].name
}

func (l byName) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byName) Len() int {
	return len(l)
}

type Contry struct {
	name   string
	states []State
}

type State struct {
	name      string
	districts []District
}

type District struct {
	name    string
	persons []Person
}

type Person struct {
	aadharId string
	gender   string
	age      int
	district *District
}

type SubscriberIntf interface {
	subscribe(string, interface{})
}

type MessagingMiddlewareIntf interface {
	subscribe(string, SubscriberIntf)
}

func main() {

	tree := []int{9, 5, 10, 6, 18, 4, 3}

	for i := 0; ; {

	}
}

func sort(arr []int) {

	p := partition(arr)
	if p == -1 {
		return
	}
	sort(arr[:p])
	sort(arr[p+1:])

}

func partition(arr []int) int {
	if len(arr) <= 1 {
		return -1
	}
	i := -1
	k := len(arr) - 1
	visited := make([]int, 0, 10)
	fmt.Println("Original", arr)
	for j := 0; j != k; j++ {
		if contains(arr[j], visited) {
			arr = append(arr[:j], arr[j+1:]...)
			k = k - 1
			j = j - 1
			fmt.Println("Duplicates Removed", arr)
			continue
		}
		if arr[j] <= arr[k] {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
		visited = append(visited, arr[j])
	}
	arr[i+1], arr[k] = arr[k], arr[i+1]

	return i + 1
}

func contains(a int, arr []int) bool {
	for _, e := range arr {
		if a == e {
			return true
		}
	}

	return false
}
