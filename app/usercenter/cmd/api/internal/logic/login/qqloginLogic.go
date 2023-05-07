package login

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QqloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQqloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QqloginLogic {
	return &QqloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QqloginLogic) Qqlogin(req *types.QqLoginReq) (resp *types.QqLoginResp, err error) {
	// todo: add your logic here and delete this line
	return
}
