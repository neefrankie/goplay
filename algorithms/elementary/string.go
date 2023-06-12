package elementary

// StringSearch discovers all occurrences of a word in a text string.
func StringSearch(src, target string) []int {
	l1 := len(src)
	l2 := len(target)
	pos := make([]int, 0)
	for i := 0; i < l1; i++ {
		var j int
		for j = 0; j < l2; j++ {
			if src[i+j] != target[j] {
				break
			}
		}
		if j == l2 {
			pos = append(pos, i)
		}
	}

	return pos
}
