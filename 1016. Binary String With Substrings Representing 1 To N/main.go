//转成2进制判断一下是不是子字符串就行了
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertToBin(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	for ;num > 0 ; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}


func queryString(S string, N int) bool {
	for i := 1; i <= N;i++{
		binary := convertToBin(i)
		if !strings.Contains(S, binary){
			return false
		}
	}

	return true
}

func main(){
	fmt.Println(queryString("110101011011000011011111000000", 15))
}

