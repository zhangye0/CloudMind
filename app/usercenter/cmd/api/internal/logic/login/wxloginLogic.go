package login

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxloginLogic {
	return &WxloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxloginLogic) Wxlogin(req *types.WxLoginReq) (resp *types.WxLoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
