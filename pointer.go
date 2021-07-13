package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	fmt.Println("pointer", &i)

	// Doesn't dereference i pointer, hence the value doesn't change
	zeroval(i)
	fmt.Println("zeroval:", i)
	fmt.Println("pointer", &i)

	// Dereference i pointer with & syntax, hence the value change to 0
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
