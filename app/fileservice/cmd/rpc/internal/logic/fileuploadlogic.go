package logic

import (
	"context"

	"CloudMind/app/fileservice/cmd/rpc/pb/internal/svc"
	"CloudMind/app/fileservice/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadLogic) FileUpload(in *pb.FileUploadReq) (*pb.FileUploadResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileUploadResp{}, nil
}
