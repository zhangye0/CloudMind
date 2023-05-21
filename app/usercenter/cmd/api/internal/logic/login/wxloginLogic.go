package login

import (
	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"
	"context"

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

func (l *WxloginLogic) Wxlogin(req *types.WxLoginReq) (*types.WxLoginResp, error) {
	// todo: add your logic here and delete this line
	return nil, nil
}
