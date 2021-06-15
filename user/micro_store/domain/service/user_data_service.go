package service

import (
	"errors"
	"github.com/coder279/micro_store/domain/model"
	"github.com/coder279/micro_store/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User)(int64,error)
	DeleteUser(int64)error
	UpdateUser(user *model.User,isChangePwd bool)(err error)
	FindUserByName(string)(*model.User,error)
	CheckPwd(userName string,pwd string)(isOK bool,err error)
}


//实例化工厂里面的函数
func NewUserDataService(userRepository repository.UserRepository)IUserDataService{
	return &UserDataService{UserRepository:userRepository}
}
type UserDataService struct {
	UserRepository repository.UserRepository
}
func GeneratePassword(userPassword string)([]byte,error){
	return bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string,hashed string)(isOk bool,err error){
	if err = bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(userPassword));err != nil {
		return false,errors.New("密码不正确")
	}else{
		return true,nil
	}
}

//插入用户信息
func (u *UserDataService) AddUser(user *model.User)(int64,error){
	pwdByte,err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID,err
	}
	user.HashPassword = string(pwdByte)
	return user.ID,u.UserRepository.CreateUser(user)
}
//删除用户
func (u *UserDataService) DeleteUser(userId int64)(err error){
	return u.UserRepository.DeleteUserById(userId)
}

//修改用户
func (u *UserDataService) UpdateUser(user *model.User,isChangePwd bool)(err error){
	if isChangePwd{
		password,err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(password)
	}
	return u.UserRepository.UpdateUser(user)

}

//根据名字查找数据
func (u *UserDataService) FindUserByName(userName string) (user *model.User,err error){
	return u.UserRepository.FindByName(userName)
}

func (u *UserDataService)CheckPwd(userName string,pwd string) (isOK bool,err error){
	user,err := u.UserRepository.FindByName(userName)
	if err != nil {
		return false,err
	}
	return ValidatePassword(pwd,user.HashPassword)
}




