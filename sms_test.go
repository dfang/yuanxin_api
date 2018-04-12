package main

import (
	"testing"
)

func TestSendSms(t *testing.T) {
	//params := make(map[string]interface{})
	//params["account"] = "N9718791"
	//params["password"] = "Gzcl888888"
	//params["phone"] = "13530605832"
	//params["msg"] = url.QueryEscape("【253云通讯】您好，您的验证码是999999")
	//params["report"] = "false"

	sms := SMSData{
		//Account: "N9718791",
		//Password: "Gzcl888888",
		Phone: "17671757383",
		//Message: url.QueryEscape("【253云通讯】您好，您的验证码是999999"),
		//Report: true,
	}

	SendSms(&sms)
}