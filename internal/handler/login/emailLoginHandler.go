package login

import (
	"net/http"

	"CloudMind/internal/logic/login"
	"CloudMind/internal/svc"
	"CloudMind/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EmailLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewEmailLoginLogic(r.Context(), svcCtx)
		resp, err := l.EmailLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
