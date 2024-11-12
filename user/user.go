package user

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/HCH1212/tiktok_e-commence_rpc/utils"
	"gorm.io/gorm"
)

type UserImpl struct{}

func (i *UserImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 在rpc客户端判断参数合法性，传入参数则已合法
	_, err = utils.ByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) { // 用户不存在
		hashPassword := ""
		hashPassword, err = utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		u := model.User{
			Email:    req.Email,
			Password: hashPassword,
		}
		dao.DB.Create(&u)
		// TODO 直接赋值会报空指针错误
		//resp.Id = uint64(u.ID)
		resp = &user.RegisterResp{Id: uint64(u.ID)}
		return
	} else if err != nil {
		return nil, err
	} else {
		return nil, errors.New("用户已存在")
	}
}

func (i *UserImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// 在rpc客户端判断参数合法性，传入参数则已合法
	u, err := utils.ByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) { // 用户不存在
		return nil, errors.New("用户不存在")
	} else if err != nil {
		return nil, err
	}
	if !utils.VerifyPassword(req.Password, u.Password) {
		return nil, errors.New("密码错误")
	}
	resp = &user.LoginResp{Id: uint64(u.ID)}
	return
	// 让rpc调用Login的时候再去调用一遍GetToken来生成双Token
}
