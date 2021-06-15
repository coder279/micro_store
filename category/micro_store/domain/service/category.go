package service

import (
	"github.com/coder279/category/domain/model"
	"github.com/coder279/category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64 , error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string)(*model.Category,error)
	FindCategoryByLevel(uint32)([]model.Category,error)
	FindCategoryByParent(int64)([]model.Category,error)
}


//创建
func NewCategoryDataService(CategoryRepository repository.ICategoryRepository) ICategoryDataService{
	return &CategoryDataService{ CategoryRepository }
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}


//插入
func (u *CategoryDataService) AddCategory(Category *model.Category) (int64 ,error) {
	 return u.CategoryRepository.CreateCategory(Category)
}

//删除
func (u *CategoryDataService) DeleteCategory(CategoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(CategoryID)
}

//更新
func (u *CategoryDataService) UpdateCategory(Category *model.Category) error {
	return u.CategoryRepository.UpdateCategory(Category)
}

//查找
func (u *CategoryDataService) FindCategoryByID(CategoryID int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(CategoryID)
}

//查找
func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}

//根据名称查找数据
func (u *CategoryDataService) FindCategoryByName(categoryName string)(*model.Category,error){
	return u.CategoryRepository.FindCategoryByName(categoryName)
}

//根据等级查找
func (u *CategoryDataService) FindCategoryByLevel(level uint32)([]model.Category,error){
	return u.CategoryRepository.FindCategoryByLevel(level)
}

//根据父级查找
func (u *CategoryDataService) FindCategoryByParent(parent_id int64)([]model.Category,error){
	return u.CategoryRepository.FindCategoryByParent(parent_id)
}