package logic

import (
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMoveLogic {
	return &FileMoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileMoveLogic) FileMove(in *pb.FileMoveReq) (*pb.FileMoveResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileMoveResp{}, nil
}
