/*
结果肯定是给的数组里面的某个字符串，要做的就是遍历数组
然后看a[i]和其它字符串比是不是其它字符串的字串，如果是的话
那么肯定不是它了。如果不是可能会是它，最后求出来所有可能会是的，最长的就是结果
*/
package main

import "fmt"

func helper(s1, s2 string)int{
	len1 := len(s1)
	len2 := len(s2)
	if len1 > len2{
		return len1
	}
	index1 :=0
	index2 := 0
	for index1 = 0; index1 < len1;index1++{
		for index2 < len2 && s1[index1] != s2[index2]{
			index2++
		}
		if index2 == len2{
			break
		}
		index2++
	}
	if index1 == len1{
		return -1
	}
	return len1
}
func findLUSlength(strs []string) int {
	lenth := len(strs)
	result := -1
	tmp := -1
	for i :=0; i < lenth;i++{
		for j := 0; j < lenth;j++{
			if i == j{
				continue
			}
			tmp = helper(strs[i], strs[j])
			if tmp == -1{
				break
			}
		}
		if tmp > result{
			result = tmp
		}
	}
	return result
}

func main(){
	test := []string{"abcabc","abcabc","abcabc","abc","abc","cca"}
	fmt.Println(findLUSlength(test))
}
