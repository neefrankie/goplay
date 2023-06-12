package search

func NaiveStringMatcher(search string, pattern string) int {
	n := len(search)
	m := len(pattern)
	for s := 0; s < n-m; s++ {
		j := 0
		for ; j < m && search[s+j] == pattern[j]; j++ {

		}

		if j >= m {
			return s
		}
	}

	return -1
}
