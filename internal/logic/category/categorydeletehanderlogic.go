package category

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zShop/internal/svc"
)

type CategoryDeleteHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryDeleteHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryDeleteHanderLogic {
	return &CategoryDeleteHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryDeleteHanderLogic) CategoryDeleteHander() error {
	// todo: add your logic here and delete this line

	return nil
}
