package main

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	p, q := false, false
	f := func() int {
		if !p {
			p = true
			return x
		}
		if !q {
			q = true
			return y
		}
		x, y = y, x+y
		return y
	}
	return f
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
