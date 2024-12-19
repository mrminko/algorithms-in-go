package main

import "fmt"

func main() {
	inputs := []int{5, 4, 3, 2, 1, 0}
	N := len(inputs)
	for i := 0; i < N; i++ {
		didSwap := sweep(inputs, i)
		if !didSwap {
			break
		}
	}
	fmt.Println("After all sweep", inputs)
}

func sweep(numbers []int, prevIndex int) (didSwap bool) {
	i1 := 0
	i2 := 1
	N := len(numbers)
	didSwap = false
	for i2 < (N - prevIndex) {
		if numbers[i1] > numbers[i2] {
			numbers[i1], numbers[i2] = numbers[i2], numbers[i1]
			didSwap = true
		}
		i1++
		i2++
	}
	return didSwap
}
