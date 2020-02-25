package main

func isSubsequence(s string, t string) bool {
	begin := 0
	for _,value := range s{
		if begin >= len(t){
			return false
		}
		for int32(t[begin]) != value{
			begin++
			if begin >= len(t){
				return false
			}
		}
		begin++
	}
	return true
}

func main(){

}
