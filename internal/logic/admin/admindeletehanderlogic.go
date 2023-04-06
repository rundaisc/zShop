package admin

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"zShop/internal/svc"
)

type AdminDeleteHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDeleteHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDeleteHanderLogic {
	return &AdminDeleteHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminDeleteHanderLogic) AdminDeleteHander() error {
	// todo: add your logic here and delete this line

	return nil
}
