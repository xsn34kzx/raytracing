package main

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

// Methods

func (lhs *Vec3) Length() float64 {
	return math.Sqrt(lhs.LengthSquared())
}

func (lhs *Vec3) LengthSquared() float64 {
	return lhs.X*lhs.X + lhs.Y*lhs.Y + lhs.Z*lhs.Z
}

func (lhs *Vec3) Dot(rhs Vec3) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Y
} 

func (lhs *Vec3) Cross(rhs Vec3) Vec3 {
	return Vec3{
		lhs.Y*rhs.Z - lhs.Z*lhs.Y,
		lhs.Z*lhs.X - lhs.X*rhs.Z,
		lhs.X*rhs.Y - lhs.Y*lhs.X,
	}
}

func (lhs *Vec3) Add(rhs Vec3) {
	lhs.X += rhs.X
	lhs.Y += rhs.Y
	lhs.Z += rhs.Y
}

func (lhs *Vec3) Subtract(rhs Vec3) {
	lhs.X -= rhs.X
	lhs.Y -= rhs.Y
	lhs.Z -= rhs.Y
}

func (v *Vec3) Scale(c float64) {
	v.X *= c
	v.Y *= c
	v.Z *= c
}

func (v *Vec3) String() string {
	return fmt.Sprintf("%v %v %v", v.X, v.Y, v.Z)
}

// Functions

func Add(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3{lhs.X + rhs.X, lhs.Y + rhs.Y, lhs.Z + rhs.Z}
}

func Subtract(lhs Vec3, rhs Vec3) Vec3 {
	return Vec3{lhs.X - rhs.X, lhs.Y - rhs.Y, lhs.Z - rhs.Z}
}

func Scale(v Vec3, c float64) Vec3 {
	return Vec3{v.X*c, v.Y*c, v.Z*c}
}

func UnitVector(v Vec3) Vec3 {
	return Scale(v, 1 / v.Length())
}
