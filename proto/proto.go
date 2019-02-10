package proto

import (
	//"./server"
	"./world"
	"fmt"
)

func Start() {
	//server.Start()
	fmt.Println("creating world")
	var w = world.Create(10, 100)
	for i := 0; i < 3; i++ {
		world.Update(w)
		fmt.Scanln()
	}
}
