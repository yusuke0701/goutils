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
	RemoveCache()
}

// NewHelper は、Helperを作成するための関数です。
// enableCache を true にすると、APIの結果をローカルのキャッシュにします。
func NewHelper(c *http.Client, enableCache bool) Helper {
	if enableCache && _userInfoMeAPIResCache == nil {
		_userInfoMeAPIResCache = make(map[string]*CallUserInfoMeAPIRes, 100)
	}
	return &helper{httpClient: c, enableCache: enableCache}
}

type helper struct {
	httpClient  *http.Client
	enableCache bool
}

var _userInfoMeAPIResCache map[string]*CallUserInfoMeAPIRes

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
	if h.enableCache {
		value, ok := _userInfoMeAPIResCache[token]
		if ok {
			return value, nil
		}
	}
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
	if h.enableCache {
		_userInfoMeAPIResCache[token] = apiRes
	}
	return apiRes, nil
}

// RemoveCache は、ローカルに保存したキャッシュを削除する
func (helper) RemoveCache() {
	_userInfoMeAPIResCache = make(map[string]*CallUserInfoMeAPIRes, 100)
}
