package admin

import (
	"context"

	"zShop/internal/svc"
	"zShop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminListHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminListHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminListHanderLogic {
	return &AdminListHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminListHanderLogic) AdminListHander(req *types.AdminListRequest) (resp *types.AdminListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
