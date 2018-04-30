package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"unsafe"
)

type SMSData struct {
	Account  string
	Password string
	Phone    string
	Message  string
	Report   bool
}

type Sendable interface {
	Send() error
}

func NewSMSAccount() SMSData {
	return SMSData{
		Account:  "N3404255",
		Password: "ReshIl6Umad0ef",
		// Message:  url.QueryEscape("【253云通讯】您好，您的验证码是999999"),
		Report: true,
	}
}

func (config SMSData) Send(phone, code string) (*string, error) {
	// sms.Account = "N9718791"
	// sms.Password = "Gzcl888888"
	// sms.Message = url.QueryEscape("【253云通讯】您好，您的验证码是999999")
	// sms.Report = true
	params := make(map[string]interface{})
	params["account"] = config.Account
	params["password"] = config.Password
	params["report"] = config.Report
	params["phone"] = phone
	params["msg"] = smsCaptchaTemplate(code)

	fmt.Println(params)

	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	reader := bytes.NewReader(bytesData)
	url := "http://smssh1.253.com/msg/send/json" //请求地址请参考253云通讯自助通平台查看或者询问您的商务负责人获取
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return str, err
}

func GenCaptcha() string {
	// rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", RangeIn(100000, 999999))
	// return RangeIn(100000, 999999)
}

func RangeIn(low, hi int) int {
	rand.Seed(time.Now().UnixNano())
	return low + rand.Intn(hi-low)
}

func smsCaptchaTemplate(code string) string {
	msg := fmt.Sprintf("【源芯】尊敬的用户：您的校验码：%s，工作人员不会索取，请勿泄漏。", code)
	return url.QueryEscape(msg)
}
