package proto

import (
	"fmt"
	"time"
)

func Start() {
	fmt.Println("creating world")
	w := CreateWorld(10, 100)
	StartServer(w)

	// var waitgroup sync.WaitGroup
	// waitgroup.Add(1)
	// waitgroup.Wait()

	// for i := 0; i < 3; i++ {
	// 	println("proto.update")
	// 	world.Update(w)
	// 	fmt.Scanln()
	// }

	ticker := time.NewTicker(500 * time.Millisecond)
	go UpdateProto(ticker, w)
	time.Sleep(60000 * time.Millisecond)
	ticker.Stop()
	fmt.Scanln()
}

func UpdateProto(ticker *time.Ticker, world *World) {
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		world.Update()
	}
}

func GetTime() float64 {
	return float64(time.Now().UnixNano()/1000000) / 1000.0
}
