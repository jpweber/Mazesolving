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

			// place node and/or keep moving along a clear path
			if col == pathChar && nodePlaced == false {
				node := Node{
					Row: int64(i),
					Col: int64(j),
				}
				nodePlaced = true
				m.Nodes = append(m.Nodes, node)
				// TODO:
				fmt.Println("Placing init Node at Row", i, "Column", j)
			}

			// hit wall and set node one space back
			if col == wallChar {
				node := Node{
					Row: int64(i),
					Col: int64(j - 1),
				}
				m.Nodes = append(m.Nodes, node)
				fmt.Println("Placing additional Node at Row", i, "Column", j-1)

				// if we hit a wall in a row, start looking down that colume until
				// we hit a wall at another row and plot that point
				subColidx := j - 1
			subColumn:
				for {

					if m.Graph[i][subColidx] == wallChar {
						node := Node{
							Row: int64(i - 1),
							Col: int64(subColidx),
						}
						m.Nodes = append(m.Nodes, node)
						fmt.Println("Placing additional Node at Row", i, "Column", subColidx)
						break subColumn
					} else {
						i++
					}
				}
			}
		}

	}
}

func (m *Maze) PlotNodes() {
	for _, n := range m.Nodes {
		m.Graph[n.Row][n.Col] = "*"
	}
}
