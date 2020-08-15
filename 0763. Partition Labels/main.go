package main

func partitionLabels(S string) []int {
	lens := len(S)
	end := lens - 1
	begin := 0
	result := make([]int, 0, 1)
	bianli := make(map[uint8]int)
	for begin <= end {
		for S[begin] != S[end] {
			if _, ok := bianli[S[end]]; !ok {
				bianli[S[end]] = end
			}
			end--
		}
		tmpBegin := begin
		for begin < end {
			if value, ok := bianli[S[begin]]; ok {
				if value > end {
					end = value
				}
			}
			begin++
		}
		result = append(result, end-tmpBegin+1)
		if end == lens-1 {
			break
		}
		begin = end + 1
		end = lens - 1
	}
	return result
}
