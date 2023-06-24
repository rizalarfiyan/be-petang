package constant

import "time"

const (
	AuthKeyLength = 128
	AuthExpire    = 10 * time.Minute

	RedisKeyAuth = "auth:"

	TemplateChangePassword = "template/change-password.html"
)
