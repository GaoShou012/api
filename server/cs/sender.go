package cs

type Sender interface {
	GetUserId() uint64
	GetUserType() uint64
	GetUsername() string
	GetNickname() string
	GetThumb() string
}

func GetSenderInfoFrom(sender Sender) map[string]interface{} {
	m := make(map[string]interface{})
	m["UserId"] = sender.GetUserType()
	m["UserType"] = sender.GetUserType()
	m["Username"] = sender.GetUsername()
	m["Nickname"] = sender.GetNickname()
	m["Thumb"] = sender.GetThumb()
	return m
}
