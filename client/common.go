package client

// UserInfo 用户信息
type UserInfo struct {
	ID       int64  `json:"ID"`
	Name     string `json:"Name"`
	NickName string `json:"NickName"`
	Avatar   []byte `json:"Avatar"`
}

// UserList 用户列表
type UserList struct {
	Users []*UserInfo `json:"Users"`
}

//GroupInfo 群组信息
type GroupInfo struct {
	ID    int64   `json:"ID"`
	Name  string  `json:"Name"`
	Users []int64 `json:"Users"`
}

// GroupList 群组列表
type GroupList struct {
	Groups []*GroupInfo `json:"Groups"`
}

//FriendList 好友列表
type FriendList struct {
	Users []int64 `json:"Users"`
}
