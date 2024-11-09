package auth

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
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

func (i *AuthImpl) GetToken(ctx context.Context, req *authservice.GetTokenArgs) (*authservice.GetTokenResult, error) {
	// accessToken过期时间一周, refreshToken过期时间一月
	accessTokenTime := time.Now().Add(7 * 24 * time.Hour)
	refreshTokenTime := time.Now().Add(4 * 7 * 24 * time.Hour)

	accessClaims := Claims{
		UserID: req.Req.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tiktok_e-commence",
			Subject:   "token",
		},
	}
	refreshClaims := Claims{
		UserID: req.Req.Id,
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
	var result = &auth.TwoToken{AccessToken: accessTokenStr, RefreshToken: refreshTokenStr}
	return &authservice.GetTokenResult{result}, nil
}

func (i *AuthImpl) ParseAccessToken(ctx context.Context, req *authservice.ParseAccessTokenArgs) (*authservice.ParseAccessTokenResult, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.Req.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return accessTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		var result = &auth.UserId{Id: claims.UserID}
		return &authservice.ParseAccessTokenResult{result}, nil
	}
	return nil, errors.New("invalid token")
}

func (i *AuthImpl) ParseRefreshToken(ctx context.Context, req *authservice.ParseRefreshTokenArgs) (*authservice.ParseRefreshTokenResult, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.Req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 有效
	if token.Valid {
		var result = &auth.UserId{Id: claims.UserID}
		return &authservice.ParseRefreshTokenResult{result}, nil
	}
	return nil, errors.New("invalid token")
}

func (i *AuthImpl) ExecRefreshToken(ctx context.Context, req *authservice.ExecRefreshTokenArgs) (*authservice.ExecRefreshTokenResult, error) {
	ParseRefreshTokenReq := &auth.RefreshToken{RefreshToken: req.Req.RefreshToken}
	res, err := i.ParseRefreshToken(ctx, &authservice.ParseRefreshTokenArgs{ParseRefreshTokenReq})
	if err != nil {
		return nil, err
	}
	id := res.Success.Id
	GetTokenReq := &auth.UserId{Id: id}
	res2, err := i.GetToken(ctx, &authservice.GetTokenArgs{GetTokenReq})
	if err != nil {
		return nil, err
	}
	result := &auth.TwoToken{AccessToken: res2.Success.AccessToken, RefreshToken: res2.Success.RefreshToken}
	return &authservice.ExecRefreshTokenResult{result}, nil
}
