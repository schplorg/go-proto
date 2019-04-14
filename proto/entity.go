package proto

import (
	"math"
	"math/rand"
)

type Entity struct {
	Pos  Vec3
	Type int
}

func CreateEntity(t int) Entity {
	pos := Vec3{X: rand.Float64(), Y: rand.Float64(), Z: rand.Float64()}
	ent := Entity{Pos: pos, Type: t}
	return ent
}

func (entity *Entity) Move(dir Vec3) {
	d := Norm(dir)
	m := math.Min(1, Mag(dir))
	entity.Pos = Add(entity.Pos, Mul(d, m))
}

func (entity *Entity) Update(world *World) {
	sin := math.Sin(world.Time)
	cos := math.Cos(world.Time)
	entity.Move(CreateVec3(sin, 0, cos))
}
