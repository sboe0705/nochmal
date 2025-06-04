package dijkstra

type Node interface {
	GetEdges() []Edge
	GetDistance() int
	SetDistance(int)
}

func CreateNode() Node {
	return &nodeImpl{[]Edge{}, -1}
}

type nodeImpl struct {
	edges    []Edge
	distance int
}

func (n *nodeImpl) GetEdges() []Edge {
	return n.edges
}

func (n *nodeImpl) GetDistance() int {
	return n.distance
}

func (n *nodeImpl) SetDistance(distance int) {
	n.distance = distance
}
