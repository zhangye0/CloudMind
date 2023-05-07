package logic

import (
	"context"

	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RealNameAuthenticationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRealNameAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RealNameAuthenticationLogic {
	return &RealNameAuthenticationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RealNameAuthenticationLogic) RealNameAuthentication(in *pb.RealNameAuthenticationReq) (*pb.RealNameAuthenticationResp, error) {
	// todo: add your logic here and delete this line

	return &pb.RealNameAuthenticationResp{}, nil
}
