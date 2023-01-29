package logic

import (
	"context"

	"zShop/internal/svc"
	"zShop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryHanderLogic {
	return &CategoryHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryHanderLogic) CategoryHander(req *types.CategoryRequest) (resp *types.CategoryResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
