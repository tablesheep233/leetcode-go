package Tribonacci_1137

func Tribonacci(n int) int {
	if n < 3 {
		if n > 0 {
			return 1
		}
		return 0
	}
	t0 := 0
	t1 := 1
	t2 := 1
	for i := 2; i < n; i++ {
		t := t0 + t1 + t2
		t0 = t1
		t1 = t2
		t2 = t
	}
	return t2
}
