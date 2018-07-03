package main
/*
最长回文字符串，可以参照下面的网址的讲解
https://www.felix021.com/blog/read.php?2040


 */
import "fmt"

func longestPalindrome(s string) string {
	arryLen := len(s)
	array := make([]byte,(arryLen * 2 + 1),(arryLen * 2 + 1))
	j := 0
	i := 0
	for i = 0;i < len(array);i++{
		if i % 2 == 0{
			array[i] = '#'
		}else{
			array[i] = s[j]
			j++
		}
	}
	fmt.Println(array)
	p := make([]int,(arryLen * 2 + 1),(arryLen * 2 + 1))
	rightMaxFlag := 0
	maxIndex := 0
	maxNum := 0
	xiaobiaoflag := 0
	for i:= 0; i < len(array);i++{
		if rightMaxFlag > i{
			if rightMaxFlag - i <= p[2 * maxIndex - i]{
				p[i] = rightMaxFlag - i
			}else{
				p[i] = p[2 * maxIndex - i]
			}
		}else{
			p[i] = 1
		}
		for (i - p[i] >=0) && (i + p[i] < len(array)) && array[i + p[i]] == array[i - p[i]]{
			p[i]++
		}
		if p[i] > maxNum{
			maxNum = p[i]
			xiaobiaoflag = i
		}

		if i + p[i] > rightMaxFlag{
			rightMaxFlag = i + p[i]
			maxIndex = i
			//fmt.Println(i,rightMaxFlag,maxIndex)
		}
	}
	begin := 0
	fmt.Println(xiaobiaoflag,maxNum)
	begin =  (xiaobiaoflag - maxNum +1)/2
	return string(s[begin:begin + maxNum-1])
}

func longestPalindromeLujingGuiHua(s string) string {
	len1 := len(s)
	result := make([][]bool,len1,len1)
	for i := range result {
		result[i] = make([]bool, len1)
	}
	maxlen := 0
	beginIndex := 0
	for i := 0; i < len(s);i++{
		for j := 0; j <= i;j++{
			if i - j < 2{
				result[j][i] = (s[i]==s[j])
			}else {
				if result[j+1][i-1] == true && s[i] == s[j] {
					result[j][i] = true
				} else {
					result[j][i] = false
				}
			}
				len := i - j + 1
				if len > maxlen && result[j][i] {
					maxlen = len
					beginIndex = j
				}

		}

	}
	fmt.Println(beginIndex,maxlen)
	return s[beginIndex:beginIndex+maxlen]
}

func main(){
	testData := "a"
	resultData := longestPalindromeLujingGuiHua(testData)
	fmt.Println(resultData)
}
