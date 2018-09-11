package Practice_3_10

import (
	"bytes"
)

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
