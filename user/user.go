package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/HCH1212/tiktok_e-commence_rpc/utils"
	"gorm.io/gorm"
	"time"
)

type UserImpl struct{}

func (i *UserImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 在rpc客户端判断参数合法性，传入参数则已合法

	// 先看redis中是否有用户
	res, err := dao.RDB.Get(ctx, req.Email).Result()
	if err == nil && res != "" {
		return nil, errors.New("用户已存在")
	}

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

		// 添加到缓存
		userData, _ := json.Marshal(u)
		dao.RDB.Set(ctx, u.Email, userData, 24*time.Hour) // 缓存 24 小时

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

	// 检查缓存中是否存在用户
	var u *model.User
	res, err := dao.RDB.Get(ctx, req.Email).Result()
	if err == nil {
		u = &model.User{}
		_ = json.Unmarshal([]byte(res), u)
	} else {
		// 缓存未命中，从数据库查询
		u, err = utils.ByEmail(req.Email)
		if errors.Is(err, gorm.ErrRecordNotFound) { // 用户不存在
			return nil, errors.New("用户不存在")
		} else if err != nil {
			return nil, err
		}

		// 更新缓存
		userData, _ := json.Marshal(u)
		dao.RDB.Set(ctx, u.Email, userData, 24*time.Hour) // 缓存 24 小时
	}

	if u == nil {
		return nil, errors.New("u nil")
	}
	if !utils.VerifyPassword(req.Password, u.Password) {
		return nil, errors.New("密码错误")
	}
	resp = &user.LoginResp{Id: uint64(u.ID)}
	return
	// 让rpc调用Login的时候再去调用一遍GetToken来生成双Token
}
