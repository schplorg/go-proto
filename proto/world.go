package proto

import (
	"encoding/json"
	"fmt"
)

type World struct {
	Time     float64
	Size     int
	Entities []Entity
	Chunks   [][]*Chunk
}

func CreateWorld(size int, heroCount int) *World {
	entities := make([]Entity, heroCount)
	for i := 0; i < len(entities); i++ {
		f := float64(size)
		entities[i] = CreateEntity(0, GetRandomPosition(f))
	}
	chunks := CreateChunks(size)
	var world = World{GetTime(), size, entities, chunks}
	println("worldSize: " + fmt.Sprint(world.Size))
	return &world
}
func CreateChunks(size int) [][]*Chunk {
	println("CreateChunks")
	chunks := make([][]*Chunk, size)
	for x := 0; x < size; x++ {
		chunks[x] = make([]*Chunk, size)
		for y := 0; y < size; y++ {
			chunks[x][y] = CreateChunk()
		}
	}
	return chunks
}

func (w *World) Update() {
	w.ClearChunks()
	w.UpdateChunks()
	w.Time = GetTime()
	println("Update w.Entities")
	for i := 0; i < len(w.Entities); i++ {
		e := &w.Entities[i]
		e.Update(w)
	}
	println("Update w.Entities done")
}
func (w *World) UpdateChunks() {
	println("UpdateChunks")
	for i := 0; i < len(w.Entities); i++ {
		e := &w.Entities[i]
		w.Chunks[e.ChunkX][e.ChunkY].Append(i)
	}
}
func (w *World) ClearChunks() {
	println("ClearChunks")
	chunks := w.Chunks
	chunkLength := len(chunks)
	for x := 0; x < chunkLength; x++ {
		for y := 0; y < chunkLength; y++ {
			chunks[x][y].Clear()
		}
	}
}
func (w *World) GetJSON() string {
	b, e := json.Marshal(w)
	if e == nil {
		return string(b)
	}
	return ""
}
