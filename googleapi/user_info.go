package googleapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UserInfoAPIRes は、UserInfoMeAPIのjson形式でのレスポンス構造体を表す。
type UserInfoAPIRes struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

// CallUserInfoMeAPI は、自分のユーザー情報を取得するAPIを呼び出す
// ref: https://any-api.com/googleapis_com/oauth2/console/userinfo/oauth2_userinfo_v2_me_get
func CallUserInfoMeAPI(httpClient *http.Client, token string) (*UserInfoAPIRes, error) {
	req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/userinfo/v2/me", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("alt", "json")

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := httpClient.Do(req)
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

	userInfoAPIRes := new(UserInfoAPIRes)
	if err := json.Unmarshal(b, userInfoAPIRes); err != nil {
		return nil, err
	}
	return userInfoAPIRes, nil
}
