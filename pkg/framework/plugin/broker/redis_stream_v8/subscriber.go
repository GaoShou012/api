package broker_redis_stream_v8

type subscriber struct {
	unSubscribe func()
}

func (s *subscriber) SetupUnSubscribe(fn func()) {
	s.unSubscribe = fn
}

func (s *subscriber) UnSubscribe() error {
	s.unSubscribe()
	return nil
}


