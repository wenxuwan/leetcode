package main

import "fmt"

func isMatch(s string, p string) bool {
	result := make([][]bool, len(s)+1)
	for index, _ := range result {
		result[index] = make([]bool, len(p)+1)
	}
	result[0][0] = true
	for i := 1; i <= len(p); i++ {
		result[0][i] = result[0][i-1] && p[i-1] == '*'
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				result[i][j] = result[i-1][j-1]
			} else {
				if p[j-1] == '*' {
					if result[i][j-1] || result[i-1][j] {
						result[i][j] = true
					}
				}
			}
		}
	}
	return result[len(s)][len(p)]
}

func main() {
	fmt.Println(isMatch("a", "a**"))
}
