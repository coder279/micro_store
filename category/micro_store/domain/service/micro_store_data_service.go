package service

import (
	"github.com/coder279/micro_store/domain/model"
	"github.com/coder279/micro_store/domain/repository"
)

type IMicro_storeDataService interface {
	AddMicro_store(*model.Micro_store) (int64 , error)
	DeleteMicro_store(int64) error
	UpdateMicro_store(*model.Micro_store) error
	FindMicro_storeByID(int64) (*model.Micro_store, error)
	FindAllMicro_store() ([]model.Micro_store, error)
}


//创建
func NewMicro_storeDataService(micro_storeRepository repository.IMicro_storeRepository) IMicro_storeDataService{
	return &Micro_storeDataService{ micro_storeRepository }
}

type Micro_storeDataService struct {
	Micro_storeRepository repository.IMicro_storeRepository
}


//插入
func (u *Micro_storeDataService) AddMicro_store(micro_store *model.Micro_store) (int64 ,error) {
	 return u.Micro_storeRepository.CreateMicro_store(micro_store)
}

//删除
func (u *Micro_storeDataService) DeleteMicro_store(micro_storeID int64) error {
	return u.Micro_storeRepository.DeleteMicro_storeByID(micro_storeID)
}

//更新
func (u *Micro_storeDataService) UpdateMicro_store(micro_store *model.Micro_store) error {
	return u.Micro_storeRepository.UpdateMicro_store(micro_store)
}

//查找
func (u *Micro_storeDataService) FindMicro_storeByID(micro_storeID int64) (*model.Micro_store, error) {
	return u.Micro_storeRepository.FindMicro_storeByID(micro_storeID)
}

//查找
func (u *Micro_storeDataService) FindAllMicro_store() ([]model.Micro_store, error) {
	return u.Micro_storeRepository.FindAll()
}

