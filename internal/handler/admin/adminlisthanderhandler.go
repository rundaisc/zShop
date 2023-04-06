package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zShop/internal/logic/admin"
	"zShop/internal/svc"
	"zShop/internal/types"
)

func AdminListHanderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewAdminListHanderLogic(r.Context(), svcCtx)
		resp, err := l.AdminListHander(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
