package ggvm_util

func Fib(x int) func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
