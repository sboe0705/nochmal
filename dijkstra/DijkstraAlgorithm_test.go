package dijkstra

import (
	"nochmal/logging"
	"testing"

	"github.com/sboe0705/assertions"
)

func TestDetermineNodes(t *testing.T) {
	// given
	nodes := createSampleGraph()

	// when
	determinedNodes := determineNodes(nodes[0])

	// then
	assertions.AssertArray(t, nodes, *determinedNodes, "Not all nodes determined")
}

func TestSortByDistance(t *testing.T) {
	// given
	node0 := CreateNamedNode("Node 0")
	node0.SetDistance(1)
	node1 := CreateNamedNode("Node 1")
	node1.SetDistance(2)
	node2 := CreateNamedNode("Node 2")
	node2.SetDistance(3)
	node3 := CreateNamedNode("Node 3")
	node3.SetDistance(-1)
	nodes := []Node{node0, node1, node2, node3}

	// when
	sortByDistance(nodes)

	// then
	assertions.AssertArray(t, []Node{node0, node1, node2, node3}, nodes, "Array is sorted in wrong order")
}

func TestDetermineDistances(t *testing.T) {
	logging.Init()

	// given
	nodes := createSampleGraph()

	// when (first run)
	DetermineDistances(nodes[0])

	// then
	assertSampleDistances(t, nodes)

	// when (repeated run)
	DetermineDistances(nodes[0])

	// then
	assertSampleDistances(t, nodes)

	// when (run with changed distances)
	nodes[4].SetDistance(0)
	DetermineDistances(nodes[0])

	// then
	assertNodeDistance(t, nodes[0], 0)
	assertNodeDistance(t, nodes[1], 7)
	assertNodeDistance(t, nodes[2], 9)
	assertNodeDistance(t, nodes[3], 6)
	assertNodeDistance(t, nodes[4], 0)
	assertNodeDistance(t, nodes[5], 9)
}

func assertSampleDistances(t *testing.T, nodes []Node) {
	assertNodeDistance(t, nodes[0], 0)
	assertNodeDistance(t, nodes[1], 7)
	assertNodeDistance(t, nodes[2], 9)
	assertNodeDistance(t, nodes[3], 20)
	assertNodeDistance(t, nodes[4], 20)
	assertNodeDistance(t, nodes[5], 11)
}

func createSampleGraph() []Node {
	node0 := CreateNamedNode("Node 0")
	node1 := CreateNamedNode("Node 1")
	node2 := CreateNamedNode("Node 2")
	node3 := CreateNamedNode("Node 3")
	node4 := CreateNamedNode("Node 4")
	node5 := CreateNamedNode("Node 5")

	node0.ConnectWith(node1, 7)
	node0.ConnectWith(node2, 9)
	node0.ConnectWith(node5, 14)
	node1.ConnectWith(node2, 10)
	node1.ConnectWith(node3, 15)
	node2.ConnectWith(node3, 11)
	node2.ConnectWith(node5, 2)
	node3.ConnectWith(node4, 6)
	node4.ConnectWith(node5, 9)

	return []Node{node0, node1, node2, node3, node4, node5}
}

func TestDetermineDistancesWithPresetRootNodeDistance(t *testing.T) {
	logging.Init()

	// given
	nodes := createSampleGraph()

	// when
	nodes[0].SetDistance(5)
	nodes[4].SetDistance(0)
	DetermineDistances(nodes[0])

	// then
	assertNodeDistance(t, nodes[0], 5)
	assertNodeDistance(t, nodes[1], 12)
	assertNodeDistance(t, nodes[2], 11)
	assertNodeDistance(t, nodes[3], 6)
	assertNodeDistance(t, nodes[4], 0)
	assertNodeDistance(t, nodes[5], 9)
}

func assertNodeDistance(t *testing.T, node Node, expectedDistance int) {
	assertions.AssertEquals(t, expectedDistance, node.GetDistance(), "Wrong distance for "+node.ToString())
}
