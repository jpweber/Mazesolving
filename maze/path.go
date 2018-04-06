package maze

import "fmt"

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
				node := Node{
					Row: int64(i),
					Col: int64(j),
				}
				if m.NorthNeighbor(node) {
					m.PlotNodePoint(node)
				}
			}

			// hit wall and set node one space back
			if col == wallChar {
				node := Node{
					Row: int64(i),
					Col: int64(j - 1),
				}
				m.NorthNeighbor(node)
				m.PlotNodePoint(node)

				fmt.Println("Placing additional Node at Row", i, "Column", j-1)
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
	if m.NodePoints[node.Col] != nil {
		fmt.Println("Found northern neighbor:", node.Col, m.NodePoints[node.Col][len(m.NodePoints[node.Col])-1])
		return true
	}

	return false
}
func (m *Maze) PlotNodePoint(node Node) {
	fmt.Println("Plotting node point")
	m.Nodes = append(m.Nodes, node)
	if m.NodePoints == nil {
		m.NodePoints = make(map[int64][]int64)
	}
	m.NodePoints[int64(node.Col)] = append(m.NodePoints[int64(node.Col)], node.Row)
}

func (m *Maze) DrawNodes() {
	for _, n := range m.Nodes {
		m.Graph[n.Row][n.Col] = "*"
	}
}
