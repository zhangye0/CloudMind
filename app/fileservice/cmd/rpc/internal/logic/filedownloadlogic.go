package logic

import (
	"context"

	"CloudMind/app/fileservice/cmd/rpc/internal/svc"
	"CloudMind/app/fileservice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDownloadLogic) FileDownload(in *pb.FileDownloadReq) (*pb.FileDownloadResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileDownloadResp{}, nil
}
