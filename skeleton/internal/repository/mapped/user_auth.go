package mapped

import "time"

const (
	AccountUserAuthType UserAuthType = "account" // 账号类型
	MobilePhoneAuthType UserAuthType = "mobile"  // 手机号类型
	EmailAuthType       UserAuthType = "email"   // 邮箱类型
)

type UserAuthType = string

type UserAuthMapped struct {
	Id        uint64    // 自增主键
	UserId    string    // 用户 ID
	AuthType  string    // 授权类型
	Account   string    // 授权账号，通常为邮箱、手机号、第三方账号等
	AuthData  string    // 授权数据，如加密后的密码、第三方授权等
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
