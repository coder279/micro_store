package repository
import (
	"github.com/jinzhu/gorm"
	"github.com/coder279/micro_store/domain/model"
)
type ICategoryRepository interface{
    InitTable() error
    FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll()([]model.Category,error)
    FindCategoryByLevel(uint32)([]model.Category,error)
    FindCategoryByParent(int64)([]model.Category,error)
    FindCategoryByName(string)(*model.Category,error)

}
//创建CategoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository  {
	return &CategoryRepository{mysqlDb:db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *CategoryRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

//根据ID查找Category信息
func (u *CategoryRepository)FindCategoryByID(CategoryID int64) (Category *model.Category,err error) {
	Category = &model.Category{}
	return Category, u.mysqlDb.First(Category,CategoryID).Error
}

//创建Category信息
func (u *CategoryRepository) CreateCategory(Category *model.Category) (int64, error) {
	return Category.ID, u.mysqlDb.Create(Category).Error
}

//根据ID删除Category信息
func (u *CategoryRepository) DeleteCategoryByID(CategoryID int64) error {
	return u.mysqlDb.Where("id = ?",CategoryID).Delete(&model.Category{}).Error
}

//更新Category信息
func (u *CategoryRepository) UpdateCategory(Category *model.Category) error {
	return u.mysqlDb.Model(Category).Update(Category).Error
}

//获取结果集
func (u *CategoryRepository) FindAll()(CategoryAll []model.Category,err error) {
	return CategoryAll, u.mysqlDb.Find(&CategoryAll).Error
}

//根据层级进行筛选
func(u *CategoryRepository) FindCategoryByLevel(level uint32)(Category []model.Category,err error){
	return Category,u.mysqlDb.Where("category_level = ?",level).Find(Category).Error
}

//根据名称进行筛选
func(u *CategoryRepository) FindCategoryByName(name string)(Category *model.Category,err error){
	Category = &model.Category{}
	return Category,u.mysqlDb.Where("category_name = ?",name).Find(Category).Error
}

//根据parent筛选
func (u *CategoryRepository) FindCategoryByParent(parent int64)(Category []model.Category,err error){
	return Category,u.mysqlDb.Where("category_parent = ?",parent).Find(Category).Error
}


