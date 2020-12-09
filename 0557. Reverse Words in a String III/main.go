package main

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reverseWords(s string) string {
	var result string
	begin := 0
	for end, value := range s {
		if value == ' ' {
			result += reverse(s[begin:end]) + " "
			begin = end + 1
		}
	}
	result += reverse(s[begin:len(s)])
	return result
}

func main() {

}
