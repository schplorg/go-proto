package proto

import "math/rand"

func (entity *Entity) UpdatePlant(world *World) {
	if entity.Energy < 50 {
		d := world.Delta * 10.0
		r := rand.Float64()
		entity.Energy += d * r
	} else {
		entity.Energy = 0
		neighborPlants := 0
		for _, p := range entity.Neighbors {
			n := world.Entities[p]
			if n.Type == 1 {
				neighborPlants++
			}
		}
		if neighborPlants < 3 {
			pos := GetRandomVector(1)
			pos.Z = 0.0
			pos = Norm(pos)
			rot := GetZeroVector()
			world.CreateEntity(1, Add(entity.Pos, pos), rot)
			//println(world.EntityCount)
		}
	}
}
