package user

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RealNameAuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRealNameAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RealNameAuthenticationLogic {
	return &RealNameAuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RealNameAuthenticationLogic) RealNameAuthentication(req *types.RealNameAuthenticationReq) (resp *types.RealNameAuthenticationResp, err error) {
	// todo: add your logic here and delete this line
	return
}
