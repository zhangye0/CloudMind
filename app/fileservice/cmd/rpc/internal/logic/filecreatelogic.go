package logic

import (
	"context"

	"CloudMind/app/fileservice/cmd/rpc/pb/internal/svc"
	"CloudMind/app/fileservice/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileCreateLogic {
	return &FileCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileCreateLogic) FileCreate(in *pb.FileCreateReq) (*pb.FileCreateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileCreateResp{}, nil
}
