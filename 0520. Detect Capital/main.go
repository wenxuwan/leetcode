package main

import "fmt"

func detectCapitalUse(word string) bool {
	begin := 0
	totle := 0
	if word[0] >= 65 && word[0] <= 90{
		begin = 1
	}
	for i := 1; i < len(word);i++{
		if word[i] >= 65 && word[i] <= 90{
			totle++
		}
		if totle < i && begin == 1 && totle != 0{
			return false
		}
	}
	if totle != 0 && begin == 0{
		return false
	}
	return true
}

func main(){
	fmt.Println(detectCapitalUse("FlaG"))
}
