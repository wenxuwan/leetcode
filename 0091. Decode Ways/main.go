package main

import "fmt"

func numDecodings(s string) int {
	lenth := len(s)
	result := make([]int,lenth + 1,lenth + 1)
	if('0' < s[0] && s[0] <= '9'){
		result[0] =  1
		result[1] =  1
	}else{
		return 0
	}
	for i := 2;i < lenth + 1;i++{
		fmt.Println(result)
		if s[i - 1] == '0'{
			//如果前面是0，则退出
			if s[i - 2] == '0' && i != 2{
				return 0
			}else{
				if s[i - 2] <= '2'{
					result[i] = result[i - 2]
				}else{
					return 0
				}
			}
		}else{
			//前面等于0那么就不能再重新拆分组合了
			if s[i - 2] == '0' && i != 2{
				result[i] = result[i - 1]
			}else{
				if (s[i - 2] == '2' && s[i - 1] <= '6') || (s[i - 2] == '1'){
					result[i] += result[i - 1]
					result[i] += result[i - 2]
				}else{
					result[i] += result[i - 1]
				}
			}
		}
	}
	fmt.Println(result)
	return result[lenth]
}

func main(){
	fmt.Println(numDecodings("17"))
}
