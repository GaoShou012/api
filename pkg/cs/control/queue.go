package control

import (
	"cs/env"
	"fmt"
)

func QueueName(code string) string {
	return fmt.Sprintf("cs:customer:queue:%s", code)
}

/*
	加入排队，并且返回排队的位置
*/
func JoinQueue(code string, sessionId string) (uint64, error) {
	key := QueueName(code)
	return env.Queue.Join(key, sessionId)
}
