package Practice_3_11

import (
	"bytes"
	"strings"
)


func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n < 3 {
		return s
	}
	for n >= 3 {
		n -= 3
	}
	buf.WriteString(s[0:n])
	if n != 0 {
		buf.WriteString(",")
	}

	for n <= len(s) - 3 {
		buf.WriteString(s[n:n+3] + ",")
		n += 3
	}
	ss := buf.String()
	return ss[:len(ss) - 1]
}

func CommaV2(s string) string {
	mark := 0
	var ss string
	if s[0] == '-' {
		s = s[1:]
		mark = 1
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		ss = Comma(s[:dot]) + "." + Reverse(Comma(Reverse(s[dot+1:])))
	} else {
		ss = Comma(s)
	}
	if mark == 1{
		return "-" + ss
	} else {
		return ss
	}
}