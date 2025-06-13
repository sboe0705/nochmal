package dijkstra

import (
	"nochmal/logging"
	"testing"

	"github.com/sboe0705/assertions"
)

func TestDetermineDistances(t *testing.T) {
	logging.Init()

	// given
	node1 := CreateNamedNode("Node 1")
	node2 := CreateNamedNode("Node 2")
	node3 := CreateNamedNode("Node 3")
	node4 := CreateNamedNode("Node 4")
	node5 := CreateNamedNode("Node 5")
	node6 := CreateNamedNode("Node 6")

	node1.ConnectWith(node2, 7)
	node1.ConnectWith(node3, 9)
	node1.ConnectWith(node6, 14)
	node2.ConnectWith(node3, 10)
	node2.ConnectWith(node4, 15)
	node3.ConnectWith(node4, 11)
	node3.ConnectWith(node6, 2)
	node4.ConnectWith(node5, 6)
	node5.ConnectWith(node6, 9)

	// when
	DetermineDistances(node1)

	// then
	assertions.AssertEquals(t, 0, node1.GetDistance(), "Wrong distance for node1")
	assertions.AssertEquals(t, 7, node2.GetDistance(), "Wrong distance for node2")
	assertions.AssertEquals(t, 9, node3.GetDistance(), "Wrong distance for node3")
	assertions.AssertEquals(t, 20, node4.GetDistance(), "Wrong distance for node4")
	assertions.AssertEquals(t, 20, node5.GetDistance(), "Wrong distance for node5")
	assertions.AssertEquals(t, 11, node6.GetDistance(), "Wrong distance for node6")
}
