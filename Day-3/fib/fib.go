package fib

func CreateFib(s int) []int {
	a, b := 0, 1
	f := []int{0, 1}
	for i := len(f); i < s; i++ {
		a, b = b, a+b
		f = append(f, b)
	}
	return f

}
