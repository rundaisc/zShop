package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zShop/internal/logic/category"
	"zShop/internal/svc"
)

func CategoryDeleteHanderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewCategoryDeleteHanderLogic(r.Context(), svcCtx)
		err := l.CategoryDeleteHander()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
