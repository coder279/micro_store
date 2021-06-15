package handler
import (
	"context"
    "github.com/coder279/micro_store/domain/service"
	"github.com/coder279/micro_store/proto/category"
)

type Category struct{
	CategoryDataService service.ICategoryDataService
}

func (c *Category)CreateCategory(ctx context.Context,res *service.CategoryDataService.)




