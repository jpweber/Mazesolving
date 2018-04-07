package maze

import (
	"image"
	"image/png"
	"os"
)

func (m *Maze) Output() {

	solvedMaze := image.NewRGBA(image.Rect(0, 0, 5, 5))
	for i, row := range m.Graph {
		for j, col := range row {
			solvedMaze.Set(j, i, pixelToRGBA(col))
		}

	}

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create("test.png")
	if err != nil {
		// Handle error
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, solvedMaze)

	// Don't forget to close files
	outputFile.Close()
}
