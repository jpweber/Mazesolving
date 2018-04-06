package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpweber/mazes/maze"
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
	mainMaze.NodeFinder()

	mainMaze.PlotNodes()
	// print nodes
	for _, n := range mainMaze.Nodes {
		fmt.Println(n.Row, n.Col, n.Neighbors)
	}
	// print maze
	for _, x := range mainMaze.Graph {
		for _, y := range x {
			fmt.Printf("%s", y)
		}
		// fmt.Printf("\n")
	}
}
