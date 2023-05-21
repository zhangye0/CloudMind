package logic

import (
	"CloudMind/app/filecenter/model"
	"context"
	time2 "time"

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

	t := time2.Now()
	time := t.Unix()

	_, err := l.svcCtx.FileModel.TxInsert(l.ctx, l.svcCtx.GormDB, &model.File{
		Name:       in.Name,
		Type:       "folder",
		Path:       in.Path,
		Size:       "0",
		ShareLink:  "xxxxxx",
		ModifyTime: time,
	})

	if err != nil {
		return &pb.FileCreateResp{
			Error: "创建失败",
		}, err
	}

	return nil, nil
}
