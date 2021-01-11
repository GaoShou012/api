package countdown

type OnTimeout func(counter uint64, args ...interface{})
