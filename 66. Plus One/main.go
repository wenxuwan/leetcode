/*简单的数字操作，保证最前面一位如果进位了，那就要在最前面添加1*/
package main

import "fmt"

func plusOne(digits []int) []int {
	jinwei := 0
	a := make([]int,1,len(digits))
	a[0] = 1
	i := 0
	for i = len(digits) - 1; i >= 0; i--{
		if digits[i] < 9{
			digits[i] = digits[i] + 1
			jinwei = 0
			break
		}else{
			jinwei = 1
			digits[i] = 0
		}
	}
	if i < 0 && jinwei == 1{
		digits = append(a,digits[0:]...)
	}
	return  digits
}

func main()  {
	arr := []int{9,9,9}
	fmt.Print(plusOne(arr))
}
