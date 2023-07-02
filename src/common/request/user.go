package request

type LoginReq struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}
