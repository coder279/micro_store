package handler
import (
	"context"
    "github.com/coder279/category/domain/service"
	category "github.com/coder279/category/proto/category"
)

type Category struct{
	CategoryDataService service.ICategoryDataService
}

func (c *Category)CreateCategory(ctx context.Context,res *service.CategoryDataService.)




