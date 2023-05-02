package logic

import (
	"context"

	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetinfoLogic {
	return &GetinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetinfoLogic) Getinfo(in *pb.GetInfoReq) (*pb.GetInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetInfoResp{}, nil
}
