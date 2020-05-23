package main

import "fmt"

func main() {
	randomNumbers := numbers{1, 2, 3}

	randomNumbers.printAll()
	fmt.Println(randomNumbers)
	fmt.Println(randomNumbers.getNumbers())
}
