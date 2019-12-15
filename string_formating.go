package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 16)
	fmt.Printf("%x\n", 15)

	fmt.Printf("%f\n", 15.16)

	fmt.Printf("%e\n", 1234567800000.0)

	fmt.Printf("%E\n", 1234567800000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.211111111111, 11111111111113333333333.45)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 1.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 1.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "suprise!, happy birthday!")
	fmt.Printf("|%--16s|%--16s|\n", "foo", "suprise!, happy birthday!")

	s := fmt.Sprintf("|%--16s|%--16s|\n", "foo", "suprise!, happy birthday!")
	println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
