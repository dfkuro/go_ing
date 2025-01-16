package main

import "fmt"

func main() {
	// Create the array
	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	searchNumber := 23
	fmt.Println(binarySearch(searchField, searchNumber))
}

func binarySearch(searchField []int, searchNumber int) int {
	fmt.Printf("Searching %v on %v \n", searchNumber, searchField)
	low := 0
	high := len(searchField) - 1

	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("low: %v, mid: %v, high: %v\n", low, mid, high)

		if searchField[mid] == searchNumber {
			fmt.Println("Found")
			return mid
		} else if searchField[mid] < searchNumber {
			fmt.Println("Not found 1")
			low = mid + 1
		} else {
			fmt.Println("Not found -1")
			high = mid - 1
		}
	}
	return -1
}
