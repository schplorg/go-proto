package proto

import (
	//"./server"
	"./world"
	"fmt"
)

func Start() {
	//server.Start()
	fmt.Println("creating world")
	var w = world.Create()
	for i := 0; i < 3; i++ {
		world.Update(w)
		fmt.Scanln()
	}
}
