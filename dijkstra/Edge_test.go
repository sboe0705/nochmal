package dijkstra

import (
	"testing"

	"github.com/sboe0705/assertions"
)

func TestCreateEdge(t *testing.T) {
	node1 := CreateNode()
	node2 := CreateNode()
	costs := 4

	// when
	edge := CreateEdge(node1, node2, costs)

	// then
	assertions.AssertInt(t, costs, edge.GetCosts(), "Wrong costs")
	assertions.AssertObject(t, node1, edge.GetSource(), "Wrong source node")
	assertions.AssertObject(t, node2, edge.GetDestination(), "Wrong destination node")
}
