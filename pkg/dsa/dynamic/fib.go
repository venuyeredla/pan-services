package dynamic

// fib(0)=0,fib(1)=1
// f(n)=f(n-1)+fib(n-2) return nth term in sequence.
func fib(n int, cache []int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fib(n-1, cache) + fib(n-2, cache)
	}
}

func fibC(n int) int {
	cache := make([]int, 7)
	cache[0] = 0
	cache[1] = 1
	for i := 2; i < len(cache); i++ {
		cache[i] = -1
	}
	return fib(n, cache)
}

func FibSeries(size int) []int {
	// Count series from zero.
	fib := make([]int, size)
	var i int = 0
	fib[0] = 0
	fib[1] = 1
	for i < size {
		fib[i] = fib[i-1] + fib[i-2]
		i++
	}
	return fib
}

func FibNth(n int) int {
	if n < 2 {
		return n
	} else {
		return FibNth(n-2) + FibNth(n-1)
	}

}
