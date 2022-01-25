package lib

func Fib(x int) int {
	if x == 0 {
		return 1
	}
	return Fib(x-1) + Fib(x-2)
}
