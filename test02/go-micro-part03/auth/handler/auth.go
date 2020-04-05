package handler

import (
	"context"
	"go-micro/go-micro-part03/auth/model/access"
	auth "go-micro/go-micro-part03/proto"
	"go-micro/go-micro-part03/proto/common"
	"strconv"

	log "github.com/micro/go-micro/v2/logger"
)

type AuthHandler struct{}

// 生成Token
func (a *AuthHandler) CreateToken(ctx context.Context, requ *auth.TokenRequest, resp *auth.TokenResponse) error {

	tokenService := access.GetTokenService()
	token, err := tokenService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatUint(requ.UserId, 10),
		Name: requ.UserName,
	})

	resp.BaseResponse = &common.BaseResponse{
		IsSuccess:    true,
		ErrorMessage: "测试成功",
	}

	if err != nil {
		return err
	}
	resp.BaseResponse.IsSuccess = true
	resp.Token = token

	return nil

}

// 删除Token
func (a *AuthHandler) DeleteToken(ctx context.Context,requ *auth.TokenRequest, resp *auth.TokenResponse) error {

	log.Info("DeleteToken", requ.Token)

	tokenService := access.GetTokenService()
	err := tokenService.DelUserAccessToken(requ.Token)
	if err != nil {
		return err
	}

	return nil
}


// 获取Token
func (a *AuthHandler) GetToken(ctx context.Context, requ *auth.TokenRequest, resp *auth.TokenResponse) error  {

	tokenService := access.GetTokenService()

	log.Infof("[GetCachedAccessToken] 获取缓存的token，%d", requ.UserId)
	token, err := tokenService.GetCachedAccessToken(&access.Subject{
		ID: strconv.FormatUint(requ.UserId, 10),
	})
	if err != nil {

		log.Errorf("[GetCachedAccessToken] 获取缓存的token失败，err：%s", err)
		return err
	}
	resp.Token = token
	return nil

}