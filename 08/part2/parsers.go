package main

import (
	"strings"
)

func extractInputAndMap(s string) ([]string, NodeMap) {
	nodeSlice := strings.Split(s, "\n")
	return strings.Split(nodeSlice[0], ""), parseStringGraphIntoMap(nodeSlice[2:])
}

func (n NodeMap) getStartFinishLocations() ([]string, []string) {
	startLocations := []string{}
	endLocations := []string{}
	for key, _ := range n {
		if string(key[len(key)-1]) == "A" {
			startLocations = append(startLocations, key)
		} else if string(key[len(key)-1]) == "Z" {
			endLocations = append(endLocations, key)
		}
	}
	return startLocations, endLocations
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
