package logic

import (
	"context"

	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcloginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAcloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcloginLogic {
	return &AcloginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AcloginLogic) Aclogin(in *pb.AccountLoginReq) (*pb.AccountLoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AccountLoginResp{}, nil
}
