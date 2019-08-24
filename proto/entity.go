package proto

import (
	"math"
	"strconv"
)

type Entity struct {
	Pos       Vec3
	Type      int
	ChunkX    int
	ChunkY    int
	Neighbors []int
}

func CreateEntity(t int, pos Vec3) Entity {
	ent := Entity{Pos: pos, Type: t, ChunkX: -1, ChunkY: -1, Neighbors: make([]int, 0)}
	ent.ChunkX = int(ent.Pos.X)
	ent.ChunkY = int(ent.Pos.Y)
	return ent
}

func (entity *Entity) Move(dir Vec3, world *World) {
	d := Norm(dir)
	m := math.Min(1, Mag(dir))
	entity.Pos = Add(entity.Pos, Mul(d, m))
	entity.Pos.X = math.Max(0, math.Min(entity.Pos.X, float64(world.Size)-0.01))
	entity.Pos.Y = math.Max(0, math.Min(entity.Pos.Y, float64(world.Size)-0.01))
	entity.ChunkX = int(entity.Pos.X)
	entity.ChunkY = int(entity.Pos.Y)
}

func (entity *Entity) Update(world *World) {
	println("Entity Update")
	sin := math.Sin(world.Time / 4.0)
	cos := math.Cos(world.Time / 4.0)
	entity.Move(CreateVec3(sin, 0, cos), world)
	entity.Neighbors = entity.Neighbors[:0]
	println("world.Chunks " + strconv.Itoa(entity.ChunkX) + " " + strconv.Itoa(entity.ChunkY))
	chunk := world.Chunks[entity.ChunkX][entity.ChunkY]
	for neighbor, _ := range chunk.Entities {
		entity.Neighbors = append(entity.Neighbors, neighbor)
	}
}
