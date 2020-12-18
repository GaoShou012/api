package queue

import "cs/meta"

/*
	支持多租户模式
	name 队列的名字，一个租户一个队列名字
*/
type Queue interface {
	Init() error

	/*
		设置队列最大长度
	*/
	SetLenMax(name string, num uint64) error
	GetLenMax(name string) (uint64, error)

	/*
		获取队列首位会话信息
	*/
	GetTheFirstSession(name string) (*Item, error)

	/*
		加入队列
	*/
	Join(name string, sessionId string) (uint64, error)

	/*
		离开队列
		如果不在队列中，离开时报错
	*/
	Leave(name string, sessionId string) error

	/*
		获取队列整个长度
	*/
	Len(name string) (uint64, error)

	/*
		获取会话在队列中的位置
		随着会话离开队列，位置一直发生变化
	*/
	GetOffset(name string, session meta.Session) (uint64, error)
}
