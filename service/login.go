package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TID     string `json:"tid"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

func Login() string {

	url := "https://gwapi.fleet.vip/gw/fms/auth/login"
	method := "POST"

	payload := strings.NewReader(`{
    "mobile": "",
    "password": ""
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.79")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var response map[string]interface{}
	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	data := response["data"].(map[string]interface{})
	fmsSessionID := data["fms_session_id"].(string)
	return fmsSessionID
}

func GetToken(session string) string {
	url := "https://gwapi.fleet.vip/gw/fms/msg/token"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return ""
	}
	req.Header.Add("Fms-Session-Id", session)
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return ""
	}

	if tokenResponse.Code != 1 {
		return ""
	}

	return tokenResponse.Data.Token
}
