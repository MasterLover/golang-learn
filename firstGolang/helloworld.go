package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func updatePerson(p *person) {
	p.name = "json"
	p.age = 20
}

func main() {

	fmt.Println("hello world")
	//time.Sleep(1000)

	p := person{"kaka", 18}
	//updatePerson(p)
	updatePerson(&p)
	fmt.Println("after update")
	fmt.Println("name:", p.name, "age:", p.age)
}

// 将updatePerson函数的参数改成指针参数时，
