package main

import (
	"math"
)

// Vector3 is a 3 axis vector
type Vector3 struct {
	x, y, z float64
}

// Add returns the sum of the vector and given vector
func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{a.x + b.x, a.y + b.y, a.z + b.z}
}

// Cross returns the cross product of the vector and given vector
// The magnitude of the cross product is given by
// ||a×b|| = ||a||||b|| |sin θ|,
// where θ is the small angle between vectors a and b. Thus, if a and b are unit
// vectors, the magnitude of the cross product is the magnitude of sin θ.
// Note, that the cross product of two parallel vectors will be the zero vector 0.
// This is consistent with the geometric notion that the cross product produces a
// vector orthogonal to the original two vectors. If the original vectors are parallel,
// then there is no unique direction perpendicular to both vectors (i.e. there are
// infinitely many orthogonal vectors, all parallel to any plane perpendicular to
// either vector).
func (a Vector3) Cross(b Vector3) Vector3 {
	return Vector3{a.y*b.z - a.z*b.y, a.z*b.x - a.x*b.z, a.x*b.y - a.y*b.x}
}

 // Dot returns the dot product of the vector and given vector
// a*b=||A||||b||cos(theta) where theta is the smallest angle
// between the vectors.
// If a and b are unit vectors, ||a||||b|| = 1 and cos(theta) = a*b
// To find angle between vectors, calculate unit vectors then
// cos(theta) = ua*ub
// Dot product of orthogonal vectors is zero
// Dot product is negative if the smallest angle between them
// exceeds 90deg

func (a Vector3) Dot(b Vector3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

// Magnitude returns the vector magnitude
func (a Vector3) Magnitude() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}

// Reflect returns the vector reflection off the surface
// defined by point and normal
func (a Vector3) Reflect(point, normal Vector3) Vector3 {
	// v' = v - 2(n*v)n
	return a.Sub(normal.Scale(normal.Dot(a) * 2))
}

// Scale returns the vector scaled by the scalar c
func (a Vector3) Scale(c float64) Vector3 {
	return Vector3{a.x * c, a.y * c, a.z * c}
}

// Sub returns the difference between the vector and given vector
func (a Vector3) Sub(b Vector3) Vector3 {
	return Vector3{a.x - b.x, a.y - b.y, a.z - b.z}
}

// Unit returns the unit vector
func (a Vector3) Unit() Vector3 {
	// multiplication is faster than division
	invmag := 1 / a.Magnitude()
	return Vector3{a.x * invmag, a.y * invmag, a.z * invmag}
}
