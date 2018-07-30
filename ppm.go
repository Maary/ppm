package ppm

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

//PpImage is an ppm image object
type PpImage struct {
	name          string
	mode          string
	wdith, height int
	maxPixel      uint16
	pixel         [][]Vector
	Colors
}

//Vector used for rgb color
type Vector struct {
	X, Y, Z float64
}

type Colors interface {
	CreateColors(interface{})
}

//CreateColors init the pixel indo
func (p *PpImage) CreateColors(colorsSet [][]Vector) *PpImage {

	for i := 0; i < p.height; i++ {
		for j := 0; j < p.wdith; j++ {
			p.pixel[i][j].X = colorsSet[i][j].X
			p.pixel[i][j].Y = colorsSet[i][j].Y
			p.pixel[i][j].Z = colorsSet[i][j].Z
		}
	}
	return p
}

const color = 255.9999999

var (
	nilSizeErr    = errors.New("nil image can not to draw")
	nilNameErr    = errors.New("nil image name can not to draw")
	nameFormatErr = errors.New("file name format shoule be .ppm")
)

//NewPPMImage init an ppm image
func NewPPMImage() *PpImage {
	return &PpImage{
		name:     "",
		mode:     "",
		wdith:    0,
		height:   0,
		maxPixel: 0,
	}
}

//SetName set the image name.
func (p *PpImage) SetName(name string) *PpImage {
	p.name = name
	return p
}

//SetMode set the ppm format, Mode default is "p6".
func (p *PpImage) SetMode(mode ...string) *PpImage {
	if len(mode) == 0 {
		p.mode = "p6"
	}
	p.mode = mode[0]
	return p
}

//SetWidthAndHeight set image's size.
func (p *PpImage) SetWidthAndHeight(wdith, height int) *PpImage {
	p.wdith = wdith
	p.height = height
	return p
}

//SetPixelNum set num of pixel, MaxPixel's max size is 256.
func (p *PpImage) SetPixelNum(maxPixel uint16) *PpImage {
	p.maxPixel = maxPixel
	return p
}

//TODO
func (p *PpImage) FillColor(colors [][]Vector) *PpImage {
	p.pixel = colors
	return p
}

//Draw the ppm info to file.
func (p *PpImage) Draw(colors interface{}) (err error) {
	if p.height == 0 || p.wdith == 0 {
		return nilSizeErr
	}
	if p.name == "" {
		return nilNameErr
	}
	if !strings.HasSuffix(p.name, "ppm") {
		return nameFormatErr
	}
	//TODO
	f, err := os.Create(p.name)
	defer f.Close()
	defer fmt.Println("OK => SAVED: ", p.name)
	if err != nil {
		return err
	}

	for _, colorW := range p.pixel {
		for _, colorH := range colorW {
			_, err = fmt.Fprintf(f, "%d %d %d\n", colorH.X, colorH.Y, colorH.Z)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
