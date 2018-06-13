package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//AuthTokenResponse 环信获取token的response
type AuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Application string `json:"application"`
}

// type UserResponse struct {
// 	Action      string `json:"action"`
//   Application string `json:"application"`
//   Path string `json:"path"`
//   URI string `json:"uri"`

// }

// AuthTokenEndpoint 环信获取token
func AuthTokenEndpoint() string {

	orgName := "1155180513228554"
	appName := "origincore"

	url := fmt.Sprintf("http://a1.easemob.com/%s/%s/token", orgName, appName)

	// clientID := "YXA6ZVuWgFbEEeiuvY_QCzohVw"
	// clientSecret := "YXA6rUFs_LaAd1xl9uf0iAq4gnLfKCY"

	var jsonStr = []byte(`{
    "grant_type": "client_credentials",
    "client_id": "YXA6ZVuWgFbEEeiuvY_QCzohVw",
    "client_secret": "YXA6rUFs_LaAd1xl9uf0iAq4gnLfKCY"
    }`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	var s = new(AuthTokenResponse)
	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	// fmt.Println("token: ", s.AccessToken)
	// return string(body)
	return s.AccessToken

}

func RegisterEasemobUser() {
	orgName := "1155180513228554"
	appName := "origincore"

	url := fmt.Sprintf("http://a1.easemob.com/%s/%s/users", orgName, appName)
	var jsonStr = []byte(`[
    {
      "username": "string",
      "password": "string"
    }
  ]`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var i interface{}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &i)

	// fmt.Println("%v ", i)

}
