package user

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailloginLogic {
	return &EmailloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailloginLogic) Emaillogin(req *types.WxLoginReq) (resp *types.WxLoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
