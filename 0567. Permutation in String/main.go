/*
这个问题做法应该是比较多的，可以取一个map以s1中的
字符作为key然后value作为出现的个数，然后就是遍历
s2，只要出现不一样的，就可以直接跳到下一个开始了。
但如果发现出现一个前面出现过的，那么前面的指针就可以
往后移动到其第一次出现的地方了。
*/
package main

import "fmt"

func mapCopy(m1 map[int32]int, m2 map[int32]int){
	for key,value := range m1{
		m2[key] = value
	}
}
func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[int32]int)
	begin := 0
	for _,value := range(s1){
		s1Map[value]++
	}
	tmp := make(map[int32]int)
	mapCopy(s1Map, tmp)
	for index, value := range(s2){
		if _, ok := s1Map[value]; ok{
			if _, ok := tmp[value]; ok{
				if tmp[value] > 0{
					tmp[value]--
				}
				if tmp[value] == 0{
					delete(tmp, value)
				}
			}else{
				for begin < index{
					if int32(s2[begin]) != value{
						tmp[int32(s2[begin])]++
					}else{
						begin++
						break
					}
					begin++
				}
			}
		}else{
			begin = index + 1
			mapCopy(s1Map, tmp)
		}
		if len(tmp) == 0{
			return true
		}
	}
	return false
}

func main(){
	fmt.Println(checkInclusion("adc","dcda"))
}
