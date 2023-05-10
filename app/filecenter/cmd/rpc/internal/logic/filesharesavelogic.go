package logic

import (
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileShareSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileShareSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileShareSaveLogic {
	return &FileShareSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileShareSaveLogic) FileShareSave(in *pb.FileShareSaveReq) (*pb.FileShareSaveResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileShareSaveResp{}, nil
}
