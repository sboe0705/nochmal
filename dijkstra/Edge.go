package dijkstra

type Edge[T any] interface {
	GetSource() Node[T]
	GetDestination() Node[T]
	GetOtherNode(node Node[T]) Node[T]
	GetCosts() int
}

func CreateEdge[T any](source, destination Node[T], costs int) Edge[T] {
	return &edgeImpl[T]{source, destination, costs}
}

type edgeImpl[T any] struct {
	source      Node[T]
	destination Node[T]
	costs       int
}

func (e *edgeImpl[T]) GetSource() Node[T] {
	return e.source
}

func (e *edgeImpl[T]) GetDestination() Node[T] {
	return e.destination
}

func (e *edgeImpl[T]) GetCosts() int {
	return e.costs
}

func (e *edgeImpl[T]) GetOtherNode(node Node[T]) Node[T] {
	if e.GetSource() == node {
		return e.GetDestination()
	} else {
		return e.GetSource()
	}
}
