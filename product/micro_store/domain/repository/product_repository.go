package repository
import (
	"github.com/jinzhu/gorm"
	"github.com/coder279/micro_store/product/domain/model"
)
type IProductRepository interface{
    InitTable() error
    FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll()([]model.Product,error)

}
//创建productRepository
func NewProductRepository(db *gorm.DB) IProductRepository  {
	return &ProductRepository{mysqlDb:db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *ProductRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Product{}).Error
}

//根据ID查找Product信息
func (u *ProductRepository)FindProductByID(productID int64) (product *model.Product,err error) {
	product = &model.Product{}
	return product, u.mysqlDb.First(product,productID).Error
}

//创建Product信息
func (u *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, u.mysqlDb.Create(product).Error
}

//根据ID删除Product信息
func (u *ProductRepository) DeleteProductByID(productID int64) error {
	return u.mysqlDb.Where("id = ?",productID).Delete(&model.Product{}).Error
}

//更新Product信息
func (u *ProductRepository) UpdateProduct(product *model.Product) error {
	return u.mysqlDb.Model(product).Update(product).Error
}

//获取结果集
func (u *ProductRepository) FindAll()(productAll []model.Product,err error) {
	return productAll, u.mysqlDb.Find(&productAll).Error
}

