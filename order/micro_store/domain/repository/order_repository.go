package repository
import (
	"github.com/jinzhu/gorm"
	"github.com/coder279/order/domain/model"
)
type IOrderRepository interface{
    InitTable() error
    FindOrderByID(int64) (*model.Order, error)
	CreateOrder(*model.Order) (int64, error)
	DeleteOrderByID(int64) error
	UpdateOrder(*model.Order) error
	FindAll()([]model.Order,error)

}
//创建orderRepository
func NewOrderRepository(db *gorm.DB) IOrderRepository  {
	return &OrderRepository{mysqlDb:db}
}

type OrderRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *OrderRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Order{}).Error
}

//根据ID查找Order信息
func (u *OrderRepository)FindOrderByID(orderID int64) (order *model.Order,err error) {
	order = &model.Order{}
	return order, u.mysqlDb.First(order,orderID).Error
}

//创建Order信息
func (u *OrderRepository) CreateOrder(order *model.Order) (int64, error) {
	return order.ID, u.mysqlDb.Create(order).Error
}

//根据ID删除Order信息
func (u *OrderRepository) DeleteOrderByID(orderID int64) error {
	return u.mysqlDb.Where("id = ?",orderID).Delete(&model.Order{}).Error
}

//更新Order信息
func (u *OrderRepository) UpdateOrder(order *model.Order) error {
	return u.mysqlDb.Model(order).Update(order).Error
}

//获取结果集
func (u *OrderRepository) FindAll()(orderAll []model.Order,err error) {
	return orderAll, u.mysqlDb.Find(&orderAll).Error
}

