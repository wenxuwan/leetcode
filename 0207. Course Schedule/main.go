/*
拓扑排序,可以用bfs。这类问题就是依赖解决问题，首先从入度为0的点出发，
然后把所有以它为起点的点的入度都减一。然后以此循环。
*/
package main

import "fmt"

func bfs(m map[int]int,key int,p [][]int){
	for i :=0;i < len(p);i++{
		if p[i][1] == key{
			m[p[i][0]]--
			if m[p[i][0]] == 0{
				m[p[i][0]] = -1
				bfs(m, p[i][0],p)
			}
		}
	}
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	depent := make(map[int]int,numCourses)
	for i := 0; i < numCourses;i++{
		depent[i] = 0
	}
	for i := 0 ;i < len(prerequisites);i++{
		depent[prerequisites[i][0]]++
	}
	fmt.Println(depent)

	for key, value := range depent{
		if value == 0{
			bfs(depent, key, prerequisites)
		}
	}
	fmt.Println(depent)
	for i := 0 ;i < numCourses;i++{
		if depent[i] > 0{
			return false
		}
	}
	return true
}

func main(){
	testData := [][]int{{1,0},{2,1}}
	fmt.Println(canFinish(3,testData))
}
