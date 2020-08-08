package userinfo

// NewMockHelper は、Helperのモックを作成するための関数です。
func NewMockHelper() Helper {
	return &mockHelper{}
}

type mockHelper struct{}

// CallUserInfoMeAPI のモックデータ
var (
	CallUserInfoMeAPIMockData  *CallUserInfoMeAPIRes
	CallUserInfoMeAPIMockError error
)

func (mockHelper) CallUserInfoMeAPI(token string) (*CallUserInfoMeAPIRes, error) {
	if CallUserInfoMeAPIMockData != nil {
		tmp := CallUserInfoMeAPIMockData
		CallUserInfoMeAPIMockData = nil
		return tmp, nil
	}
	if CallUserInfoMeAPIMockError != nil {
		tmp := CallUserInfoMeAPIMockError
		CallUserInfoMeAPIMockError = nil
		return nil, tmp
	}
	return nil, nil
}
