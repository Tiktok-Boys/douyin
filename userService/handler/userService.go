package handler

import (
	"context"
	"go-micro.dev/v4/errors"
	proto "userService/proto"
)

type UserService struct {
	auth   *Authentication
	mapper *Mapper
}

func NewUserService() *UserService {
	return &UserService{
		auth:   NewAuthentication(),
		mapper: NewMapper(),
	}
}

func (u *UserService) Register(ctx context.Context, req *proto.DouyinUserRegisterRequest, rsp *proto.DouyinUserRegisterResponse) error {
	username := req.Username
	password := req.Password
	userId, err := u.mapper.getIdByUsername(ctx, username)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "获取用户名使用情况失败"
		return err
	}
	if userId != -1 {
		rsp.StatusCode = 1
		rsp.StatusMsg = "用户名已被注册!"
		return errors.InternalServerError("500", "Username already registered")
	}
	err = u.mapper.createUser(ctx, username, password)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "注册失败"
		return err
	}
	userId, err = u.mapper.getIdByUsername(ctx, username)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "获取用户名使用情况失败"
		return err
	}
	var token string
	token, err = u.auth.GenerateToken(ctx, username, password, userId)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "token获取失败"
		return err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "注册成功"
	rsp.Token = token
	rsp.UserId = userId
	return nil
}

func (u *UserService) Login(ctx context.Context, req *proto.DouyinUserLoginRequest, rsp *proto.DouyinUserLoginResponse) error {
	username := req.Username
	password := req.Password
	isVerified, err := u.mapper.isPasswordCorrect(ctx, username, password)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "登录时发生错误"
		return err
	}
	if !isVerified {
		rsp.StatusCode = 1
		rsp.StatusMsg = "用户名或密码错误"
		return errors.InternalServerError("500", "Incorrect identity")
	}
	var userId int64
	userId, err = u.mapper.getIdByUsername(ctx, username)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "获取用户名使用情况失败"
		return err
	}
	var token string
	token, err = u.auth.GenerateToken(ctx, username, password, userId)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "token获取失败"
		return err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "登录成功"
	rsp.Token = token
	rsp.UserId = userId
	return nil
}

func (u *UserService) UserInfo(ctx context.Context, req *proto.DouyinUserRequest, rsp *proto.DouyinUserResponse) error {
	token := req.Token
	targetUserID := req.UserId
	selfUserId, err := u.auth.ValidateToken(ctx, token)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "鉴权失败"
		return err
	}
	var user *proto.User
	user, err = u.mapper.getUserInfo(ctx, targetUserID, selfUserId, token)
	if err != nil {
		rsp.StatusCode = -1
		rsp.StatusMsg = "获取用户信息失败"
		return err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "获取成功"
	rsp.User = user
	return nil
}
