package recursion

func Converge(n int) int {
	if n == 0 {

		return 1

	} else {

		return n + Converge(n-1)

	}
}
