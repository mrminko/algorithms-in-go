package main

import "fmt"

func InsertionSort(keys []int) {
	for i := 1; i < len(keys); i++ {
		temp := keys[i]
		j := i - 1
		for j >= 0 && keys[j] < temp {
			keys[j+1] = keys[j]
			j--
		}
		keys[j+1] = temp
	}
}

func main() {
	keys := []int{8, 9, 47, 35, 1, 2, 0, 84, 53, 2, 12}
	InsertionSort(keys)
	fmt.Println("Sorted:\n\r", keys)
}
