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
	Health    float64
	Energy    float64
	Neighbors []int
}

func (world *World) CreateEntity(t int, pos Vec3, rot Vec3) {
	id := world.EntityCount
	world.EntityCount++
	entity := Entity{
		Pos:       pos,
		Rot:       rot,
		Id:        id,
		Type:      t,
		ChunkX:    -1,
		ChunkY:    -1,
		Health:    100,
		Energy:    50,
		Neighbors: make([]int, 0),
	}
	entity.Pos.X = math.Max(0, math.Min(entity.Pos.X, float64(world.Size)-0.01))
	entity.Pos.Y = math.Max(0, math.Min(entity.Pos.Y, float64(world.Size)-0.01))
	entity.ChunkX = int(entity.Pos.X)
	entity.ChunkY = int(entity.Pos.Y)
	world.Additions = append(world.Additions, &entity)
}
func (world *World) RemoveEntity(entity *Entity) {
	world.Removals = append(world.Removals, entity.Id)
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
	// entity logic
	entity.Neighbors = entity.Neighbors[:0]
	world.GetChunkNeighbors(entity, 2)
	switch entity.Type {
	case 0:
		{
			entity.UpdateHero(world)
			break
		}
	case 1:
		{
			entity.UpdatePlant(world)
			break
		}
	}
}

func (entity *Entity) UpdateChunkPosition(world *World) {
	world.Chunks[entity.ChunkX][entity.ChunkY].Append(entity.Id)
}
