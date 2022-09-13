package debug

func Space(level int) string {
	res := ""

	for i := 0; i < level; i++ {
		res += "  "
	}

	return res
}
