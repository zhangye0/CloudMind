package user

import (
	"net/http"

	"CloudMind/app/usercenter/cmd/api/internal/logic/user"
	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateMoneyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateMoneyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateMoneyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMoney(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
