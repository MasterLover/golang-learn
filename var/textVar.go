package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println(a)
	fmt.Printf("type of %T\n", a)

	var b string = "222222"
	fmt.Println(b)
	fmt.Printf("type of %T\n", b)

	var c = 100
	fmt.Println(c)
	fmt.Printf("type of %T\n", c)

	d := 100
	fmt.Println(d)
	fmt.Printf("type of d %T\n", d)

	g := 0.151213215454
	fmt.Printf("type of g  %T\n", g)

	var xx, yy int = 100, 200
	fmt.Println(xx, yy)

	var (
		vv int  = 3005
		jj bool = true
	)
	fmt.Println(vv, jj)
}
