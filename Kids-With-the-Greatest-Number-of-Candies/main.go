package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := maxCandy(candies, len(candies))
	var helper func(i int) []bool
	helper = func(i int) []bool {
		if i == len(candies) {
			return []bool{}
		}
		curr := (candies[i] + extraCandies) >= maxCandies
		return append([]bool{curr}, helper(i+1)...)
	}

	return helper(0)
}

func maxCandy(candies []int, n int) int {
	if n == 1 {
		return candies[0]
	}
	max := maxCandy(candies, n-1)
	if candies[n-1] > max {
		return candies[n-1]
	}
	return max
}
