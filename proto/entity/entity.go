package entity

import (
	"../vector"
	"math/rand"
)

type Entity struct {
	Pos vector.Vec3
}

func Create() Entity {
	pos := vector.Vec3{X: rand.Float32(), Y: rand.Float32(), Z: rand.Float32()}
	ent := Entity{Pos: pos}
	return ent
}
