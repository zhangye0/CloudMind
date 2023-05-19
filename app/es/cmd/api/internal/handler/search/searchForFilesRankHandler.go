package search

import (
	"net/http"

	"CloudMind/app/es/cmd/api/internal/logic/search"
	"CloudMind/app/es/cmd/api/internal/svc"
	"CloudMind/app/es/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchForFilesRankHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchForFilesRankReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := search.NewSearchForFilesRankLogic(r.Context(), svcCtx)
		resp, err := l.SearchForFilesRank(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
