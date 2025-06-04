package dijkstra

type Edge interface {
	GetSource() Node
	GetDestination() Node
	GetCosts() int
}

func CreateEdge() Edge {
	return &EdgeImpl{}
}

type EdgeImpl struct {
	source      Node
	destination Node
	costs       int
}

func (e *EdgeImpl) GetSource() Node {
	return e.source
}

func (e *EdgeImpl) GetDestination() Node {
	return e.destination
}

func (e *EdgeImpl) GetCosts() int {
	return e.costs
}
