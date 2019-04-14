package proto

import (
	"encoding/json"
	"fmt"
)

type World struct {
	Time     float64
	Size     int
	Entities []Entity
}

func CreateWorld(size int, heroCount int) *World {
	entities := make([]Entity, heroCount)
	for i := 0; i < len(entities); i++ {
		entities[i] = CreateEntity(0)
	}
	var world = World{GetTime(), size, entities}
	println("worldSize: " + fmt.Sprint(world.Size))
	return &world
}

func (w *World) Update() {
	w.Time = GetTime()
	for i := 0; i < len(w.Entities); i++ {
		e := &w.Entities[i]
		e.Update(w)
		//println(e.Pos.String())
	}
}

func (w *World) GetJSON() string {
	b, e := json.Marshal(w)
	//println("b = " + string(b))
	//println("e = " + fmt.Sprintf("%v", e))
	if e == nil {
		return string(b)
	} else {
		return ""
	}
}
