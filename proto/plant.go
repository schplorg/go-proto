package proto

import "math/rand"

func (entity *Entity) UpdatePlant(world *World) {
	if entity.Energy < 100 {
		entity.Energy += (int)((world.Delta * 10) * rand.Float64())
	} else {
		entity.Energy = 0
		pos := GetRandomVector(1)
		pos = Norm(pos)
		rot := GetZeroVector()
		world.CreateEntity(1, Add(entity.Pos, pos), rot)
	}
}
