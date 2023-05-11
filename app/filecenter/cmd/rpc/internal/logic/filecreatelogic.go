package logic

import (
	"CloudMind/app/filecenter/model"
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

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

	_, err := l.svcCtx.FileModel.Insert(l.ctx, &model.File{
		Name:       in.Name,
		Type:       "folder",
		Path:       "xxxxxx",
		Size:       "0",
		ShareLink:  "xxxxxx",
		ModifyTime: 0,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}
