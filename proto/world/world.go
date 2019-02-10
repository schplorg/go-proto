package world

import (
	"../entity"
	"fmt"
)

type World struct {
	Size     int
	Entities []entity.Entity
}

func Create(size int, heroCount int) *World {
	var entities = make([]entity.Entity, heroCount)
	for i := 0; i < len(entities); i++ {
		entities[i] = entity.Create()
	}
	var world = World{size, entities}
	println("worldSize: " + fmt.Sprint(world.Size))
	return &world
}

func Update(w *World) {
	for i := 0; i < len(w.Entities); i++ {
		e := w.Entities[i]
		println(e.Pos.String())
	}
}
