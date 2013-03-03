package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
)

const (
	MAX_REFLECT_DEPTH = 7
	SPECULAR_SPREAD   = 5
)

var (
	AMBIENT RGB = RGB{255, 255, 255}.Scale(.15)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pixwidth := flag.Int("w", 500, "Rendered image pixel width")
	outfile := flag.String("o", "out.ppm", "Rendered image filename")
	flag.Parse()

	f, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	w := *pixwidth            // image pixel width
	h := (w * 4) / 5          // 5:4 aspect ratio
	height := .4              // viewscreen height
	width := (height * 5) / 4 // viewscreen width
	ph := height / float64(h) // pixel height
	pw := width / float64(w)  // pixel width

	// Render scene in memory
	memrend := make([][]RGB, h)

	// Write image header
	fmt.Fprintf(f, "P6\n%d %d\n255\n", w, h)

	// for each row
	for i := 0; i < h; i++ {
		// delegate this line to a goroutine
		go func(i int) {
			// y coord of row
			y := (height / 2) - (ph * (float64(i) + .5))
			line := make([]RGB, w)
			// for each column
			for j := 0; j < w; j++ {
				// x coord of column
				x := -(width / 2) + (pw * (float64(j) + .5))
				// z coord is on screen
				z := 0.0
				pt := Vector3{x, y, z}
				ur := pt.Sub(camera).Unit()
				if o, _ := trace(pt, ur, objects); o != nil {
					// pixel colored by object hit
					line[j] = shade(o, pt, ur, objects, lights, 0)
				}
			}
			memrend[i] = line
		}(i)
	}

	// Write rendered scene to file
	for _, line := range memrend {
		binary.Write(f, binary.LittleEndian, line)
	}
}

// trace returns the first object hit by the ray and the distance to it.
// pt - ray origin
// ur - ray unit vector
// objects - all objects in the scene
func trace(pt, ur Vector3, objects []Object) (o Object, t float64) {
	t = math.Inf(0)
	for _, object := range objects {
		if ti := object.Intersect(pt, ur); ti < t {
			t = ti
			o = object
		}
	}
	return
}

// shade recursively returns the shade of each pixel  
func shade(o Object, pt, ur Vector3, objects []Object, lights []*Light, depth uint) (color RGB) {
	if depth > MAX_REFLECT_DEPTH {
		return RGB{}
	}
	vr := ur.Reflect(pt, o.Normal(pt))
	// TODO: t never equals Inf. Is this because of the interface?
	if o, t := trace(pt, vr, objects); t == math.Inf(0) {
		// background color
		return RGB{}
	} else {
		mColor := shade(o, pt.Add(vr.Scale(t)), vr, objects, lights, depth+1)
		color = directShade(o, pt, ur, vr.Scale(-1), mColor)
	}
	for _, light := range lights {
		ul := pt.Sub(light.Point()).Unit()
		// Trace a ray from the light source to the object. If the first object hit is
		// the object, add the light to the pixel.
		if om, _ := trace(light.Point(), ul, objects); o == om {
			color.Add(directShade(o, pt, ur, light.Point(), light.Color()))
		}
	}
	return color
}

// directShade returns the sum of all light from a pixel
func directShade(o Object, pt, ur, ul Vector3, cl RGB) (radiance RGB) {
	view := ur.Scale(-1)
	n := o.Normal(pt)
	// reflected ray
	vr := ul.Reflect(pt, n)
	diffuse := cl.Mul(o.Color()).Scale(-1 * ul.Dot(n) * o.Diffuse())
	specular := cl.Scale(math.Pow(vr.Dot(view), SPECULAR_SPREAD) * o.Specular())
	return AMBIENT.Add(diffuse).Add(specular)
}
