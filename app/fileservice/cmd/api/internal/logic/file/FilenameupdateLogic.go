package file

import (
	"context"

	"CloudMind/app/fileservice/cmd/api/internal/svc"
	"CloudMind/app/fileservice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilenameupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilenameupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilenameupdateLogic {
	return &FilenameupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilenameupdateLogic) Filenameupdate(req *types.FileNameUpdateReq) (resp *types.FileNameUpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
