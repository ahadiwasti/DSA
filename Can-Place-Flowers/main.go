package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPlaceFlowers([]int{0, 1, 0, 0, 0, 1}, 0, 2))    // false
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 0, 2)) // true
}

func canPlaceFlowers(bed []int, i, n int) bool {
	if n == 0 {
		return true
	}
	if i >= len(bed) {
		return false
	}

	if bed[i] == 1 {
		return canPlaceFlowers(bed, i+2, n)
	}

	leftOK := i == 0 || bed[i-1] == 0
	rightOK := i == len(bed)-1 || bed[i+1] == 0

	if leftOK && rightOK {
		bed[i] = 1
		if canPlaceFlowers(bed, i+2, n-1) {
			bed[i] = 0
			return true
		}
		bed[i] = 0
	}
	return canPlaceFlowers(bed, i+1, n)
}
