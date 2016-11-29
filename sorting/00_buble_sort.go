package sorting

import "github.com/eljuanchosf/adventures-in-golang/shared"

//BubleSort sorts an array of integers by swapping two elements at once, iteratively
func BubleSort(array []int) []int {
	elementCount := len(array) - 1
	for {
		swapped := false
		for elementA := 0; elementA < elementCount; elementA++ {
			elementB := elementA + 1
			if array[elementA] > array[elementB] {
				swapped = true
				shared.SwapInt(array, elementA, elementB)
			}
		}
		if !swapped {
			break
		}
	}
	return array
}
