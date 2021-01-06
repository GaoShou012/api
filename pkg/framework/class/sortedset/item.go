package sortedset

type Item interface {
	Key() string
	Val() int64
}
