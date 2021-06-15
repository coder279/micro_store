package handler

import (
	"context"
	"github.com/coder279/micro_store/domain/model"
	"github.com/coder279/micro_store/domain/service"
	user "github.com/coder279/micro_store/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

//功能函数--类型转化
func UserForResponse(userModel *model.User) *user.UserInfoResponse{
	response := &user.UserInfoResponse{}
	response.FirstName = userModel.FirstName
	response.UserName = userModel.UserName
	return response
}

//注册
func (u *User) Register (ctx context.Context,userRegisterRequest *user.UserRegisterRequest,userResponse *user.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:userRegisterRequest.UserName,
		FirstName:userRegisterRequest.FirstName,
		HashPassword:userRegisterRequest.Password,
	}
	_,err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	userResponse.Message = "注册成功"
	return nil
}

//登录
func (u *User)Login(ctx context.Context,userLoginRequest *user.UserLoginRequest,userLoginResponse *user.UserLoginResponse) error {
	isOK,err := u.UserDataService.CheckPwd(userLoginRequest.UserName,userLoginRequest.Password)
	if err != nil {
		return err
	}
	if isOK {
		userLoginResponse.Message = "登录成功"
	}else{
		userLoginResponse.Message = "登录失败"
	}
	return nil
}


func (u *User) GetUserInfo(ctx context.Context, userInfoRequest *user.UserInfoRequest,userInfoResponse *user.UserInfoResponse) error {
	user,err := u.UserDataService.FindUserByName(userInfoRequest.UserName)
	if err != nil {
		return err
	}
	userInfoResponse = UserForResponse(user)
	return nil
}