package mail

type Driver struct {
	// 检查验证码
	Send(email Email, config map[string]string) bool
}
