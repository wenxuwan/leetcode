package main

func merge(A []int, m int, B []int, n int)  {
	tlen :=  m + n
	begin1 := m - 1
	begin2 := n - 1
	flag := tlen - 1
	for begin1 >= 0 && begin2 >= 0{
		if A[begin1] > B[begin2]{
			A[flag] = A[begin1]
			begin1--
		}else{
			A[flag] = B[begin2]
			begin2--
		}
		flag--
	}
	for begin2 >= 0{
		A[flag] = B[begin2]
		begin2--
		flag--
	}
}

func main(){

}
