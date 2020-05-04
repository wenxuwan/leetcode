package main

import "fmt"

func isUgly(num int) bool {
	value := []int{5,3,2}
	for num != 1{
		tmp := num
		for _,v := range value{
			if num % v == 0{
				num = num / v
			}
		}
		if tmp == num{
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(isUgly(7))
}
