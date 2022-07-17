package graph

import "algorithm-go/structure/unionfind"

func equationsPossible(equations []string) bool {
	uf := unionfind.New(26)

	for _, equation := range equations {
		v1, v2, equal := parser(equation)
		if equal {
			uf.Union(v1, v2)
		}
	}

	for _, equation := range equations {
		v1, v2, equal := parser(equation)
		if !equal {
			if uf.Connected(v1, v2) {
				return false
			}
		}
	}

	return true
}

func parser(equation string) (int, int, bool) {
	v1, v2 := equation[0], equation[3]

	var equal bool
	if equation[1] == '=' {
		equal = true
	}

	return int(v1 - 'a'), int(v2 - 'a'), equal
}
