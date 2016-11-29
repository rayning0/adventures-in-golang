package sorting

import "github.com/eljuanchosf/adventures-in-golang/shared"

//SelectiveSort iteratively finds the smallest element and swaps it for the first element of the subarray
func SelectiveSort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	for startIndex := 0; startIndex < len(array); startIndex++ {
		position := shared.FindSmallestElement(array, startIndex, len(array)-1)
		if array[startIndex] > array[position] {
			shared.SwapInt(array, position, startIndex)
		}
	}
	return array
}
