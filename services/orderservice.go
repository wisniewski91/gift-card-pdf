package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/wisniewski91/gift-card-pdf/types"
)

type autehObj struct {
	grant_type    string
	client_id     string
	client_secret string
}

func autorize() (types.PayUAuth, error) {

	apiUrl := "https://secure.snd.payu.com"
	resource := "/pl/standard/user/oauth/authorize"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", "")
	data.Set("client_secret", "")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	var authObj types.PayUAuth
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("Error making request:", err)
		return authObj, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&authObj)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return authObj, err
	}

	return authObj, nil

}

func Order() {
	autorize()
}
