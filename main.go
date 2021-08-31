package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"time"
)

const RGBTHRESHOLD = 32767
const FRAMECOUNT = 6572
const FRAMEDELAY = 30
const DRAWINGCHARACTER = "."

type Frame struct {
	Source image.Image
}

func (f *Frame) Clear() {
	fmt.Print("\033[12A")
	time.Sleep(FRAMEDELAY * time.Millisecond)
}

func (f *Frame) Display() {
	bounds := f.Source.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			imageColor := f.Source.At(x, y)
			r, g, b, _ := imageColor.RGBA()
			character := DRAWINGCHARACTER
			if r < RGBTHRESHOLD || g < RGBTHRESHOLD || b < RGBTHRESHOLD {
				character = " "
			}
			fmt.Printf("%s", character)
		}
		fmt.Printf("%c", 10)
	}
}

func (f *Frame) Load(frameNumber int) {
	infile, err := os.Open(fmt.Sprintf("frames/badapple%04d.png", frameNumber))
	if err != nil {
		log.Fatalln(infile, err)
	}
	defer infile.Close()

	f.Source, _, err = image.Decode(infile)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	for i := 1; i <= FRAMECOUNT; i++ {
		f := Frame{}
		f.Clear()
		f.Load(i)
		f.Display()
	}
}
