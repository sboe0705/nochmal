package dijkstra

import (
	"fmt"
	"log/slog"
)

type Node[T any] interface {
	GetEdges() []Edge[T]
	AddEdge(edge Edge[T])
	GetDistance() int
	SetDistance(int) error
	ConnectWith(node Node[T], costs int)
	IsVisited() bool
	SetVisited(bool)
	GetValue() T
	ToString() string
}

func CreateNode() Node[any] {
	return &nodeImpl[any]{nil, "", []Edge[any]{}, -1, false}
}

func CreateNamedNode(name string) Node[any] {
	return &nodeImpl[any]{nil, name, []Edge[any]{}, -1, false}
}

func CreateValuedNode[T any](value T, nameOf func(value T) string, initialDistanceOf func(value T) int) Node[T] {
	return &nodeImpl[T]{value, nameOf(value), []Edge[T]{}, initialDistanceOf(value), false}
}

type nodeImpl[T any] struct {
	value    T
	name     string
	edges    []Edge[T]
	distance int
	visited  bool
}

func (n *nodeImpl[T]) GetEdges() []Edge[T] {
	return n.edges
}

func (n *nodeImpl[T]) AddEdge(edge Edge[T]) {
	n.edges = append(n.edges, edge)
}

func (n *nodeImpl[T]) GetDistance() int {
	return n.distance
}

func (n *nodeImpl[T]) SetDistance(distance int) error {
	if n.distance < 0 || n.distance >= distance {
		n.distance = distance
		return nil
	} else {
		slog.Warn("Node distance has been increased after once set!", "node", n.ToString())
		return fmt.Errorf("already set distance cannot be increased")
	}
}

func (n *nodeImpl[T]) ConnectWith(node Node[T], costs int) {
	edge := CreateEdge(n, node, costs)
	n.AddEdge(edge)
	node.AddEdge(edge)
}

func (n *nodeImpl[T]) IsVisited() bool {
	return n.visited
}

func (n *nodeImpl[T]) SetVisited(visited bool) {
	n.visited = visited
}

func (n *nodeImpl[T]) GetValue() T {
	return n.value
}

func (n *nodeImpl[T]) ToString() string {
	if len(n.name) > 0 {
		return n.name
	}
	return fmt.Sprintf("%+v", n)
}
