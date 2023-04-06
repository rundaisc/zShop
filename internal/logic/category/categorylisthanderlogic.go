package category

import (
	"context"

	"zShop/internal/svc"
	"zShop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryListHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListHanderLogic {
	return &CategoryListHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListHanderLogic) CategoryListHander(req *types.CategoryListRequest) (resp *types.CategoryListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
