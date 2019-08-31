package proto

import (
	"encoding/json"
	"fmt"
)

type World struct {
	Time        float64
	Size        int
	EntityCount int
	Entities    []Entity
	Chunks      [][]*Chunk
}

func CreateWorld(size int, heroCount int) *World {
	entities := make([]Entity, heroCount)
	for i := 0; i < len(entities); i++ {
		f := float64(size)
		pos := GetRandomVector(f)
		rot := GetRandomVector(3.141)
		pos.Z = 0.0
		rot.X = 0.0
		rot.Y = 0.0
		entities[i] = CreateEntity(i, 0, pos, rot)
	}
	chunks := CreateChunks(size)
	var world = World{GetTime(), size, len(entities), entities, chunks}
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
	w.Time = GetTime()
	for i := 0; i < len(w.Entities); i++ {
		e := &w.Entities[i]
		e.Update(w)
	}
}
func (w *World) ClearChunks() {
	chunks := w.Chunks
	chunkLength := len(chunks)
	for x := 0; x < chunkLength; x++ {
		for y := 0; y < chunkLength; y++ {
			chunks[x][y].Clear()
		}
	}
}
func (w *World) GetChunkNeighbors(entity *Entity, level int) {
	lenChunks := len(w.Chunks)
	f := 1 + ((level - 1) * 2)
	startX := entity.ChunkX - (level - 1)
	startY := entity.ChunkY - (level - 1)
	for x := 0; x < f; x++ {
		for y := 0; y < f; y++ {
			k := startX + x
			l := startY + y
			if k < 0 || k >= lenChunks || l < 0 || l >= lenChunks {
				continue
			}
			chunk := w.Chunks[k][l]
			for _, neighbor := range chunk.Entities {
				if neighbor != entity.Id {
					entity.Neighbors = append(entity.Neighbors, neighbor)
				}
			}
		}
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
