package proto

func (entity *Entity) UpdateHero(world *World) {
	// sin := math.Sin(entity.Rot.Z * world.Time / 4.0)
	// cos := math.Cos(entity.Rot.Z * world.Time / 4.0)
	// m := CreateVec3(sin, sin*cos*cos, 0)
	// m = Mul(m, 0.01)
	// entity.Move(m, world)
	if entity.Energy < 100 {
		for _, p := range entity.Neighbors {
			n := world.Entities[p]
			if n.Type == 1 {
				if n.ChunkX == entity.ChunkX && n.ChunkY == entity.ChunkY {
					//world.RemoveEntity(n)
				} else {
					toTarget := Add(n.Pos, Mul(entity.Pos, -1.0))
					toTarget = Mul(toTarget, 0.01)
					entity.Move(toTarget, world)
				}
			}
		}
	}
}
