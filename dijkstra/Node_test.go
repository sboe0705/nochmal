package dijkstra

import (
	"testing"

	"github.com/sboe0705/assertions"
)

func TestCreateNode(t *testing.T) {
	// when
	node := CreateNode()

	// then
	assertions.AssertInt(t, -1, node.GetDistance(), "Wrong initial distance")
	assertions.AssertArray(t, []Edge{}, node.GetEdges(), "Invalid initial edges")
}
