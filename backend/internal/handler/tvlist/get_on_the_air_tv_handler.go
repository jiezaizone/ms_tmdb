package tvlist

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ms_tmdb/internal/logic/tvlist"
	"ms_tmdb/internal/svc"
	"ms_tmdb/internal/types"
)

func GetOnTheAirTvHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tvlist.NewGetOnTheAirTvLogic(r.Context(), svcCtx)
		err := l.GetOnTheAirTv(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
