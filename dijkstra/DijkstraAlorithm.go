package dijkstra

import (
	"cmp"
	"log/slog"
	"slices"
)

const debugPrefix = "Dijkstra Algoritm - "

func DetermineDistances(rootNode Node) {
	rootNode.SetDistance(0)
	nodesToVisit := &[]Node{rootNode}
	visitedNodes := &[]Node{}
	for len(*nodesToVisit) > 0 {
		nextNode := getNextNodeToVisit(nodesToVisit)
		visit(nextNode, nodesToVisit, visitedNodes)
	}
}

func getNextNodeToVisit(nodesToVisit *[]Node) Node {
	sortByDistance(*nodesToVisit)
	nextNode := (*nodesToVisit)[0]
	nextNode.SetVisited(true)
	*nodesToVisit = (*nodesToVisit)[1:]
	return nextNode
}

func visit(node Node, nodesToVisit, visitedNodes *[]Node) {
	slog.Debug(debugPrefix+"visiting node", "node", node.ToString())
	for i := range node.GetEdges() {
		edge := node.GetEdges()[i]
		nextNode := edge.GetOtherNode(node)
		markNodeForVisit(nextNode, nodesToVisit, visitedNodes)
		if !isAlreadyVisited(nextNode, visitedNodes) {
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

func markNodeForVisit(node Node, nodesToVisit, visitedNodes *[]Node) {
	if isMarkedForVisit(node, nodesToVisit) {
		slog.Debug(debugPrefix+"node is already marked for visit", "node", node.ToString())
		return
	}
	if isAlreadyVisited(node, visitedNodes) {
		slog.Debug(debugPrefix+"node has already been visited", "node", node.ToString())
		return
	}
	slog.Debug(debugPrefix+"node is marked for visit", "node", node.ToString())
	*nodesToVisit = append(*nodesToVisit, node)
}

func isMarkedForVisit(node Node, nodesToVisit *[]Node) bool {
	for i := range *nodesToVisit {
		if (*nodesToVisit)[i] == node {
			return true
		}
	}
	return false
}

func isAlreadyVisited(node Node, visitedNodes *[]Node) bool {
	for i := range *visitedNodes {
		if (*visitedNodes)[i] == node {
			return true
		}
	}
	return false
}

func sortByDistance(nodes []Node) {
	slices.SortFunc(nodes, func(a, b Node) int {
		return cmp.Compare(a.GetDistance(), b.GetDistance())
	})
}
