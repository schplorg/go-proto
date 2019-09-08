package proto

import "math"

func (entity *Entity) UpdateHero(world *World) {
	sin := math.Sin(entity.Rot.Z * world.Time / 4.0)
	cos := math.Cos(entity.Rot.Z * world.Time / 4.0)
	m := CreateVec3(sin, sin*cos*cos, 0)
	m = Mul(m, 0.01)
	entity.Move(m, world)
}
