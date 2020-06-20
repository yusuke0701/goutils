package firebase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type loginAPIReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// VerifyEmailCredential verify email and password at the Firebase Auth.
// ref: https://firebase.google.com/docs/reference/rest/auth#section-sign-in-email-password
func VerifyEmailCredential(email, password string) (bool, error) {
	body, err := json.Marshal(&loginAPIReq{Email: email, Password: password})
	if err != nil {
		return false, err
	}

	u, err := url.Parse("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword")
	if err != nil {
		return false, fmt.Errorf("failed to parse url: %v", err)
	}

	q := u.Query()
	q.Set("key", apiKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(body))
	if err != nil {
		return false, fmt.Errorf("failed to create requset: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := new(http.Client).Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to do request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	resb, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response: %v", err)
	}
	fmt.Printf("failed to login: email = %s, res = %#v\n", email, string(resb))
	return false, nil
}
