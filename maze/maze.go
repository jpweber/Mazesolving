package maze

import (
	"bufio"
	"log"
	"os"
)

type Maze struct {
	Graph [][]string
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

func textDecoder(reader *bufio.Reader) Maze {
	var line string
	var err error
	var maze Maze
	i := 0
	for {
		row := []string{}
		line, err = reader.ReadString('\n')
		if len(line) > 5 {
			for _, char := range line {
				row = append(row, string(char))
			}

			log.Println(len(row))
			maze.Graph = append(maze.Graph, row[2:])
		}
		i++
		if err != nil {
			break
		}
	}
	maze.findStart()
	maze.findGoal()
	// maze.Graph = graph
	return maze
}

// Read  file and pass it on to the maze decoder
func Read(fileName string) Maze {
	// open file and read contents
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	return textDecoder(reader)

}

func (m *Maze) findStart() {
	// init start map
	m.Start = make(map[int64]int64)
loop:
	for i, row := range m.Graph {
		for j, col := range row {
			if col == "S" {
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
			if col == "G" {
				log.Println("Found Goal:", i, ",", j)
				m.Goal[int64(i)] = int64(j)
				break loop
			}
		}
	}
}
