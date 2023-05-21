package logic

import (
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileNameUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileNameUpdateLogic {
	return &FileNameUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileNameUpdateLogic) FileNameUpdate(in *pb.FileNameUpdateReq) (*pb.FileNameUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileNameUpdateResp{}, nil
}
