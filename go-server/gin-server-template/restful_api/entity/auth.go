package entity

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInfo struct {
	UserID   int64  `json:"id"`
	Username string `json:"username"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type AuthInfo struct {
	UserID   int64  `json:"id"`
	Username string `json:"username"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
