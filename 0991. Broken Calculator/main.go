package main

import "fmt"

func brokenCalc(X int, Y int) int {
	if X >= Y{
		return X - Y
	}
	if Y % 2 == 0{
		return 1 + brokenCalc(X, Y / 2)
	}else{
		return 1 + brokenCalc(X, Y + 1)
	}
}

func main(){
	fmt.Println(brokenCalc(4,10))
}
