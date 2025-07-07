package auth_responses

type LoginResponseData struct {
	Token string `json:"token"`
}

func NewLoginResponseData(token string) *LoginResponseData {
	return &LoginResponseData{
		Token: token,
	}
}
