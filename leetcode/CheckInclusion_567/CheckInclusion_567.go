package CheckInclusion_567

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return true
	}
	var cs1, cs2 [26]int
	for i, c := range s1 {
		cs1[c-'a']++
		cs2[s2[i]-'a']++
	}
	if cs1 == cs2 {
		return true
	}
	for i := m; i < n; i++ {
		cs2[s2[i]-'a']++
		cs2[s2[i-m]-'a']--
		if cs1 == cs2 {
			return true
		}
	}
	return false
}
