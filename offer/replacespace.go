package offer

func replaceSpace(s string) string {
	var res, word string
	sep := "%20"

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res += word
			res += sep
			word = ""
		} else {
			word += string(s[i])
		}
	}
	// 尾部单词，若以空格结尾，那么word为空串
	res += word

	return res
}
