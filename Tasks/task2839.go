func swap(s string, i, j int) string {
	b := []byte(s)
	b[i], b[j] = b[j], b[i]

	return string(b)
}

func canBeEqual(s1 string, s2 string) bool {
	return s1 == s2 || swap(s1, 0, 2) == s2 || swap(s1, 1, 3) == s2 || swap(swap(s1, 0, 2), 1, 3) == s2
}
