package model

import (
	"fmt"
	"testing"

	taobao "github.com/smartwalle/taobao"
)

// df
var APPKEY = "24929444"
var APPSECRET = "75042ef08f6b1b6fd00fb89fa98dc89f"

// var APPKEY = "1024929444"
// var APPSECRET = "sandbox05b816d34a7be68d7b8642681"

// clb
// var APPKEY = "24927822"
// var APPSECRET = "3b18b0ffafd79aacf203441a54e68992"

// worked
// var APPKEY = "23201003"
// var APPSECRET = "1e2dfd16981d75142597fd10131b17b5"

func TestUser_RegisterIMUser(t *testing.T) {

	var u1 = taobao.OpenIMUserInfo{}
	u1.UserId = "admin12345"
	u1.Password = "a6facfa821ba92c3c17f4c3fae5442c2"
	u1.Nick = "我是管理员"

	var u2 = taobao.OpenIMUserInfo{}
	u2.UserId = "test12345"
	u2.Password = "123456"

	// taobao.OpenIMAddUsersParam{}
	var p = taobao.OpenIMAddUsersParam{}
	p.AddOpenIMUser(&u1)
	p.AddOpenIMUser(&u2)

	// var APPKEY = "24929444"
	// var APPSECRET = "f39c810bc8cf801151520242a51bfef2"

	fmt.Println(taobao.RequestWithKey(APPKEY, APPSECRET, p))
}

func TestUser_ListIMUser(t *testing.T) {
	var p = taobao.OpenIMGetUsersParam{}
	p.UserIds = []string{"admin12345", "test12345"}
	fmt.Println("===== OpenIMGetUserParam =====")

	// var APPKEY = "24929444"
	// var APPSECRET = "f39c810bc8cf801151520242a51bfef2"
	// taobao.UpdateKey(APPKEY, APPSECRET)
	// fmt.Println(taobao.Request(p))
	fmt.Println(taobao.RequestWithKey(APPKEY, APPSECRET, p))
}
