package Practice_4_1

func BitCount(sha1, sha2 [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {
		if sha1[i] != sha2[i] {
			count ++
		}
	}
	return count
}
