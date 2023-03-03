package factorial

func Factorial(n int) int {
	ans := 1
	if n == 0 {
		return 1
	} else {
		for i := 1; i <= n; i++ {
			ans *= i
		}
	}
	return ans
}
