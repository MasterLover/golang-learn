package main

import (
	"fmt"
)

func main() {
	const length int = 110
	fmt.Println(length)
	//length = 200

	var a = 21
	var b = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	println()
	s, s2 := hello("hello", "world")
	fmt.Println(s, s2)
}
func hello(a string, b string) (string, string) {
	c := a + "," + b
	return c, c + "\t as"
}
