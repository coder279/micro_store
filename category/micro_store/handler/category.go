package handler
import (
	"context"
	"github.com/coder279/category/common"
	"github.com/coder279/category/domain/model"
	"github.com/coder279/category/domain/service"
	"github.com/coder279/category/proto"
	"github.com/prometheus/common/log"
)

type Category struct{
	CategoryDataService service.ICategoryDataService
}

func (c *Category) DeleteCategory(context.Context, *go_micro_service_category.DeleteCategoryRequest, *go_micro_service_category.DeleteCategoryResponse) error {
	panic("implement me")
}

func (c *Category) FincAllCategory(context.Context, *go_micro_service_category.FindAllRequest, *go_micro_service_category.CategoryAllResponse) error {
	panic("implement me")
}

func categoryToResponse(CategorySlice []model.Category,response *go_micro_service_category.CategoryAllResponse){
	for _,cg := range CategorySlice{
		cr := &go_micro_service_category.CategoryResponse{}
		err := common.SwapTo(cg,cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category,cr)
	}
}

//创建分类的服务
func (c *Category)CreateCategory(ctx context.Context,req *go_micro_service_category.CategoryRequest,res *go_micro_service_category.CreateCategoryResponse)error{
	category := &model.Category{}
	//赋值
	err := common.SwapTo(req,category)
	if err != nil {
		return err
	}
	category_id,err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	res.Message = "添加成功"
	res.CategoryId = category_id
	return nil
}

//update分类服务
func (c *Category) UpdateCategory(ctx context.Context,req *go_micro_service_category.CategoryRequest,res *go_micro_service_category.UpdateCategoryResponse)error{
	category := &model.Category{}
	err := common.SwapTo(req,category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	res.Message = "分类更新成功"
	return nil
}

//删除分类服务
func (c *Category) DeleteCatory(ctx context.Context,req *go_micro_service_category.DeleteCategoryRequest,res *go_micro_service_category.DeleteCategoryResponse) error{
	err := c.CategoryDataService.DeleteCategory(req.CategoryId)
	if err != nil {
		return err
	}
	res.Message = "分类删除成功"
	return nil
}
//通过名字查找分类
func (c *Category) FindCategoryByName(ctx context.Context,req *go_micro_service_category.FindByNameRequest,res *go_micro_service_category.CategoryResponse) error{
	category,err := c.CategoryDataService.FindCategoryByName(req.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category,res)
}

//根据id进行查找
func (c *Category) FindCategoryById(ctx context.Context,req *go_micro_service_category.FindByIdRequest,res *go_micro_service_category.CategoryResponse) error{
	category,err := c.CategoryDataService.FindCategoryByID(req.Id)
	if err != nil {
		return err
	}
	return common.SwapTo(category,res)

}

//根据parent查找parent
func (c *Category) FindCategoryByParent(ctx context.Context,req *go_micro_service_category.FindByParentRequest,res *go_micro_service_category.CategoryAllResponse) error {
	category, err := c.CategoryDataService.FindCategoryByParent(req.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(category,res)
	return nil
}

func (c *Category) FindCategoryByLevel(ctx context.Context,req *go_micro_service_category.FindByLevelRequest,res *go_micro_service_category.CategoryAllResponse) error{
	category,err := c.CategoryDataService.FindCategoryByLevel(req.Level)
	if err != nil {
		return err
	}
	categoryToResponse(category,res)
	return nil
}

func (c *Category) FindCategoryAll(ctx context.Context,req *go_micro_service_category.FindAllRequest,res *go_micro_service_category.CategoryAllResponse) error{
	category,err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(category,res)
	return nil
}





