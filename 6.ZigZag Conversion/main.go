package main

import "fmt"

/*
	该题主要是找规律：
		Input: s = "PAYPALISHIRING", numRows = 4
		Output: "PINALSIGYAHRPI"
		Explanation:

		P     I    N
		A   L S  I G
		Y A   H R
		P     I

以这个example为例，4行竖着放,其实就是个倒Z,我们可以看到每行的每个元素之间的关系如下：
	index + (2 * numRows - (2 * (i + 1))) //中间隔着的坐标( i [0,numRows))
	index + (2 * numRows -2) //对称的变量的坐标

 */
func convert(s string, numRows int) string {
	result := make([]byte,len(s))
	j := 0
	index := 0
	for i := 0; i < numRows && i < len(s);i++{
		index = i
		result[j] = s[index]
		j++
		for ;index < len(s);{
			nextIndex := index + (2 * numRows - (2 * (i + 1))) //中间的变量坐标
			nextIndex2 := index + (2 * numRows -2) //对称的变量的坐标
			if nextIndex == index && nextIndex2 == index{
				result = []byte(s)
				break
			}
			if nextIndex < len(s) && nextIndex != index{
				fmt.Println(j,nextIndex)
				result[j] = s[nextIndex]
				j++
				index = nextIndex
			}
			if nextIndex2 < len(s) && nextIndex2 != index{
				if nextIndex == nextIndex2{
					// Do nothing
				}else{
					result[j] = s[nextIndex2]
					j++
					index = nextIndex2
				}
			}
			index = nextIndex2
		}
	}
	return string(result)
}


func main(){
	s := "AB"
	numRows := 1
	ss := convert(s,numRows)
	fmt.Println(ss)
}
