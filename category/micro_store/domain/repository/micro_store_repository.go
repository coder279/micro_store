package repository
import (
	"github.com/jinzhu/gorm"
	"github.com/coder279/micro_store/domain/model"
)
type IMicro_storeRepository interface{
    InitTable() error
    FindMicro_storeByID(int64) (*model.Micro_store, error)
	CreateMicro_store(*model.Micro_store) (int64, error)
	DeleteMicro_storeByID(int64) error
	UpdateMicro_store(*model.Micro_store) error
	FindAll()([]model.Micro_store,error)

}
//创建micro_storeRepository
func NewMicro_storeRepository(db *gorm.DB) IMicro_storeRepository  {
	return &Micro_storeRepository{mysqlDb:db}
}

type Micro_storeRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *Micro_storeRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Micro_store{}).Error
}

//根据ID查找Micro_store信息
func (u *Micro_storeRepository)FindMicro_storeByID(micro_storeID int64) (micro_store *model.Micro_store,err error) {
	micro_store = &model.Micro_store{}
	return micro_store, u.mysqlDb.First(micro_store,micro_storeID).Error
}

//创建Micro_store信息
func (u *Micro_storeRepository) CreateMicro_store(micro_store *model.Micro_store) (int64, error) {
	return micro_store.ID, u.mysqlDb.Create(micro_store).Error
}

//根据ID删除Micro_store信息
func (u *Micro_storeRepository) DeleteMicro_storeByID(micro_storeID int64) error {
	return u.mysqlDb.Where("id = ?",micro_storeID).Delete(&model.Micro_store{}).Error
}

//更新Micro_store信息
func (u *Micro_storeRepository) UpdateMicro_store(micro_store *model.Micro_store) error {
	return u.mysqlDb.Model(micro_store).Update(micro_store).Error
}

//获取结果集
func (u *Micro_storeRepository) FindAll()(micro_storeAll []model.Micro_store,err error) {
	return micro_storeAll, u.mysqlDb.Find(&micro_storeAll).Error
}

