package sms

import (
	"encoding/json"
	aliyunsmsClient "github.com/KenmyZhang/aliyun-communicate"
	"gohub/pkg/logger"
)

// Aliyun 实现 sms.Driver interface
type Aliyun struct{}

// Send 实现 sms.Driver interface 的send方法
func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	smsClient := aliyunsmsClient.New("http://dysmsapi.aliyuncs.com/")

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	logger.DebugJSON("短信[阿里云]", "配置信息", config)

	result, err := smsClient.Execute(
		config["access_key_id"],
		config["access_key_secret"],
		phone,
		config["sign_name"],
		message.Template,
		string(templateParam),
	)

	logger.DebugJSON("短信[阿里云]", "请求内容", smsClient.Request)
	logger.DebugJSON("短信[阿里云]", "请求响应", result)

	if err != nil {
		logger.ErrorString("短信[阿里云]", "发送失败", err.Error())
		return false
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析响应 JSON 错误", err.Error())
		return false
	}

	if result.IsSuccessful() {
		logger.DebugString("短信[阿里云]", "发送成功", "")
		return true
	} else {
		logger.ErrorString("短信[阿里云]", "发送失败", string(resultJSON))
		return false
	}
}
