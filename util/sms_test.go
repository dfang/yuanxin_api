package util

import (
	"testing"
)

// GOCACHE=off go test -v *.go
func TestSendSms(t *testing.T) {
	code := GenCaptcha()
	// c := fmt.Sprintf("%d", rangeIn(100000, 999999))

	t.Log(code)

	// result, err := NewSMSAccount().Send("13530605832", code)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// t.Log(*result)
}
