package model

type UserInfo struct {
	ID       int64
	Username string `json:"username"`
	Role     string `json:"role"`
}
