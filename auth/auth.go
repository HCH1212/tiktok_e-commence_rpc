package auth

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type AuthImpl struct{}

var accessTokenKey = []byte(viper.GetString("jwt.accessTokenKey"))
var refreshTokenKey = []byte(viper.GetString("jwt.refreshTokenKey"))

type Claims struct {
	UserID uint64
	jwt.RegisteredClaims
}

func (i *AuthImpl) GetToken(ctx context.Context, req *auth.UserId) (resp *auth.TwoToken, err error) {
	// accessToken过期时间一周, refreshToken过期时间一月
	accessTokenTime := time.Now().Add(7 * 24 * time.Hour)
	refreshTokenTime := time.Now().Add(4 * 7 * 24 * time.Hour)

	accessClaims := Claims{
		UserID: req.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tiktok_e-commence",
			Subject:   "token",
		},
	}
	refreshClaims := Claims{
		UserID: req.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tiktok_e-commence",
			Subject:   "token",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenStr, err := accessToken.SignedString(accessTokenKey)
	if err != nil {
		return nil, err
	}
	refreshTokenStr, err := refreshToken.SignedString(refreshTokenKey)
	if err != nil {
		return nil, err
	}
	resp = &auth.TwoToken{AccessToken: accessTokenStr, RefreshToken: refreshTokenStr}
	return
}

func (i *AuthImpl) ParseAccessToken(ctx context.Context, req *auth.AccessToken) (resp *auth.UserId, err error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return accessTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		resp = &auth.UserId{Id: claims.UserID}
		return
	}
	return nil, errors.New("invalid token")
}

func (i *AuthImpl) ParseRefreshToken(ctx context.Context, req *auth.RefreshToken) (resp *auth.UserId, err error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		resp = &auth.UserId{Id: claims.UserID}
		return
	}
	return nil, errors.New("invalid token")
}

func (i *AuthImpl) VerifyToken(ctx context.Context, req *auth.AccessToken) (resp *auth.UserId, err error) {
	res, err := i.ParseAccessToken(ctx, req)
	if err != nil {
		return nil, err
	}

	_, err = utils.ById(res.Id)
	// 用户已经不存在
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *AuthImpl) ExecRefreshToken(ctx context.Context, req *auth.RefreshToken) (resp *auth.TwoToken, err error) {
	res, err := i.ParseRefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}
	res2, err := i.GetToken(ctx, res)
	if err != nil {
		return nil, err
	}
	resp = &auth.TwoToken{AccessToken: res2.AccessToken, RefreshToken: res2.RefreshToken}
	return
}
