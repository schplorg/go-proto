package proto

import (
	"strconv"
)

type Chunk struct {
	Entities []int
	Last     int
}

func CreateChunk() *Chunk {
	chunk := Chunk{Entities: make([]int, 0), Last: -1}
	return &chunk
}

func (c *Chunk) Clear() {
	c.Last = -1
	c.Entities = c.Entities[:0]
}
func (c *Chunk) Append(entity int) {
	c.Last++
	c.Entities = append(c.Entities, entity)
}

func (c *Chunk) Update() {
	for _, e := range c.Entities {
		println(strconv.Itoa(e))
	}
}
func CreateChunks(size int) [][]*Chunk {
	chunks := make([][]*Chunk, size)
	for x := 0; x < size; x++ {
		chunks[x] = make([]*Chunk, size)
		for y := 0; y < size; y++ {
			chunks[x][y] = CreateChunk()
		}
	}
	return chunks
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
