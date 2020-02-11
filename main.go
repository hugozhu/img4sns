package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var wide bool
var dir string
var output string

func combine(file1 string, file2 string, fileOutput string) {
	imgFile1, err := os.Open(file1)
	imgFile2, err := os.Open(file2)
	if err != nil {
		fmt.Println(err)
	}
	img1, err := png.Decode(imgFile1)
	img2, err := png.Decode(imgFile2)
	if err != nil {
		fmt.Println(err)
	}

	var rgba *image.RGBA
	var sp2 image.Point

	//starting position of the second image (bottom left)
	if wide {
		sp2 = image.Point{img1.Bounds().Dx(), 0}
	} else {
		sp2 = image.Point{0, img1.Bounds().Dy()}
	}
	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	r := image.Rectangle{image.Point{0, 0}, r2.Max}
	rgba = image.NewRGBA(r)
	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)

	out, err := os.Create(fileOutput)
	if err != nil {
		fmt.Println(err)
	}
	png.Encode(out, rgba)
}

func init() {
	flag.BoolVar(&wide, "wide", false, "Wide picture output")
	flag.StringVar(&dir, "dir", "tmp", "Folder to read")
	flag.StringVar(&output, "output", "output", "Output file")
	flag.Parse()
}

func main() {
	files, err := filepath.Glob(dir + "/" + "*.[pP][nN][gG]")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var file1, file2, file3 string
	file1 = files[0]
	file3 = "output.png"
	for i, file := range files {
		if i > 0 {
			file2 = file
			log.Println(file1, file2, file3)
			combine(file1, file2, file3)
			file1 = file3
		}
	}
}
