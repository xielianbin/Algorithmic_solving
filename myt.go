package main

import "fmt"

func t1() {
	n := 9
	var ff func(a int)
	ff = func(a int) {
		fmt.Printf("ff a=%d\n", a)
		if a > 0 {
			ff(a - 1)
		}
	}
	ff(n)

}

func main() {
	t1()
}
