package Practice_4_5

func NoneEmpty(s []string) []string {
	count := len(s)
	for i := range s {
		if i == len(s)-1 {
			return s
		}
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
			count--
		}
	}
	return s
}
