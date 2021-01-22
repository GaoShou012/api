package broker_redis_pubsub_v7_cluster

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


