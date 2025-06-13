package dijkstra

type Edge interface {
	GetSource() Node
	GetDestination() Node
	GetOtherNode(node Node) Node
	GetCosts() int
}

func CreateEdge(source, destination Node, costs int) Edge {
	return &edgeImpl{source, destination, costs}
}

type edgeImpl struct {
	source      Node
	destination Node
	costs       int
}

func (e *edgeImpl) GetSource() Node {
	return e.source
}

func (e *edgeImpl) GetDestination() Node {
	return e.destination
}

func (e *edgeImpl) GetCosts() int {
	return e.costs
}

func (e *edgeImpl) GetOtherNode(node Node) Node {
	if e.GetSource() == node {
		return e.GetDestination()
	} else {
		return e.GetSource()
	}
}
