package maze

import (
	"fmt"
	"log"
)

const wallChar = "#"
const pathChar = " "

func (m *Maze) NodeFinder() {

	// i (row)
	// j (col)

	// also skip the first row it is going to be our border edge
	for i, row := range m.Graph {
		// TODO:
		fmt.Println("Scanning for node placement on row:", i)
		nodePlaced := false
		for j, col := range row {
			// TODO:
			fmt.Println("Scanning for row placement in Column:", j)

			// if we hit a wall square at the first step just keep going
			if col == wallChar && j == 0 {
				continue
			}

			// if we hit a wall square and the previous square is also
			// a sqaure, advance, don't place a node
			if col == wallChar && row[j-1] == wallChar {
				continue
			}

			// place first node point found on row
			// ever path space after this doesn't need a node until we hit a decision point
			if col == pathChar && nodePlaced == false {
				node := Node{
					Row: int64(i),
					Col: int64(j),
				}

				nodePlaced = true
				m.NorthNeighbor(node)
				m.PlotNodePoint(node)

				// TODO:
				fmt.Println("Placing init Node at Row", i, "Column", j)
				continue
			}

			// just walking along the path checking up for neighbors
			if col == pathChar {
				log.Println(col, i, j)
				node := Node{
					Row: int64(i),
					Col: int64(j),
				}
				if m.NorthNeighbor(node) {
					m.WestNeighbor(node)
					m.PlotNodePoint(node)
				}
				continue
			}

			// hit wall and set node one space back
			if col == wallChar {
				fmt.Println("hit wall")
				node := Node{
					Row: int64(i),
					Col: int64(j - 1),
				}
				// check if we already know this point before moving one
				if m.NodeExists(node) {
					fmt.Println("Node already exists, do nothing", node)
					continue
				}
				m.NorthNeighbor(node)
				m.WestNeighbor(node)
				m.PlotNodePoint(node)

				fmt.Println("Placing additional Node at Row", i, "Column", j-1)
				continue
			}
		}

	}
}

func (m *Maze) NorthNeighbor(node Node) bool {
	// row = node.Row
	// col = node.Col
	// stay in same row, decrement column, check for wall char
	// first bail early if we hit a wall
	if m.Graph[node.Row-1][node.Col] == wallChar {
		//TODO:
		fmt.Println("wall north of me, don't checking north")
		return false
	}

	// scan up the row to see if we have a neighbor
	if m.NodePointsCol[node.Col] != nil {
		neighborNode := m.NodePointsCol[node.Col][len(m.NodePointsCol[node.Col])-1]
		fmt.Println("Found northern neighbor:", node.Col, neighborNode)
		// make neighbor connection

		// m.Neighbors[node] = append(m.Neighbors[node], neighborNode)
		// m.Neighbors[node] = map[Node]bool{neighborNode: true}
		m.ConnectNeighbors(node, neighborNode)
		return true
	}

	return false
}

func (m *Maze) WestNeighbor(node Node) bool {
	// row = node.Row
	// col = node.Col
	// stay in same row, decrement column, check for wall char
	// first bail early if we hit a wall
	if m.Graph[node.Row][node.Col-1] == wallChar {
		//TODO:
		fmt.Println("wall west of me, done checking west")
		return false
	}

	// scan up the row to see if we have a neighbor
	if m.NodePointsRow[node.Row] != nil {
		neighborNode := m.NodePointsRow[node.Row][len(m.NodePointsRow[node.Row])-1]
		fmt.Println("Found Western neighbor:", node.Row, neighborNode)
		// make neighbor connection

		// m.Neighbors[node] = append(m.Neighbors[node], neighborNode)
		// m.Neighbors[node] = map[Node]bool{neighborNode: true}
		m.ConnectNeighbors(node, neighborNode)
		return true
	}

	return false
}

func (m *Maze) PlotNodePoint(node Node) {

	fmt.Println("Plotting node point")

	// Plot not to mast list of nodes
	m.Nodes = append(m.Nodes, node)

	// plot node in column map
	if m.NodePointsCol == nil {
		m.NodePointsCol = make(map[int64][]Node)
	}
	m.NodePointsCol[int64(node.Col)] = append(m.NodePointsCol[int64(node.Col)], node)

	// plot node in row map
	if m.NodePointsRow == nil {
		m.NodePointsRow = make(map[int64][]Node)
	}
	m.NodePointsRow[int64(node.Row)] = append(m.NodePointsRow[int64(node.Row)], node)
}

func (m *Maze) DrawNodes() {
	for _, n := range m.Nodes {
		m.Graph[n.Row][n.Col] = "*"
	}
}

func (m *Maze) NodeExists(node Node) bool {
	for _, n := range m.Nodes {
		if n == node {
			return true
		}
	}
	return false
}

func (m *Maze) ConnectNeighbors(node1, node2 Node) {
	log.Println("node1:", node1, "node2:", node2)
	log.Println("Making neighbor connection")
	if m.Neighbors == nil {
		m.Neighbors = make(map[Node][]Node)
	}
	m.Neighbors[node1] = append(m.Neighbors[node1], node2)
	// and reverse connection
	m.Neighbors[node2] = append(m.Neighbors[node2], node1)
}
