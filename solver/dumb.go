package solver

import (
	"fmt"

	"github.com/jpweber/mazes/maze"
)

var visited []maze.Node
var usedChoice []maze.Node
var path []maze.Node
var lastMultiChoice []maze.Node
var goalReached bool

func DumbAlg(m maze.Maze) []maze.Node {
	// for _, n := range m.Nodes {
	// 	fmt.Println(n.Row, n.Col)
	// 	fmt.Println(m.Neighbors[n])
	// 	fmt.Println("")
	// }

	start := m.Nodes[0]
	visit(start, m)
	return path
}

func visit(n maze.Node, m maze.Maze) {
	// early exit
	if goalReached {
		return
	}
	// technically a prepend. building a stack FILO
	visited = append([]maze.Node{n}, visited...)
	path = append(path, n)

	// fmt.Println("Visiting:", n)
	choices := m.Neighbors[n]
	// fmt.Println("Choices: ", choices)

	// check all choices to see if we found the goal
	for _, c := range choices {
		if goal(c, m) {
			fmt.Println("yay we made it")
			goalReached = true
			// fmt.Println(usedChoice)
			return
		}
	}

	// check if we are out of choices. Likely this is a dead end
	if len(choices) == 0 {
		// fmt.Println("Out of choices")
		// fmt.Println("I should go back to", lastMultiChoice, "and try a new path")
		goBackTo := lastMultiChoice[0]
		// prune our path to remove branches are aren't using
		path = prune(goBackTo, path)
		// add the used choice to list so we don't reuse it again
		usedChoice = append(usedChoice, goBackTo)
		// remove the node we are going back to from the stack
		lastMultiChoice = lastMultiChoice[1:]
		visit(goBackTo, m)
	}

	// if we have more than one choice log it as a return point
	if len(choices) > 2 && !contains(n, usedChoice) {
		// fmt.Println("Logging multi choice path")
		lastMultiChoice = append([]maze.Node{n}, lastMultiChoice...)
	}

	// pick next unvisited node
	nextNode := maze.Node{}
	for _, c := range choices {

		if !contains(c, visited) {
			// fmt.Println("Setting next node to:", c)
			nextNode = c
			break
		}
	}
	visit(nextNode, m)

}

func contains(choice maze.Node, visited []maze.Node) bool {
	for _, v := range visited {
		if choice == v {
			return true
		}
	}
	return false
}
func goal(n maze.Node, m maze.Maze) bool {

	if n == m.Goal {
		return true
	}
	return false
}

func prune(n maze.Node, path []maze.Node) []maze.Node {
	var prunedPath []maze.Node
	// new visited nodes are append to the bottom
	// search the slice in reverse for wher to prune
	for i, node := range path {
		if node == n {
			prunedPath = path[:i]
			// fmt.Println(prunedPath)
			return prunedPath
		}
	}

	return prunedPath
}
