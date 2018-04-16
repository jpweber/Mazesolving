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
	kNodeColor = Pixel{
		R: 92,
		G: 221,
		B: 240,
		A: 255,
	}
	kPathColor = Pixel{
		R: 0,
		G: 0,
		B: 255,
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
	Start Node
	Goal  Node
	// [column][row,row,row]
	NodePointsCol map[int64][]Node
	// [row][col,col,col]
	NodePointsRow map[int64][]Node
	Neighbors     map[Node][]Node
	Bounds        image.Rectangle
	Path          []Node
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
	maze.Bounds = bounds

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

func (m *Maze) DrawPath() {

	for i, n := range m.Path {
		m.Graph[n.Row][n.Col] = kPathColor
		if len(m.Path) != i+1 {
			m.connectPoints(m.Path[i], m.Path[i+1])
		}
	}

}

func (m *Maze) connectPoints(nodeA, nodeB Node) {

	var colDiff int64
	var rowDiff int64
	if nodeA.Row == nodeB.Row {
		colDiff = int64(nodeA.Col-nodeB.Col) * -1
	} else {
		rowDiff = int64(nodeA.Row-nodeB.Row) * -1
	}
	if colDiff == 0 {
		// fmt.Println("Row steps =", rowDiff)
		if rowDiff < 0 {
			for i := rowDiff; i < 0; i++ {
				m.Graph[nodeA.Row+i][nodeA.Col] = kPathColor
			}
		} else {
			for i := rowDiff; i > 0; i-- {
				m.Graph[nodeA.Row+1][nodeA.Col] = kPathColor
			}
		}
	} else {
		// fmt.Println("Col steps =", colDiff)
		if colDiff < 0 {
			for i := colDiff; i < 0; i++ {
				m.Graph[nodeA.Row][nodeA.Col+i] = kPathColor
			}
		} else {
			for i := colDiff; i > 0; i-- {
				m.Graph[nodeA.Row][nodeA.Col+1] = kPathColor
			}
		}
	}

}
func (m *Maze) DrawNodes() {
	for _, n := range m.Nodes {
		m.Graph[n.Row][n.Col] = kNodeColor
	}
}
