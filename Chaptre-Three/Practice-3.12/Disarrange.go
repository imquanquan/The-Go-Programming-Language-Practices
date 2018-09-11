package Practice_3_12

import "reflect"


func IsDisarrange(s1, s2 string) bool {
	runes1, runes2 := []rune(s1), []rune(s2)
	counts1 := make(map[rune]int)
	counts2 := make(map[rune]int)
	if len(runes1) != len(runes2) {
		return false
	}
	for i := 0; i < len(runes1); i++ {
		counts1[runes1[i]]++
		counts2[runes2[i]]++
	}
	if reflect.DeepEqual(counts2, counts1) {
		return true
	} else {
		return false
	}
}
