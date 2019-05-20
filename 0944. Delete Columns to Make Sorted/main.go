package main

import "fmt"

/*
GO语言的代码提交，题目总会出问题。所以写个C的为了pass

int minDeletionSize(char ** A, int ASize){
	int len,i,j = 0;
	int deletNum = 0;
	len = strlen(A[0]);
	for(j = 0; j < len; j++){
		for(i = 1;i < ASize;i++){
			if(A[i-1][j] > A[i][j]){
				deletNum++;
				break;
			}
		}
	}
	return deletNum;
}

*/
func minDeletionSize(A []string) int {
	dlenth := 0
	if len(A) > 100 || len(A) < 1{
		return dlenth
	}
	if len(A[0]) < 1 || len(A[0]) > 1000{
		return dlenth
	}
	for i := 0; i < len(A[0]);i++{
		for j := 1; j < len(A);j++{
			if A[j - 1][i] > A[j][i]{
				dlenth++
				break
			}
		}
	}
	return  dlenth
}

func main(){
	A := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(A))
}
