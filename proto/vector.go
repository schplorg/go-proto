package proto

import (
	"fmt"
	"math"
	"math/rand"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func CreateVec3(x float64, y float64, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func Add(a Vec3, b Vec3) Vec3 {
	return CreateVec3(a.X+b.X, a.Y+b.Y, a.Z+b.Z)
}

func Mul(a Vec3, s float64) Vec3 {
	return CreateVec3(a.X*s, a.Y*s, a.Z*s)
}

func Div(a Vec3, s float64) Vec3 {
	return CreateVec3(a.X/s, a.Y/s, a.Z/s)
}

func Mag(a Vec3) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func Norm(a Vec3) Vec3 {
	return Div(a, Mag(a))
}

func (b Vec3) String() string {
	return fmt.Sprintf("%f %f %f", b.X, b.Y, b.Z)
}
func GetRandomPosition(scale float64) Vec3 {
	pos := Vec3{X: rand.Float64() * scale, Y: rand.Float64() * scale, Z: rand.Float64() * scale}
	return pos
}
