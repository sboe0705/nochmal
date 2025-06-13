package dijkstra

import (
	"testing"

	"github.com/sboe0705/assertions"
)

func TestCreateEdge(t *testing.T) {
	// given
	node1 := CreateNode()
	node2 := CreateNode()
	costs := 4

	// when
	edge := CreateEdge(node1, node2, costs)

	// then
	assertions.AssertEquals(t, costs, edge.GetCosts(), "Wrong costs")
	assertions.AssertEquals(t, node1, edge.GetSource(), "Wrong source node")
	assertions.AssertEquals(t, node2, edge.GetDestination(), "Wrong destination node")
}

func TestGetOtherNode(t *testing.T) {
	// given
	node1 := CreateNode()
	node2 := CreateNode()
	edge := CreateEdge(node1, node2, 4)

	// when ... then
	assertions.AssertEquals(t, node1, edge.GetOtherNode(node2), "Wrong other node")
	assertions.AssertEquals(t, node2, edge.GetOtherNode(node1), "Wrong other node")
}
