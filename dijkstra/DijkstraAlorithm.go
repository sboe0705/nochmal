package dijkstra

import (
	"cmp"
	"log/slog"
	"slices"
)

const debugPrefix = "Dijkstra Algoritm - "

func DetermineDistances(rootNode Node[any]) {
	if rootNode.GetDistance() < 0 {
		rootNode.SetDistance(0)
	}
	nodesToVisit := determineNodes(rootNode)
	visitedNodes := &[]Node[any]{}
	for len(*nodesToVisit) > 0 {
		nextNode := getNextNodeToVisit(nodesToVisit)
		visit(nextNode, visitedNodes)
	}
}

func determineNodes(node Node[any]) *[]Node[any] {
	nodes := &[]Node[any]{node}
	determineNodesRecursively(node, nodes)
	return nodes
}

func determineNodesRecursively(node Node[any], nodes *[]Node[any]) {
	for i := range node.GetEdges() {
		edge := node.GetEdges()[i]
		nextNode := edge.GetOtherNode(node)
		if addIfNotContains(nextNode, nodes) {
			// visit next node recursively, only if it was added to list
			determineNodesRecursively(nextNode, nodes)
		}
	}
}

func getNextNodeToVisit(nodesToVisit *[]Node[any]) Node[any] {
	sortByDistance(*nodesToVisit)
	nextNode := (*nodesToVisit)[0]
	*nodesToVisit = (*nodesToVisit)[1:]
	return nextNode
}

func visit(node Node[any], visitedNodes *[]Node[any]) {
	node.SetVisited(true)
	slog.Debug(debugPrefix+"visiting node", "node", node.ToString())
	for i := range node.GetEdges() {
		edge := node.GetEdges()[i]
		nextNode := edge.GetOtherNode(node)
		if !contains(nextNode, visitedNodes) {
			oldDistance := nextNode.GetDistance()
			newDistance := node.GetDistance() + edge.GetCosts()
			if oldDistance == -1 {
				slog.Debug(debugPrefix+"setting distance", "node", nextNode.ToString(), "newDistance", newDistance)
				nextNode.SetDistance(newDistance)
			} else if oldDistance > newDistance {
				slog.Debug(debugPrefix+"decreasing distance", "node", nextNode.ToString(), "oldDistance", oldDistance, "newDistance", newDistance)
				nextNode.SetDistance(newDistance)
			}
		} else {
			slog.Debug(debugPrefix+"keeping distance as node has alredy been visited", "node", nextNode.ToString())
		}
	}
	*visitedNodes = append(*visitedNodes, node)
}

func addIfNotContains(node Node[any], nodes *[]Node[any]) bool {
	if contains(node, nodes) {
		return false
	}
	*nodes = append(*nodes, node)
	return true
}

func contains(node Node[any], nodes *[]Node[any]) bool {
	for i := range *nodes {
		if (*nodes)[i] == node {
			return true
		}
	}
	return false
}

func sortByDistance(nodes []Node[any]) {
	slices.SortFunc(nodes, func(a, b Node[any]) int {
		if a.GetDistance() < 0 {
			return 1
		}
		if b.GetDistance() < 0 {
			return -1
		}
		return cmp.Compare(a.GetDistance(), b.GetDistance())
	})
}
