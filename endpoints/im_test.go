package endpoints

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Checksum(appSecret, nonce, curTime string) string {

	h := sha1.New()

	// appKey := ""
	// appSecret := ""
	// nonce := rand.Int31()
	// curTime := time.Now().Unix()

	fmt.Println(appSecret)
	fmt.Println(nonce)
	fmt.Println(curTime)

	h.Write([]byte(appSecret + nonce + curTime))
	checkSum := h.Sum(nil)
	// str := bytes.NewBuffer(checkSum).String()
	str := fmt.Sprintf("%x", checkSum)
	fmt.Println(str)
	fmt.Printf("%x\n", checkSum)
	return str
}

// func Test_CheckSum(t *testing.T) {
// 	rand.Seed(time.Now().UTC().UnixNano())

// 	// appKey := "d45545b3eeb821970eab26931859871e"
// 	appSecret := "d31182026a36"
// 	nonce := strconv.FormatInt(rand.Int63(), 10)
// 	curTime := strconv.FormatInt(time.Now().Unix(), 10)

// 	fmt.Println(appSecret)
// 	fmt.Println(nonce)
// 	fmt.Println(curTime)

// 	sum := Checksum(appSecret, nonce, curTime)
// 	fmt.Println(sum)
// }

func Test_RegisterImUser(t *testing.T) {

	appKey := "d45545b3eeb821970eab26931859871e"
	appSecret := "d31182026a36"
	nonce := strconv.FormatInt(rand.Int63(), 10)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	hash := Checksum(appSecret, nonce, curTime)

	client := &http.Client{}

	params := url.Values{}
	params.Set("accid", "helloworld")
	params.Set("name", "zhangsan")
	postData := strings.NewReader(params.Encode())

	req, _ := http.NewRequest("POST", "https://api.netease.im/nimserver/user/create.action", postData)

	req.Header.Add("AppKey", appKey)
	req.Header.Add("Nonce", nonce)
	req.Header.Add("CurTime", curTime)
	req.Header.Add("CheckSum", hash)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//print raw response body for debugging purposes

	fmt.Println(string(body))
}
