package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zShop/internal/logic/admin"
	"zShop/internal/svc"
	"zShop/internal/types"
)

func AdminCreateHanderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewAdminCreateHanderLogic(r.Context(), svcCtx)
		err := l.AdminCreateHander(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
