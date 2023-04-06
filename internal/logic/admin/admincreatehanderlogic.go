package admin

import (
	"context"

	"zShop/internal/svc"
	"zShop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminCreateHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminCreateHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminCreateHanderLogic {
	return &AdminCreateHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminCreateHanderLogic) AdminCreateHander(req *types.AdminCreateRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
