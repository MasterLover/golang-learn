package main

import (
	"gin-learning/app/system"
	"gin-learning/routers"
)

func main() {
	routers.Include(system.Routers)
	r := routers.Init()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
