//这道题目实在没必要写注解了。。。
package main

func findLUSlength(a string, b string) int {
	if a == b{
		return -1
	}
	if len(a) > len(b){
		return len(a)
	}
	return len(b)
}

func main(){

}
