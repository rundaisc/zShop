package category

import (
	"context"

	"zShop/internal/svc"
	"zShop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryCreateHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryCreateHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryCreateHanderLogic {
	return &CategoryCreateHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryCreateHanderLogic) CategoryCreateHander(req *types.CategotyCreateRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
