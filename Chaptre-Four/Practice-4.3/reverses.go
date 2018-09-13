package Practice_4_3

func Reverses(s *[10]int)  {
	for i, j := 0, 9; i < j; j, i = j - 1, i + 1 {
		s[i], s[j] = s[j], s[i]
	}
}
