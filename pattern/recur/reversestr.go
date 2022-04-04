package recur

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}

	s[0], s[len(s)] = s[len(s)], s[0]
	reverseString(s[1:len(s)-1])
}
