package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("hello"))    // holle
	fmt.Println(reverseVowels("leetcode")) // leotcede
}

func reverseVowels(s string) string {
	vowelarry := []rune(s)
	left, right := 0, len(s)-1

	isVowel := func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	for left < right {
		if !isVowel(vowelarry[left]) {
			left++
			continue
		}
		if !isVowel(vowelarry[right]) {
			right--
			continue
		}
		vowelarry[left], vowelarry[right] = vowelarry[right], vowelarry[left]
		left++
		right--
	}
	return string(vowelarry)
}
