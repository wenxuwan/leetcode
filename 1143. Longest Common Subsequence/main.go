package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	result := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		result[i] = make([]int, len(text2)+1)
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				result[i+1][j+1] = result[i][j] + 1
			} else {
				if result[i+1][j] > result[i][j+1] {
					result[i+1][j+1] = result[i+1][j]
				} else {
					result[i+1][j+1] = result[i][j+1]
				}
			}
		}
	}

	return result[len(text1)][len(text2)]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}
