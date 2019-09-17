/*
题目看着很麻烦，但其实可以逐步分解掉，首先如果没有*号的话
我们很容易写，一个个判断，但如果有*号的时候会出现两种情况：
1. *号前面的字符可以和S匹配，这时候我们只需要S下标往后移，因为*可以代表N（0个或无数个）前面的字符。
2. *号前面的字符不能匹配，这时候无所谓，因为我们可以去掉它，所以P的下表往后移动两个。
这样逐步分解之后就可以写出来了。
*/

package main

func isMatch(s string, p string) bool {
	withoutAsterisk := false
	if len(p) == 0{
		if len(s) == 0{
			return true
		}else{
			return false
		}
	}
	if len(s) != 0 && (s[0] == p[0] || p[0] == '.'){
		withoutAsterisk = true
	}

	if len(p) >= 2 && (p[1] == '*'){
		return (isMatch(s,p[2:]) || withoutAsterisk && isMatch(s[1:],p))
	}else{
		return (withoutAsterisk && isMatch(s[1:],p[1:]))
	}
}

func main(){

}