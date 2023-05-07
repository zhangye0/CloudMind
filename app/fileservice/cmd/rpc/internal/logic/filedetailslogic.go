package logic

import (
	"context"

	"CloudMind/app/fileservice/cmd/rpc/pb/internal/svc"
	"CloudMind/app/fileservice/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDetailsLogic {
	return &FileDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDetailsLogic) FileDetails(in *pb.FileDetailsReq) (*pb.FileDetailsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileDetailsResp{}, nil
}
