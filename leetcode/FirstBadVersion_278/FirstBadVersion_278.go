package FirstBadVersion_278

func firstBadVersion(n int, bad int) int {
	l, r := 1, n
	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid, bad) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(version int, bad int) bool {
	return bad != version
}
