package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zShop/internal/logic/category"
	"zShop/internal/svc"
	"zShop/internal/types"
)

func CategoryListHanderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewCategoryListHanderLogic(r.Context(), svcCtx)
		resp, err := l.CategoryListHander(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
