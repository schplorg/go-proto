package proto

import (
	"fmt"
	"time"
)

func Start() {
	fmt.Println("creating world")
	w := CreateWorld(10, 100)
	StartServer(w)
	ticker := time.NewTicker(50 * time.Millisecond)
	go UpdateProto(ticker, w)
	time.Sleep(600000 * time.Millisecond)
	ticker.Stop()
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
