package bo

// CreateUserWithAccountBO 创建用户的业务对象
type CreateUserWithAccountBO struct {
	UserId   string `json:"user_id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
