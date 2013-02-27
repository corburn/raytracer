package main

import (
	"flag"
	"log"
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
}

func trace() {
}
