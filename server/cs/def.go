package cs

import "fmt"

const (
	// 存储所有的会话
	// 用于异步检查会话状态，进行内存回收
	TopicOfSessionsSet = "SessionsSortedSet"
)

func TopicOfMerchantSessionSet(merchantCode string) string {
	return fmt.Sprintf("SessionSortedSet:%s", merchantCode)
}
