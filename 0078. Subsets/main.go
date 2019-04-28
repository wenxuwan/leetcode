package main
import "fmt"

func dfs(result* [][]int, current int, tmp []int, data []int){

	*result = append(*result, tmp)

	for i := current ; i < len(data);i++{
		tmpSlice := make([]int,len(tmp))
		copy(tmpSlice, tmp)
		tmpSlice  = append(tmpSlice,data[i])
		dfs(result,i + 1,tmpSlice,data)
	}
}
func subsets(nums []int) [][]int {
	var tmp []int
	var result [][]int
	dfs(&result,0, tmp, nums)
	return result
}

func main(){
	s := []int{1,2,3}
	fmt.Println(subsets(s))
}