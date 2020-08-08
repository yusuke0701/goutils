package userinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Helper は、APIを呼び出すためのインタフェースです。
type Helper interface {
	CallUserInfoMeAPI(token string) (*CallUserInfoMeAPIRes, error)
}

// NewHelper は、Helperを作成するための関数です。
func NewHelper(c *http.Client) Helper {
	return &helper{httpClient: c}
}

type helper struct {
	httpClient *http.Client
}

// CallUserInfoMeAPIRes は、UserInfoMeAPIのjson形式でのレスポンス構造体を表す。
type CallUserInfoMeAPIRes struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

// CallUserInfoMeAPI は、自分のユーザー情報を取得するAPIを呼び出す。
// ref: https://any-api.com/googleapis_com/oauth2/console/userinfo/oauth2_userinfo_v2_me_get
func (h *helper) CallUserInfoMeAPI(token string) (*CallUserInfoMeAPIRes, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/userinfo/v2/me", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alt", "json")

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := h.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to call user info api: %s", string(b))
	}

	apiRes := new(CallUserInfoMeAPIRes)
	if err := json.Unmarshal(b, apiRes); err != nil {
		return nil, err
	}
	return apiRes, nil
}
