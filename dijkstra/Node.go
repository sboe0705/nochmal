package dijkstra

type Node interface {
	GetEdges() []Edge
	AddEdge(edge Edge)
	GetDistance() int
	SetDistance(int)
	ConnectWith(node Node, costs int)
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

func (n *nodeImpl) AddEdge(edge Edge) {
	n.edges = append(n.edges, edge)
}

func (n *nodeImpl) GetDistance() int {
	return n.distance
}

func (n *nodeImpl) SetDistance(distance int) {
	n.distance = distance
}

func (n *nodeImpl) ConnectWith(node Node, costs int) {
	edge := CreateEdge(n, node, costs)
	n.AddEdge(edge)
	node.AddEdge(edge)
}
