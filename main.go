package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jpweber/mazes/maze"
	"github.com/jpweber/mazes/solver"
)

func main() {

	mazeFile := flag.String("f", "", "Path to maze file")

	flag.Parse()

	if *mazeFile == "" {
		fmt.Println("Path to file with maze is required")
		os.Exit(1)
	}

	// parse our txt maze to maze struct
	mainMaze := maze.Read(*mazeFile)

	// Plot our nodes (decision points)
	start := time.Now()
	mainMaze.NodeFinder()
	duration := time.Since(start)
	fmt.Println("Time to plot nodes on maze", duration)

	start = time.Now()
	mainMaze.Path = solver.DumbAlg(mainMaze)
	duration = time.Since(start)
	fmt.Println("Time to solve on maze", duration)
	// fmt.Println(mainMaze.Path)
	// mainMaze.DrawNodes()
	mainMaze.DrawPath()
	mainMaze.Output()

	// fmt.Println(mainMaze.Graph)

}
