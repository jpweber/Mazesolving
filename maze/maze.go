package maze

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
)

var (
	kWhite = Pixel{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	kBlack = Pixel{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}
	kRed = Pixel{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	kGreen = Pixel{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}
	kPathColor = Pixel{
		R: 92,
		G: 221,
		B: 240,
		A: 255,
	}
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

type Maze struct {
	Graph [][]Pixel
	Nodes []Node
	Start map[int64]int64
	Goal  map[int64]int64
	// [column][row,row,row]
	NodePointsCol map[int64][]Node
	// [row][col,col,col]
	NodePointsRow map[int64][]Node
	Neighbors     map[Node][]Node
}

type Node struct {
	Row int64
	Col int64
}

func imageDecoder(file io.Reader) Maze {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln("Could not decode image:", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	var maze Maze
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	maze.Graph = pixels
	return maze
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
func pixelToRGBA(p Pixel) color.RGBA {
	rgba := color.RGBA{
		R: uint8(p.R),
		G: uint8(p.G),
		B: uint8(p.B),
		A: uint8(p.A),
	}

	return rgba
}

// Read  file and pass it on to the maze decoder
func Read(fileName string) Maze {
	// open file and read contents
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	return imageDecoder(file)

}

func (m *Maze) findStart() {
	// init start map
	m.Start = make(map[int64]int64)
loop:
	for i, row := range m.Graph {
		for j, col := range row {
			if col == kRed {
				log.Println("Found Start:", i, ",", j)
				m.Start[int64(i)] = int64(j)
				break loop
			}
		}
	}
}

func (m *Maze) findGoal() {
	// init start map
	m.Goal = make(map[int64]int64)
loop:
	for i, row := range m.Graph {
		for j, col := range row {
			if col == kGreen {
				log.Println("Found Goal:", i, ",", j)
				m.Goal[int64(i)] = int64(j)
				break loop
			}
		}
	}
}
