package main

import (
	"fmt"
)

func zeroptr(iptr *int) {
	*iptr = 0
}

func zeroval(ival int) {
	ival = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)

	fmt.Println("zeroptr:", i)

	zeroptr(&i)

	fmt.Println("zeroptr:", i)

	fmt.Println("zeroptr:", &i)

}
