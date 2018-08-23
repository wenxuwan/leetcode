package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) == 0{
		return true
	}
	runes := []rune(s)
	begin := 0
	end := len(s) - 1
	for begin < end{
		if ('0' <= runes[begin] && runes[begin]<= '9') || ('a' <= runes[begin] && runes[begin] <= 'z') || ('A' <= runes[begin] && runes[begin] <= 'Z'){
			if ('0' <= runes[end] && runes[end]<= '9') || ('a' <= runes[end] && runes[end] <= 'z') || ('A' <= runes[end] && runes[end] <= 'Z'){
				if runes[begin] != runes[end]{
					if (('a' <= runes[end] && runes[end] <= 'z') || ('A' <= runes[end] && runes[end] <= 'Z')) && (('a' <= runes[begin] && runes[begin] <= 'z') || ('A' <= runes[begin] && runes[begin] <= 'Z')){
						if runes[begin] - runes[end] != 32 && runes[end] - runes[begin] != 32{
							return false
						}else{
							begin++
							end--
						}
					}else{
						return false
					}
				}else{
					begin++
					end--
				}
			}else{
				end--
			}
		}else{
			begin++
		}
	}
	return true
}

func main(){
	ab := "0P"
	fmt.Println(isPalindrome(ab))
}