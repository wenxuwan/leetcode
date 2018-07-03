package main

import "strconv"

func searchIpDp(s string, n int) {
	var result []string
	for i := 0; i < 3; i++ {
		intvalue, err := strconv.Atoi(s[0:i])
		if intvalue < 255 && s[0] != 0 {
			searchIpDp(s[i:])
			return result
		}

	}
}
func restoreIpAddresses(s string) []string {
	var int n = 0
	searchIpDp(s, n)
}
func main() {

}
