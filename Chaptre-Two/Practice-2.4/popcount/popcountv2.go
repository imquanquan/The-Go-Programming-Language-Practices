package popcount


func PopCountV2(x uint64) int {
	count := 0
	for i := x; i > 0; i /= 2 {
		count += int(i&1)
	}
	return count
}


