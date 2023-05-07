package logic

import (
	"context"

	"CloudMind/app/fileservice/cmd/rpc/pb/internal/svc"
	"CloudMind/app/fileservice/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeletionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDeletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeletionLogic {
	return &FileDeletionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDeletionLogic) FileDeletion(in *pb.FileDeletionReq) (*pb.FileDeletionResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FileDeletionResp{}, nil
}
