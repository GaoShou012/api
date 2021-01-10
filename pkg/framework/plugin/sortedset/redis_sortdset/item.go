package redis_sortdset

type item struct {
	iKey string
	iVal int
}

func (i *item) Key() string {
	return i.iKey
}
func (i *item) Val() int {
	return i.iVal
}
