package auth

type SigninRequest struct {
	Username string `json:"nik"`
	Password string `json:"password"`
}
