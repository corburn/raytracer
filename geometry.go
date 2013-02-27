package main

import (
	"math"
)

var (
	camera Vector3  = Vector3{0, 0, .5}
	lights []*Light = []*Light{
		&Light{color: RGB{255, 255, 255}, point: Vector3{0, 10, 0}},
		&Light{color: RGB{255, 255, 255}, point: Vector3{-5, 7, 3}},
	}
	objects []Object = []Object{
		&Plane{color: RGB{128, 128, 128}, normal: Vector3{0, 1, 0}, point: Vector3{0, -.2, 0}},
		&Sphere{color: RGB{255, 0, 0}, point: Vector3{-.3, .2, -.6}, radius: .2},      // red
		&Sphere{color: RGB{255, 154, 0}, point: Vector3{.15, -.2, -.6}, radius: .15},  // orange
		&Sphere{color: RGB{255, 255, 0}, point: Vector3{.1, .175, -.15}, radius: .05}, // yellow
		&Sphere{color: RGB{0, 255, 0}, point: Vector3{.0, .13, -.3}, radius: .025},    // green
		&Sphere{color: RGB{0, 0, 255}, point: Vector3{.3, -.2, -.2}, radius: .125},    // blue
		&Sphere{color: RGB{75, 0, 130}, point: Vector3{-.2, .0, -.4}, radius: .06},    // indigo
	}
)

type Light struct {
	color RGB
	point Vector3
}

type Object interface {
	Color() RGB
	Intersect(pt, ur Vector3) float64
	Normal(pt Vector3) Vector3
	Point() Vector3
}

type Plane struct {
	color  RGB
	normal Vector3
	point  Vector3
}

func (p *Plane) Color() RGB {
	return p.color
}

func (p *Plane) Intersect(pt, ur Vector3) float64 {
	if t := p.normal.Dot(p.point.Sub(pt)) / p.normal.Dot(ur); t >= 0 {
		return t
	}
	return math.Inf(0)
}

func (p *Plane) Normal(pt Vector3) Vector3 {
	return p.normal
}

func (p *Plane) Point() Vector3 {
	return p.point
}

type Sphere struct {
	color  RGB
	point  Vector3
	radius float64
}

func (s *Sphere) Color() RGB {
	return s.color
}

// Intersect returns the magnitude of the nearest intersect or Inf
// if it doesn't exist
func (s *Sphere) Intersect(pt, ur Vector3) float64 {
	a := ur.Dot(ur)
	b := ur.Scale(2).Dot(pt.Sub(s.point))
	c := pt.Sub(s.point).Dot(pt.Sub(s.point)) - s.radius*s.radius
	disc := b*b - 4*a*c
	if disc < 0 {
		return math.Inf(0)
	}
	t0 := (-b - math.Sqrt(disc)) / 2
	if t0 > 0 {
		return t0
	}
	t1 := (-b + math.Sqrt(disc)) / 2
	return t1
}

func (s *Sphere) Normal(pt Vector3) Vector3 {
	return pt.Sub(s.point).Unit()
}

func (s *Sphere) Point() Vector3 {
	return s.point
}

type RGB struct {
	r, g, b uint8
}
