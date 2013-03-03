package main

import (
	"math"
)

const (
	EPSILON = .000000000000001
)
// TODO: Load scene from file
var (

	camera Vector3  = Vector3{0, 0, .5}
	lights []*Light = []*Light{
		&Light{color: RGB{255, 255, 255}, point: Vector3{0, 10, 0}},
		&Light{color: RGB{255, 255, 255}, point: Vector3{-5, 7, 3}},
	}
	objects []Object = []Object{
		&Plane{color: RGB{255, 0, 255}, diffuse: .7, reflect: .5, specular: .3, normal: Vector3{0, 1, 0}, point: Vector3{0, -.2, 0}},
		&Sphere{color: RGB{255, 0, 0}, diffuse: .7, reflect: .5, specular: .3, point: Vector3{-.3, .2, -.6}, radius: .2},      // red
		&Sphere{color: RGB{255, 154, 0}, diffuse: .7, reflect: .5, specular: .3, point: Vector3{.15, -.2, -.6}, radius: .15},  // orange
		&Sphere{color: RGB{255, 255, 0}, diffuse: .7, reflect: .5, specular: .3, point: Vector3{.1, .175, -.15}, radius: .05}, // yellow
		&Sphere{color: RGB{0, 255, 0}, diffuse: .7, reflect: .5, specular: .3, point: Vector3{.0, .13, -.3}, radius: .025},    // green
		&Sphere{color: RGB{0, 0, 255}, diffuse: .7, reflect: .5, specular: .3, point: Vector3{.3, -.2, -.2}, radius: .125},    // blue
		&Sphere{color: RGB{75, 0, 130}, diffuse: .7, reflect: 0.0, specular: .3, point: Vector3{-.2, .0, -.4}, radius: .06},   // indigo
	}
)

// TODO coordinate type

type Camera struct {
	direction Vector3
	orientation Vector3
	position Vector3
}

type Light struct {
	color RGB
	point Vector3
}

func (l *Light) Color() RGB {
	return l.color
}

func (l *Light) Point() Vector3 {
	return l.point
}

type material interface {
	Color() RGB
	Diffuse() float64
	Reflect() float64
	Specular() float64
}

type Object interface {
	material
	Intersect(pt, ur Vector3) float64
	Normal(pt Vector3) Vector3
	Point() Vector3
}

type Plane struct {
	color    RGB
	diffuse  float64
	normal   Vector3
	point    Vector3
	reflect  float64
	specular float64
}

func (p *Plane) Color() RGB {
	return p.color
}

func (p *Plane) Diffuse() float64 {
	return p.diffuse
}

func (p *Plane) Intersect(pt, ur Vector3) float64 {
	if t := p.normal.Dot(p.point.Sub(pt)) / p.normal.Dot(ur); t >= 0 {
		return t - EPSILON
	}
	return math.Inf(0)
}

func (p *Plane) Normal(pt Vector3) Vector3 {
	return p.normal
}

func (p *Plane) Point() Vector3 {
	return p.point
}

func (p *Plane) Reflect() float64 {
	return p.reflect
}

func (p *Plane) Specular() float64 {
	return p.specular
}

type Sphere struct {
	color    RGB
	diffuse  float64
	point    Vector3
	radius   float64
	reflect  float64
	specular float64
}

func (s *Sphere) Color() RGB {
	return s.color
}

func (s *Sphere) Diffuse() float64 {
	return s.diffuse
}

// Intersect returns the magnitude of the nearest intersect or Inf
// if it doesn't exist
func (s *Sphere) Intersect(pt, ur Vector3) float64 {
	// if the vector is normalized, a is always 1
	b := ur.Scale(2).Dot(pt.Sub(s.point))
	c := pt.Sub(s.point).Dot(pt.Sub(s.point)) - s.radius*s.radius
	disc := b*b - 4*c
	if disc < 0 {
		return math.Inf(0)
	}
	t0 := (-b - math.Sqrt(disc)) / 2
	if t0 > 0 {
		return t0 - EPSILON
	}
	t1 := (-b + math.Sqrt(disc)) / 2
	return t1 - EPSILON
}

func (s *Sphere) Normal(pt Vector3) Vector3 {
	return pt.Sub(s.point).Unit()
}

func (s *Sphere) Point() Vector3 {
	return s.point
}

func (s *Sphere) Reflect() float64 {
	return s.reflect
}

func (s *Sphere) Specular() float64 {
	return s.specular
}

type RGB struct {
	r, g, b uint8
}

func (a RGB) Add(b RGB) RGB {
	return RGB{a.r + b.r, a.g + b.g, a.b + b.b}
}

func (a RGB) Mul(b RGB) RGB {
	return RGB{a.r * b.r, a.g * b.g, a.b * b.b}
}

func (a RGB) Scale(c float64) RGB {
	return RGB{uint8(float64(a.r) * c), uint8(float64(a.g) * c), uint8(float64(a.b) * c)}
}
