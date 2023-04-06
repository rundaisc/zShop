package admin

import (
	"context"

	"zShop/internal/svc"
	"zShop/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateHanderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUpdateHanderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateHanderLogic {
	return &AdminUpdateHanderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminUpdateHanderLogic) AdminUpdateHander(req *types.AdminUpdateRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
