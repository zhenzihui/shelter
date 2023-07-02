package response

import "projectsuika.com/shelter/src/model/user"

type LoginResp struct {
	User         *user.User `json:"user"`
	Token        string     `json:"token"`
	RefreshToken string     `json:"refreshToken"`
}
