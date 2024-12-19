package main

import "fmt"

func main() {
	lookFor := 6
	sortedList := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	index := binarySearch(sortedList, lookFor)
	if index > 0 {
		fmt.Println("Found at: ", index)
	} else {
		fmt.Println("Value not found")
	}
}

func binarySearch(sortedList []int, lookFor int) int {
	// implement two pointers to track sub-arrays instead of slicing the original one
	// search midpoint,
	// equal => return, lessThan => move left pointer, greaterThan => move right pointer
	leftIndex := 0
	rightIndex := len(sortedList) - 1
	for leftIndex <= rightIndex {
		midpoint := leftIndex + ((rightIndex - leftIndex) / 2)
		midValue := sortedList[midpoint]
		if midValue == lookFor {
			return midpoint
		}
		if midValue < lookFor {
			leftIndex = midpoint + 1
		}
		if midValue > lookFor {
			rightIndex = midpoint - 1
		}
	}
	return -1
}
