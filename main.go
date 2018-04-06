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

	mainMaze.DrawNodes()
	// print nodes
	for _, n := range mainMaze.Nodes {
		fmt.Println(n.Row, n.Col)
	}

	// print node points col list
	for k, v := range mainMaze.NodePointsCol {
		fmt.Println("Column:", k, "Rows with points", v)
	}
	for k, v := range mainMaze.NodePointsRow {
		fmt.Println("Row:", k, "Columns with points", v)
	}

	// print neighbors
	for k, v := range mainMaze.Neighbors {
		fmt.Println("Node:", k)
		for _, x := range v {
			fmt.Println("Row:", x.Row, "Col:", x.Col)

		}
		fmt.Println("")
	}

	// print maze
	for _, x := range mainMaze.Graph {
		for _, y := range x {
			fmt.Printf("%s", y)
		}
		// fmt.Printf("\n")
	}
}
