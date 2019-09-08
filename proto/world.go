package proto

import (
	"encoding/json"
	"fmt"
)

type World struct {
	Time        float64
	Delta       float64
	Size        int
	EntityCount int
	Entities    map[int]*Entity
	Removals    []int
	Additions   []*Entity
	Chunks      [][]*Chunk
}

func CreateWorld(size int) *World {
	chunks := CreateChunks(size)
	var world = World{
		GetTime(),
		0,
		size, 0,
		make(map[int]*Entity),
		make([]int, 0),
		make([]*Entity, 0),
		chunks,
	}
	for i := 0; i < 100; i++ {
		f := float64(size)
		pos := GetRandomVector(f)
		rot := GetRandomVector(3.141)
		pos.Z = 0.0
		rot.X = 0.0
		rot.Y = 0.0
		world.CreateEntity(0, pos, rot)
	}
	for i := 0; i < 100; i++ {
		f := float64(size)
		pos := GetRandomVector(f)
		rot := GetRandomVector(3.141)
		pos.Z = 0.0
		rot.X = 0.0
		rot.Y = 0.0
		world.CreateEntity(1, pos, rot)
	}
	println("worldSize: " + fmt.Sprint(world.Size))
	return &world
}

func (w *World) Update() {
	t := GetTime()
	w.Delta = t - w.Time
	w.Time = t
	w.ClearChunks()
	// Update Additions
	for _, e := range w.Additions {
		w.Entities[e.Id] = e
	}
	w.Additions = w.Additions[:0]
	// Update Removals
	for _, e := range w.Removals {
		w.Entities[e].Id = -1
	}
	w.Removals = w.Removals[:0]

	// Update Entities
	for _, e := range w.Entities {
		if e.Id == -1 {
			continue
		}
		e.Update(w)
		e.UpdateChunkPosition(w)
	}
}
func (w *World) GetJSON() string {
	w.EntityCount = len(w.Entities)
	b, e := json.Marshal(w)
	if e == nil {
		return string(b)
	}
	return ""
}
