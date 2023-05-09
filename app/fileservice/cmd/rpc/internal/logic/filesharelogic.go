package logic

import (
	"context"

	"CloudMind/app/fileservice/cmd/rpc/internal/svc"
	"CloudMind/app/fileservice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileShareLogic {
	return &FileShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileShareLogic) FileShare(in *pb.FileShareReq) (*pb.FileShareResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileShareResp{}, nil
}
