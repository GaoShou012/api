package lib_sort

type IntItem struct {
	I int
	V interface{}
}

type IntItems []*IntItem

func (s IntItems) Len() int {
	return len(s)
}
func (s IntItems) Less(i, j int) bool {
	return s[i].I < s[j].I
}
func (s IntItems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
