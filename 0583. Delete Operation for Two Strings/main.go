package main

func minDistance(word1 string, word2 string) int {
	result := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		result[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				result[i+1][j+1] = result[i][j] + 1
			} else {
				if result[i][j+1] > result[i+1][j] {
					result[i+1][j+1] = result[i][j+1]
				} else {
					result[i+1][j+1] = result[i+1][j]
				}
			}
		}
	}
	return len(word2) + len(word1) - 2*(result[len(word1)][len(word2)])
}

func main() {

}
