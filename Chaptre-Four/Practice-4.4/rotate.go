package Practice_4_4

func Rotate(s []int, c int)  {
	a := make([]int, c-1)
	copy(a, s[:c-1])
	copy(s[:], s[c-1:])
	copy(s[c+1:], a)
}
