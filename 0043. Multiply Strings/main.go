package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	flag := 0
	var result string
	if num2 == "0" || num1 == "0"{
		return "0"
	}
	for index2 := len2 - 1; index2 >= 0;index2--{
		flag++
		tmp2 := int(num2[index2] - '0')
		var jinweiValue int
		tmpResult := ""
		for index1 := len1 - 1;index1 >= 0;index1--{
			tmp1 := int(num1[index1] - '0')
			tmp :=  tmp1 * tmp2
			tmp,  jinweiValue = (tmp + jinweiValue) % 10, (tmp + jinweiValue) / 10
			tmpResult = strconv.FormatInt(int64(tmp), 10) + tmpResult
			//fmt.Println(tmpResult, result)
		}
		if jinweiValue != 0{
			tmpResult = strconv.FormatInt(int64(jinweiValue), 10) + tmpResult
		}
		if result == ""{
			result = tmpResult
			continue
		}
		begin := len(result) - flag
		test := begin
		var jin byte
		tmpResult2 := ""
		for j := len(tmpResult) - 1; j >= 0 ; j--{
			if begin >= 0{
				tmpResult2 = strconv.FormatInt(int64(((result[begin] + tmpResult[j] - 2 * '0') + jin)  % 10), 10) + tmpResult2
				jin = ((result[begin] + tmpResult[j] - 2 * '0') + jin)  / 10
				begin--
			}else{
				tmpResult2 = strconv.FormatInt(int64((tmpResult[j] - '0' + jin) % 10), 10) + tmpResult2
				jin = (tmpResult[j] - '0' + jin)  / 10
			}
			//fmt.Println(tmpResult2)
		}
		if jin != 0{
			tmpResult2 =  strconv.FormatInt(int64(jin), 10) + tmpResult2
			jin = 0
		}
		tmpResult2 = tmpResult2 + result[test + 1:]
		result = tmpResult2
	}
	return result
}

func main(){
	fmt.Println(multiply("0", "456"))
}