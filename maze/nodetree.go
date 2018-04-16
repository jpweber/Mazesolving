package maze

func (m *Maze) NodeFinder() {

	// i (row)
	// j (col)

	// also skip the first row it is going to be our border edge
	for i, row := range m.Graph {
		// TODO:
		// fmt.Println("Scanning for node placement on row:", i)
		// columns:
		for j, col := range row {
			// TODO:
			// fmt.Println("Scanning for row placement in Column:", j)
			node := Node{
				Row: int64(i),
				Col: int64(j),
			}

			// Easy exceptions
			// if we hit a wall square at the first step just keep going
			if col == kBlack && j == 0 {
				continue
			}

			// if we hit a wall square and the previous square is also
			// a sqaure, advance, don't place a node
			if col == kBlack && row[j-1] == kBlack {
				continue
			}

			// if we found the start move on for now
			if col == kRed {
				// fmt.Println("Found Start")
				m.PlotNodePoint(node)
				continue
			}

			// if we found the end move on for now
			if col == kGreen {
				// fmt.Println("Found Goal")
				m.Goal = node
				m.WestNeighbor(node)
				m.PlotNodePoint(node)
				continue
			}

			// Easy rules are handled above, now we need to check north south east west positions.
			// and then place decision nodes and make connections as needed.

			// description of rules needed
			// these are corner rules basically
			// if wall to the west wall to the north make node
			// if wall to west and wall to south make node
			// if wall to east and wall to north make node
			// if wall to east and wall to south make node
			// These are path Rules
			// if wall to east and wall to west and path north and south
			// skip node. If we are in this situation there should be node already north
			// and a node will be getting placed south of this position later
			// These are connection rules
			// when placing a node search west for connection and make it
			// then search north for node and connect it.

			// placment logic

			// Corners
			if col == kBlack {
				continue
			}

			// are we next to the start point
			if m.startNeighbor(node) {
				// fmt.Println("found start neighbor", node)
				m.WestNeighbor(node)
				m.PlotNodePoint(node)
			}
			// are we in NW corner
			if m.northWestCorner(node) {
				// fmt.Println("NWC")
				// place node
				m.PlotNodePoint(node)
				continue
			}
			// are we in NE Corner
			if m.northEastCorner(node) {
				// fmt.Println("NEC")
				// place node and make western connection
				m.WestNeighbor(node)
				m.PlotNodePoint(node)
				continue

			}
			// are we in SW corner
			if m.southWestCorner(node) {
				// fmt.Println("SWC")
				// place node and make northern connection
				m.NorthNeighbor(node)
				m.PlotNodePoint(node)
				continue

			}
			// are we in SE corner
			if m.southEastCorner(node) {
				// fmt.Println("SEC")
				// place node and make northern and western connection
				m.NorthNeighbor(node)
				m.WestNeighbor(node)
				m.PlotNodePoint(node)
				continue
			}

			// Path junction points that aren't corners
			if m.topOfColumn(node) {
				// fmt.Println("TC")
				m.WestNeighbor(node)
				m.PlotNodePoint(node)
			}
			if m.bottomOfColumn(node) {
				// fmt.Println("BC")
				m.NorthNeighbor(node)
				m.WestNeighbor(node)
				m.PlotNodePoint(node)
			}
			if m.midEastColJunction(node) {
				// fmt.Println("MEC")
				m.NorthNeighbor(node)
				m.PlotNodePoint(node)
			}
			if m.midWestColJunction(node) {
				// fmt.Println("MWC")
				m.WestNeighbor(node)
				m.NorthNeighbor(node)
				m.PlotNodePoint(node)
			}
		}

	}
}

func (m *Maze) startNeighbor(node Node) bool {
	// fmt.Println("Start neighbor check:", m.Graph[node.Row][node.Col-1])
	if m.Graph[node.Row][node.Col-1] == kRed {
		// fmt.Println("found start neighbor")
		return true
	}
	return false
}

func (m *Maze) northWestCorner(node Node) bool {
	if m.Graph[node.Row-1][node.Col] == kBlack && m.Graph[node.Row][node.Col-1] == kBlack {
		// fmt.Println("In a northwest corner")
		return true
	}
	return false
}
func (m *Maze) southWestCorner(node Node) bool {
	if m.Graph[node.Row+1][node.Col] == kBlack && m.Graph[node.Row][node.Col-1] == kBlack {
		// fmt.Println("In a southwest corner")
		return true
	}
	return false
}

func (m *Maze) northEastCorner(node Node) bool {
	// first check if we are at the end of the columns and exit early
	if node.Col == int64(len(m.Graph[node.Row])-1) {
		return false
	}
	if m.Graph[node.Row-1][node.Col] == kBlack && m.Graph[node.Row][node.Col+1] == kBlack {
		// fmt.Println("In a norteast corner")
		return true
	}
	return false
}
func (m *Maze) southEastCorner(node Node) bool {
	// first check if we are at the end of the columns and exit early
	if node.Col == int64(len(m.Graph[node.Row])-1) {
		return false
	}

	if m.Graph[node.Row+1][node.Col] == kBlack && m.Graph[node.Row][node.Col+1] == kBlack {
		// fmt.Println("In a northwest corner")
		return true
	}
	return false
}
func (m *Maze) topOfColumn(node Node) bool {
	if m.Graph[node.Row-1][node.Col] == kBlack && m.Graph[node.Row+1][node.Col] == kWhite {
		// fmt.Println("Found top of column")
		return true
	}
	return false
}
func (m *Maze) bottomOfColumn(node Node) bool {
	if m.Graph[node.Row-1][node.Col] == kWhite && m.Graph[node.Row+1][node.Col] == kBlack {
		// fmt.Println("Found bottom of column")
		return true
	}
	return false
}
func (m *Maze) midEastColJunction(node Node) bool {
	if m.Graph[node.Row][node.Col+1] == kWhite && m.Graph[node.Row][node.Col-1] == kBlack {
		// fmt.Println("Found east column junction")
		return true
	}
	return false
}
func (m *Maze) midWestColJunction(node Node) bool {
	if m.Graph[node.Row][node.Col-1] == kWhite && m.Graph[node.Row][node.Col+1] == kBlack {
		// fmt.Println("Found west column junction")
		return true
	}
	return false
}

func (m *Maze) NorthNeighbor(node Node) {

	// scan up the row to see if we have a neighbor
	if m.NodePointsCol[node.Col] != nil {
		// fmt.Println("points in this column", m.NodePointsCol[node.Col])
		neighborNode := m.NodePointsCol[node.Col][len(m.NodePointsCol[node.Col])-1]
		// fmt.Println("Found northern neighbor:", node.Col, neighborNode, "origin node:", node)

		// make neighbor connection
		m.ConnectNeighbors(node, neighborNode)
	}

}

func (m *Maze) WestNeighbor(node Node) {

	// scan up the row to see if we have a neighbor
	if m.NodePointsRow[node.Row] != nil {
		neighborNode := m.NodePointsRow[node.Row][len(m.NodePointsRow[node.Row])-1]
		// fmt.Println("Found northern neighbor:", node.Col, neighborNode, "origin node:", node)
		// fmt.Println("Found Western neighbor:", node.Row, neighborNode)
		// make neighbor connection

		m.ConnectNeighbors(node, neighborNode)
	}

}

func (m *Maze) PlotNodePoint(node Node) {

	// fmt.Println("Plotting node point", node)

	// Plot not to mast list of nodes
	// if we don't have any nodes saved yet save it.
	if len(m.Nodes) == 0 {
		m.Nodes = append(m.Nodes, node)
	} else if m.Nodes[len(m.Nodes)-1] != node {
		// if we have more than one node only save it
		// if the previous node is not itself. effectively eliminating dupes
		m.Nodes = append(m.Nodes, node)
	}

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

func (m *Maze) NodeExists(node Node) bool {
	for _, n := range m.Nodes {
		if n == node {
			return true
		}
	}
	return false
}

func (m *Maze) ConnectNeighbors(node1, node2 Node) {
	if node1 == node2 {
		// fmt.Println("found ourselves as neighbor not connecting")
		return
	}
	// log.Println("node1:", node1, "node2:", node2)
	// log.Println("Making neighbor connection")
	if m.Neighbors == nil {
		m.Neighbors = make(map[Node][]Node)
	}
	m.Neighbors[node1] = append(m.Neighbors[node1], node2)
	// and reverse connection
	m.Neighbors[node2] = append(m.Neighbors[node2], node1)
}
