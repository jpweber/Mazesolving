package maze

import (
	"image"
	"image/png"
	"log"
	"os"
)

func (m *Maze) Output() {

	solvedMaze := image.NewRGBA(m.Bounds)
	for i, row := range m.Graph {
		for j, col := range row {
			solvedMaze.Set(j, i, pixelToRGBA(col))
		}

	}

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test.png")
	if err != nil {
		log.Printf("error creating png file: %s", err)
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, solvedMaze)

	// Don't forget to close files
	outputFile.Close()
}
