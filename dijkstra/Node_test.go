package dijkstra

import (
	"testing"

	"github.com/sboe0705/assertions"
)

func TestCreateNode(t *testing.T) {
	// when
	node := CreateNode()

	// then
	assertions.AssertEquals(t, -1, node.GetDistance(), "Wrong initial distance")
	assertions.AssertArray(t, []Edge{}, node.GetEdges(), "Invalid initial edges")
	assertions.AssertFalse(t, node.IsVisited(), "Invalid initial visited state")
	assertions.AssertEquals(t, "&{name: edges:[] distance:-1 visited:false}", node.ToString(), "Wrong initial name")
}

func TestCreateNamedNode(t *testing.T) {
	nodeName := "Node 1"

	// when
	node := CreateNamedNode(nodeName)

	// then
	assertions.AssertEquals(t, -1, node.GetDistance(), "Wrong initial distance")
	assertions.AssertArray(t, []Edge{}, node.GetEdges(), "Invalid initial edges")
	assertions.AssertFalse(t, node.IsVisited(), "Invalid initial visited state")
	assertions.AssertEquals(t, nodeName, node.ToString(), "Invalid name")
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
