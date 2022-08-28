package bfs

func turnUp(code string, pos int) string {
	bytes := []byte(code)

	if bytes[pos] == '9' {
		bytes[pos] = '0'
	} else {
		bytes[pos] += 1
	}

	return string(bytes)
}

func turnDown(code string, pos int) string {
	bytes := []byte(code)

	if bytes[pos] == '0' {
		bytes[pos] = '9'
	} else {
		bytes[pos] -= 1
	}

	return string(bytes)
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, v := range deadends {
		if v == "0000" {
			return -1
		}
		visited[v] = true
	}

	step := 0
	queue := []string{"0000"}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			code := queue[i]

			if code == target {
				return step
			}

			for i := 0; i < 4; i++ {
				upCode := turnUp(code, i)
				if _, ok := visited[upCode]; !ok {
					queue = append(queue, upCode)
					visited[upCode] = true
				}

				downCode := turnDown(code, i)
				if _, ok := visited[downCode]; !ok {
					queue = append(queue, downCode)
					visited[downCode] = true
				}
			}
		}

		queue = queue[size:]
		step += 1
	}

	return -1
}
