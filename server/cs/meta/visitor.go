package meta

// 商户对访客信息进行加密时的数据结构
type Visitor struct {
	UserId   uint64
	Username string
	Nickname string
	Thumb    string
}
