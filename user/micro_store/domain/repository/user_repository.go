package repository

import (
	"github.com/coder279/micro_store/domain/model"
	"github.com/jinzhu/gorm"
	)

type IUserRepository interface {
	InitTable() error
	FindByName(string) (*model.User,error)
	FindById(int64)(*model.User,error)
	CreateUser(user *model.User) error
	DeleteUserById(int64) error
	UpdateUser(*model.User) error
	FindAllUser()([] *model.User,error)
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
//初始化表
func (u *UserRepository)InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}
//查找用户通过名称
func (u *UserRepository) FindByName(name string)(*model.User,error){
	user := &model.User{}
	return user,u.mysqlDb.Where("user_name = ?",name).Find(user).Error
}
//根据用户ID去查找
func (u *UserRepository) FindById(id int64)(*model.User,error){
	user := &model.User{}
	return user,u.mysqlDb.Where("id = ?",id).Error
}
//创建用户信息
func (u *UserRepository) CreateUser(user *model.User)(error){
	return u.mysqlDb.Create(user).Error
}
//删除用户数据
func (u *UserRepository) DeleteUserById(id int64)(error){
	user := &model.User{}
	return u.mysqlDb.Where("id = ?",id).Delete(user).Error
}
//根据id查找用户数据
func (u *UserRepository) UpdateUser(user *model.User) (error){
	return u.mysqlDb.Model(user).Update(&user).Error
}
//查找所有用户数据
func (u *UserRepository) FindAllUser() (userAll [] *model.User,err error) {
	return userAll,u.mysqlDb.Find(&userAll).Error
}




