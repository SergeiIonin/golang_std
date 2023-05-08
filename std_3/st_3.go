package main

import "fmt"

func main() {
	fmt.Println("---RECURSION---")
	fact6 := fact(6)
	fmt.Println("6! = ", fact6)

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fib(n-1) + fib(n-2)
		}
	}
	fmt.Println("fib(6) = ", fib(6)) // 0 1 1 2 3 5 8 13

	fmt.Println("---POINTERS---")
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // NB &i
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

	fmt.Println("---METHODS---")

	r := rect{width: 10, height: 5}
	fmt.Println("area of rect = ", r.area())
	fmt.Println("perim of rect = ", r.perim())
	rp := &r
	fmt.Println("area of rect = ", rp.area())
	fmt.Println("perim of rect = ", rp.perim())
}

func fact(n int) int {
	return iter(n, 1)
}

func iter(n int, accum int) int {
	if n == 0 {
		return accum
	} else {
		return iter(n-1, n*accum)
	}
}

func zeroval(ival int) { // call by value
	ival = 0 // this is a local var, it won't affect the passed ival
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. The *iptr code in the function body
// then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a
// dereferenced pointer changes the value at the referenced address.
func zeroptr(ptr *int) {
	*ptr = 0 // dereferencing the ptr and overriding the corresponding value
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int { // NB the syntax (r *rect) perim()
	return 2*r.width + 2*r.height
}
