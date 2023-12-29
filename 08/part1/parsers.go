package main

import (
	"strings"
)

func extractInputAndMap(s string) ([]string, NodeMap) {
	nodeSlice := strings.Split(s, "\n")
	return strings.Split(nodeSlice[0], ""), parseStringGraphIntoMap(nodeSlice[2:])
}
func parseStringGraphIntoMap(s []string) NodeMap {
	nodeMap := map[string]LeftRight{}
	for _, node := range s {
		split := strings.Split(node, " = ")
		key, unprocessedMap := split[0], split[1]
		mapSlice := strings.Split(string(unprocessedMap[1:9]), ", ")
		nodeMap[key] = LeftRight{mapSlice[0], mapSlice[1]}
	}
	return nodeMap
}
