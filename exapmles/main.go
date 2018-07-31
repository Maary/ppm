package main

import (
	"fmt"
	"testDir/PPM"
)

func main() {
	p1 := ppm.NewPPMImage("p1.ppm", "p6", 512, 512, 255)
	colors := make([][]ppm.Vector, 512)
	for i := range colors {
		colors[i] = make([]ppm.Vector, 512)
	}

	for i := 0; i < 512; i++ {
		for j := 0; j < 512; j++ {
			colors[i][j].X = 10
			colors[i][j].Y = 10
			colors[i][j].Z = 10
		}
	}
	fmt.Println("colors => ", colors)
	p1.CreateColors(colors)
	p1.Draw()
}
