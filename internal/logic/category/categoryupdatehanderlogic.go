package category

import (
	"context"

	"zShop/internal/svc"
	"zShop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryUpdateHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryUpdateHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryUpdateHanderLogic {
	return &CategoryUpdateHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryUpdateHanderLogic) CategoryUpdateHander(req *types.CategotyUpdateRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
