package queue

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	headCopy := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}

	copyMap := map[*Node]*Node{
		node: headCopy,
	}

	queue := []*Node{node}
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nbr := range curr.Neighbors {
			currCopy := copyMap[curr]

			if nbrCopy, ok := copyMap[nbr]; ok {
				currCopy.Neighbors = append(currCopy.Neighbors, nbrCopy)
			} else {
				nbrCopy = &Node{
					Val:       nbr.Val,
					Neighbors: make([]*Node, 0),
				}

				copyMap[nbr] = nbrCopy
				currCopy.Neighbors = append(currCopy.Neighbors, nbrCopy)

				queue = append(queue, nbr)
			}
		}
	}

	return headCopy
}
