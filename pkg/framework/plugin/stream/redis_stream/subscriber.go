package stream_redis_stream

type subscriber struct {
	unSubscribe func()
}

func (s *subscriber) setupUnSubscribe(fn func()) {
	s.unSubscribe = fn
}

func (s *subscriber) UnSubscribe() error {
	s.unSubscribe()
	return nil
}
