package MapSum_677

import "strings"

type MapSum struct {
	key []string
	val []int
}

func Constructor() MapSum {
	return MapSum{[]string{}, []int{}}
}

func (this *MapSum) Insert(key string, val int) {
	for i, s := range this.key {
		if s == key {
			this.val[i] = val
			return
		}
	}
	this.key = append(this.key, key)
	this.val = append(this.val, val)
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for i, s := range this.key {
		if strings.HasPrefix(s, prefix) {
			sum += this.val[i]
		}
	}
	return sum
}
