package proto

import (
	"math"
)

type Entity struct {
	Pos       Vec3
	Rot       Vec3
	Id        int
	Type      int
	ChunkX    int
	ChunkY    int
	Neighbors []int
}

func CreateEntity(id int, t int, pos Vec3, rot Vec3) Entity {
	ent := Entity{Pos: pos, Rot: rot, Id: id, Type: t, ChunkX: -1, ChunkY: -1, Neighbors: make([]int, 0)}
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
	world.Chunks[entity.ChunkX][entity.ChunkY].Append(entity.Id)
}

func (entity *Entity) Update(world *World) {
	sin := math.Sin(entity.Rot.Z * world.Time / 4.0)
	cos := math.Cos(entity.Rot.Z * world.Time / 4.0)
	m := CreateVec3(sin, sin*cos*cos, 0)
	m = Mul(m, 0.01)
	entity.Move(m, world)
	entity.Neighbors = entity.Neighbors[:0]
	chunk := world.Chunks[entity.ChunkX][entity.ChunkY]
	for _, neighbor := range chunk.Entities {
		if neighbor != entity.Id {
			entity.Neighbors = append(entity.Neighbors, neighbor)
		}
	}
}
