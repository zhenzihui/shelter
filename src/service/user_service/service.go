package user_service

import (
	"errors"
	"projectsuika.com/shelter/src/common/jwt_helper"
	"projectsuika.com/shelter/src/common/request"
	"projectsuika.com/shelter/src/common/response"
	"projectsuika.com/shelter/src/model/user"
)

func Login(req request.LoginReq) (*response.LoginResp, error) {
	userInfo := user.Login(req.UserName, req.Password)
	if userInfo == nil {
		return nil, errors.New(response.CodeRegistry.InvalidPassword.Desc)
	}

	token, refreshToken, _ := jwt_helper.GenerateAllTokens(userInfo.Id, userInfo.UserName)
	return &response.LoginResp{
		Token:        token,
		RefreshToken: refreshToken,
		User:         userInfo,
	}, nil
}
