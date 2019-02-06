package vector

import (
	"fmt"
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func Create(x float32, y float32, z float32) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func (b Vec3) String() string {
	return fmt.Sprintf("%f %f %f", b.X, b.Y, b.Z)
}
