package world

import (
	"../entity"
	"fmt"
)

type World struct {
	Size     int
	Entities [1000]entity.Entity
}

func Create() *World {
	var entities [1000]entity.Entity
	for i := 0; i < len(entities); i++ {
		entities[i] = entity.Create()
	}
	var world = World{10, entities}
	println("worldSize: " + fmt.Sprint(world.Size))
	return &world
}

func Update(w *World) {
	// var wo = *w
	// for _, e := range wo.Entities {
	// 	println(e.Pos.String())
	// }

	for _, e := range w.Entities {
		println(e.Pos.String())
	}
}
