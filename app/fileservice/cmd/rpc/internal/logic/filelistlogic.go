package logic

import (
	"context"

	"CloudMind/app/fileservice/cmd/rpc/internal/svc"
	"CloudMind/app/fileservice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileListLogic) FileList(in *pb.FileListReq) (*pb.FileListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileListResp{}, nil
}
