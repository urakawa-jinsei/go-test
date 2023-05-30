package hsd

func StringDistance(lhs,rhs string) int {
	return Ditance([]rune(lhs),[]rune(rhs))
}

func Ditance(a,b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}

	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}