package main

import "fmt"

func groupAnagrams(strs []string) [][]string {

	if len(strs) == 0{
		return nil
	}
	result := make([][]string,0,len(strs))
	//s1 := make([]string,0,len(strs))
	str := make([]string,0,len(strs))
	ValueMap := make(map[string][]string)
	for _, value := range strs{
		valueByte := []byte(value)
		for i := 0; i < len(valueByte) - 1;i++{
			flag := true
			for j :=0; j < len(valueByte) - i - 1;j++{
				if valueByte[j] > valueByte[j+1]{
					flag = false
					valueByte[j],valueByte[j+1] = valueByte[j+1],valueByte[j]
				}
			}
			if flag{
				break
			}
		}
		str = append(str, string(valueByte[:]))
		ValueMap[string(valueByte[:])] = append(ValueMap[string(valueByte[:])], value)
	}
	for _,v := range  ValueMap{
		result = append(result, v)
	}
	return  result
}

func main(){
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}