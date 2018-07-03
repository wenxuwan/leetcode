package main

import (
	"strconv"
	"fmt"
)

/***
这个题目真的是智障一般的水平，完全没搞懂它的要求，一次次的提交来证明它的要求
***/
func myAtoi(str string) int {
	resultValue := 0
	flagValue := 0 //0代表正数
	positiveMax  := 2147483647
	positiveMaxStr := strconv.Itoa(positiveMax)
	negativeMaxStr := "2147483648"
	beginFlag := 0
	resultSlice := make([]byte,0,len(str))
	for i := 0;i < len(str);i++{
		if str[i] == ' '{
			if beginFlag == 0 {
				continue
			}else{
				return 0
			}
		}else {
			if '1' <= str[i] && str[i] <= '9' || str[i] == '-' || str[i] == '+'{
				if '1' <= str[i] && str[i] <= '9' {
					for j := i; j < len(str); j++ {
						if str[j] < '0' || str[j] > '9'{
							break
						}else{
							resultSlice = append(resultSlice,str[j])
							//fmt.Println(str[j])
							//fmt.Println(string(resultSlice))
							//fmt.Println(positiveMaxStr)
							if len(resultSlice) > len(negativeMaxStr) || len(resultSlice) == len(negativeMaxStr) && string(resultSlice) > positiveMaxStr{
								//fmt.Println("Hello")
								resultValue = positiveMax
								break
							}else {
								resultValue = resultValue * 10 + int(str[j] - '0')
							}
						}
					}
				}else{
					if(beginFlag == 1){
						return 0
					}
					if str[i] == '-'{
						flagValue = 1
						fmt.Println("负数")
					}
					flag := 0
					for j := i + 1; j < len(str); j++{

						if str[i + 1] == '-' || str[i + 1] == '+'{
							return 0
						}
						if flag == 0 && str[j] == '0'{
							continue
						}
						flag = 1
						if '0' <= str[j] && str[j] <= '9'{
							resultSlice = append(resultSlice,str[j])
							fmt.Println(len(resultSlice),len(negativeMaxStr))
							if len(resultSlice) > len(negativeMaxStr) || len(resultSlice) == len(negativeMaxStr) && string(resultSlice) > positiveMaxStr {
								if flagValue == 1{
									resultValue = -2147483648
								}else{
									resultValue = 2147483647
								}
								break
							}else{
									resultValue = resultValue * 10 + int(str[j] - '0')
							}
						}else{
							break
						}
					}
				}
				break
			}else{
				if str[i] == '0'{
					beginFlag = 1
					continue
				}else {
					return 0
				}
			}
		}
	}
	if flagValue == 1{
		if resultValue == -2147483648{
			return resultValue
		}
		resultValue = -resultValue
	}

	return resultValue
}

func main(){
	str :=   "0-1"
	str2 := "2147483648"
	if str < str2{
		//fmt.Println("str xiaoyu str2")
	}else{
		//fmt.Println("str bigger than str2")
	}
	fmt.Println(myAtoi(str))
}
