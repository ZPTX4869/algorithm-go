package stack

import (
	"log"
	"strconv"
	"strings"
)

// 双栈解法
func decodeString(s string) string {
	var res string
	numStk := make([]int, 0)
	ltrStk := make([]string, 0)

	ptr := 0
	for ptr < len(s) {
		if s[ptr] >= '0' && s[ptr] <= '9' {
			digit := getDigit(s, &ptr)
			numStk = append(numStk, digit)
			continue
		} else if s[ptr] >= 'a' && s[ptr] <= 'z' || s[ptr] >= 'A' && s[ptr] <= 'Z' || s[ptr] == '[' {
			ltrStk = append(ltrStk, string(s[ptr]))
		} else {
			var subRes string

			i := len(ltrStk) - 1
			for ltrStk[i] != "[" {
				subRes = ltrStk[i] + subRes
				i--
			}

			num := numStk[len(numStk)-1]
			numStk = numStk[:len(numStk)-1]
			subRes = strings.Repeat(subRes, num)

			ltrStk = ltrStk[:i]
			ltrStk = append(ltrStk, strings.Split(subRes, "")...)
		}
		ptr++
	}

	res = strings.Join(ltrStk, "")
	return res
}

// 递归解法
//func decodeString(s string) string {
//	ptr := 0
//	return getString(s, &ptr)
//}

func getDigit(s string, ptr *int) int {
	digitStr := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		digitStr += string(s[*ptr])
	}

	digit, err := strconv.Atoi(digitStr)
	if err != nil {
		log.Fatalln(digitStr)
	}

	return digit
}

func getString(s string, ptr *int) string {
	if *ptr == len(s) || s[*ptr] == ']' {
		return ""
	}

	ret := ""
	curr := s[*ptr]
	rep := 1
	if curr >= '0' && curr <= '9' {
		rep = getDigit(s, ptr)
		*ptr++

		sub := getString(s, ptr)
		*ptr++

		ret = strings.Repeat(sub, rep) + getString(s, ptr)
	} else {
		*ptr++
		ret = string(curr) + getString(s, ptr)
	}

	return ret
}
