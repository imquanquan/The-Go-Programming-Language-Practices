package popcount


func PopCountV2(x uint64) int {
	count := 0
	for x != 0 {
		count ++
		x = x & (x - 1)
	}
	return count
}


