package proto

import (
	"strconv"
)

type Chunk struct {
	Entities []int
}

func Create() *Chunk {
	var chunk Chunk
	return &chunk
}

func Update(c *Chunk) {
	for _, e := range c.Entities {
		println(strconv.Itoa(e))
	}
}
