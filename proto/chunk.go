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
