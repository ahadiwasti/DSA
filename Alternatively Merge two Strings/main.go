package main

func main() {
	word1 := "abc"
	word2 := "pqr"
	result := mergeAlternately(word1, word2)
	println(result)
}

func mergeAlternately(word1 string, word2 string) string {
	n := len(word1) + len(word2)
	result := make([]byte, 0, n)
	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			result = append(result, word1[i])
			i++
		}
		if j < len(word2) {
			result = append(result, word2[j])
			j++
		}
	}
	return string(result)

}
