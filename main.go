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

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	pixwidth := flag.Int("w", 500, "Rendered image pixel width")
	_ = pixwidth
	outfile := flag.String("o", "out.ppm", "Rendered image filename")

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
				point := Vector3{x, y, z}
				dir := point.Sub(camera).Unit()
				if o, _ := trace(point, dir, objects); o != nil {
					// pixel colored by object hit
					line[j] = shade(o, point, dir, objects, lights, 0)
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

func trace(point, dir Vector3, objects []Object) (o Object, t float64) {
	t = math.Inf(0)
	for _, object := range objects {
		if ti := object.Intersect(point, dir); ti < t {
			t = ti
			o = object
		}
	}
	return
}

func shade(o Object, point, dir Vector3, objects []Object, lights []*Light, depth uint) (color RGB) {
	return o.Color()
}
