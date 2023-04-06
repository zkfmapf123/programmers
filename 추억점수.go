package main

func main() {
	solution([]string{"may", "kein", "kain", "radi"}, []int{5, 10, 1, 3}, [][]string{
		{"may", "kein", "kain", "radi"}, {"may", "kein", "brin", "deny"}, {"kon", "kain", "may", "coni"},
	})
}

func solution(name []string, yearning []int, photo [][]string) []int {

	result, m := make([]int, len(photo)), make(map[string]int, len(name))
	for i, n := range name {
		m[n] = yearning[i]
	}

	for i, ps := range photo {
		for _, p := range ps {
			val, exists := m[p]
			if !exists {
				continue
			}

			result[i] += val
		}
	}

	return result
}
