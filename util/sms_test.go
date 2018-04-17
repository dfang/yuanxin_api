package util

import (
	"testing"
)

func TestSendSms(t *testing.T) {
	sms := SMSData{
		//Account: "N9718791",
		//Password: "Gzcl888888",
		Phone: "17671757383",
		//Message: url.QueryEscape("【253云通讯】您好，您的验证码是999999"),
		//Report: true,
	}
	t.Logf("%#v", sms)

	// result, err := SendSms(&sms)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// t.Log(*result)
}
