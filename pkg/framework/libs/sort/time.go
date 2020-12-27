package lib_sort

import "time"

type TimeItem struct {
	T time.Time
	V interface{}
}

type TimeItems []*TimeItem

func (s TimeItems) Len() int {
	return len(s)
}
func (s TimeItems) Less(i, j int) bool {
	return s[i].T.Before(s[j].T)
}
func (s TimeItems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

