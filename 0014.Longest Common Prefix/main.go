package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	commonstr := strs[0]
	var strlen int = 0
	var maxlen int = len(commonstr)
	for _, strvalue := range strs {
		var commonlen = 0
		if len(commonstr) < len(strvalue) {
			strlen = len(commonstr)
		} else {
			strlen = len(strvalue)
		}
		for j := 0; j < strlen; j++ {
			if strvalue[j] == commonstr[j] {
				commonlen++
			} else {
				break
			}
		}

		if commonlen <= maxlen {
			maxlen = commonlen
			commonstr = commonstr[:maxlen]
		}

	}

	return commonstr
}
func main() {
	var strs []string
	fmt.Println(longestCommonPrefix(strs))
}
