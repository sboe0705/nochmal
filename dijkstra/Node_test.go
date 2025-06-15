package dijkstra

import (
	"fmt"
	"testing"

	"github.com/sboe0705/assertions"
)

func TestCreateNode(t *testing.T) {
	// when
	node := CreateNode()

	// then
	assertions.AssertEquals(t, -1, node.GetDistance(), "Wrong initial distance")
	assertions.AssertArray(t, []Edge[any]{}, node.GetEdges(), "Invalid initial edges")
	assertions.AssertFalse(t, node.IsVisited(), "Invalid initial visited state")
	assertions.AssertEquals(t, "&{value:<nil> name: edges:[] distance:-1 visited:false}", node.ToString(), "Wrong initial name")
	assertions.AssertEquals(t, nil, node.GetValue(), "Invalid value")
}

func TestCreateNamedNode(t *testing.T) {
	nodeName := "Node 1"

	// when
	node := CreateNamedNode(nodeName)

	// then
	assertions.AssertEquals(t, -1, node.GetDistance(), "Wrong initial distance")
	assertions.AssertArray(t, []Edge[any]{}, node.GetEdges(), "Invalid initial edges")
	assertions.AssertFalse(t, node.IsVisited(), "Invalid initial visited state")
	assertions.AssertEquals(t, nodeName, node.ToString(), "Invalid name")
	assertions.AssertEquals(t, nil, node.GetValue(), "Invalid value")
}

func TestCreateValuedNode(t *testing.T) {
	// given
	nameOf := func(value int) string {
		return fmt.Sprintf("Node %d", value)
	}

	// when
	node := CreateValuedNode(123, nameOf, distanceOf)

	// then
	assertions.AssertEquals(t, 12, node.GetDistance(), "Wrong initial distance")
	assertions.AssertArray(t, []Edge[int]{}, node.GetEdges(), "Invalid initial edges")
	assertions.AssertFalse(t, node.IsVisited(), "Invalid initial visited state")
	assertions.AssertEquals(t, "Node 123", node.ToString(), "Invalid name")
	assertions.AssertEquals(t, 123, node.GetValue(), "Wrong or missing value")
}

func distanceOf(value int) int {
	return value / 10
}

func TestSetDistance(t *testing.T) {
	// given
	node := CreateNode() // distance not yet set

	// when ... then
	assertions.AssertEquals(t, nil, node.SetDistance(5), "Invalid error returned")

	// when (increased already set distance)
	err := node.SetDistance(6)
	assertions.AssertTrue(t, err != nil, "No error returned")
	assertions.AssertEquals(t, "already set distance cannot be increased", err.Error(), "Wrong error message")
}

func TestConnectWith(t *testing.T) {
	node1 := CreateNode()
	node2 := CreateNode()
	costs := 5

	// when
	node1.ConnectWith(node2, costs)

	// then
	assertions.AssertEquals(t, 1, len(node1.GetEdges()), "Missing edge on node 1")
	assertions.AssertEquals(t, 1, len(node2.GetEdges()), "Missing edge on node 2")

	assertions.AssertEquals(t, node1.GetEdges()[0], node2.GetEdges()[0], "Different edge in connected nodes")
	assertions.AssertEquals(t, costs, node1.GetEdges()[0].GetCosts(), "Wrong costs")
}

func TestSetVisited(t *testing.T) {
	// given
	node := CreateNode()

	// when
	node.SetVisited(true)

	// then
	assertions.AssertTrue(t, node.IsVisited(), "Invalid initial visited state")
}
