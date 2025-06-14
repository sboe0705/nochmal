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

func TestDetermineDistancesFrom(t *testing.T) {
	logging.Init()

	// given
	nodes := createSampleGraph()

	// when
	DetermineDistancesFrom(nodes[0])

	// then
	assertions.AssertEquals(t, 0, nodes[0].GetDistance(), "Wrong distance for node1")
	assertions.AssertEquals(t, 7, nodes[1].GetDistance(), "Wrong distance for node1")
	assertions.AssertEquals(t, 9, nodes[2].GetDistance(), "Wrong distance for node2")
	assertions.AssertEquals(t, 20, nodes[3].GetDistance(), "Wrong distance for node3")
	assertions.AssertEquals(t, 20, nodes[4].GetDistance(), "Wrong distance for node4")
	assertions.AssertEquals(t, 11, nodes[5].GetDistance(), "Wrong distance for node5")
}

func TestDetermineDistancesFromWithPresetDistances(t *testing.T) {
	logging.Init()

	// given
	nodes := createSampleGraph()

	// when
	nodes[0].SetDistance(5) // ignore the preset distance in root node
	nodes[4].SetDistance(0)
	DetermineDistancesFrom(nodes[0])

	// then
	assertions.AssertEquals(t, 0, nodes[0].GetDistance(), "Wrong distance for node1")
	assertions.AssertEquals(t, 7, nodes[1].GetDistance(), "Wrong distance for node1")
	assertions.AssertEquals(t, 9, nodes[2].GetDistance(), "Wrong distance for node2")
	assertions.AssertEquals(t, 6, nodes[3].GetDistance(), "Wrong distance for node3")
	assertions.AssertEquals(t, 0, nodes[4].GetDistance(), "Wrong distance for node4")
	assertions.AssertEquals(t, 9, nodes[5].GetDistance(), "Wrong distance for node5")
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
