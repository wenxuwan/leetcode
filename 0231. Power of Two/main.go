package main

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	for n != 0 {
		yu := n % 2
		shang := n / 2
		if yu != 0 {
			return false
		}
		if shang == 1 {
			return true
		}
		n = shang
	}
	return false
}

func main() {

}
