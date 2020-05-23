package main

import "fmt"

type numbers []int

func (nums numbers) printAll() {
	for i, num := range nums {
		fmt.Println("Index", i, "has number is", num)
	}
}

func (nums numbers) getNumbers() numbers {
	return nums
}
