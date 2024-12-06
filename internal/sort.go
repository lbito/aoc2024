package utils

func TopologicalSort(adjencyMap map[int][]int, data []int) []int {
	//STEP 1: scope the adjencymap to only include the nodes in data
	scopedAdjencyMap := make(map[int][]int)
	setOfEdges := make(map[int]bool)
	inDegree := make(map[int]int)
	// make set of edges for quick lookup
	for _, val := range data {
		setOfEdges[val] = true
	}
	for node, edgeList := range adjencyMap {
		//if the node is not in the data, skip it
		if !setOfEdges[node] {
			continue
		}
		if _, exists := inDegree[node]; !exists {
			inDegree[node] = 0
		}
		if _, exists := scopedAdjencyMap[node]; !exists {
			scopedAdjencyMap[node] = []int{}
		}
		for _, edge := range edgeList {
			//if the edge is not in data then skip it
			if !setOfEdges[edge] {
				continue
			}
			//node and edge are both in data, add the edge to the scopedAdjencyMap
			scopedAdjencyMap[node] = append(scopedAdjencyMap[node], edge)
			inDegree[edge]++
		}
	}

	queue := make([]int, 0)
	result := make([]int, 0)

	//add the nodes with 0 inDegree to the queue
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
			delete(inDegree, node)
		}
	}

	//Khan's algorithm
	for len(queue) > 0 {
		currentNode := queue[0]
		result = append(result, currentNode)
		degreesToDecrement := scopedAdjencyMap[currentNode]

		for _, node := range degreesToDecrement {
			inDegree[node]--
		}

		for node, degree := range inDegree {
			if degree == 0 {
				delete(inDegree, node)
				queue = append(queue, node)
			}
		}
		queue = queue[1:]

	}
	return result
}
