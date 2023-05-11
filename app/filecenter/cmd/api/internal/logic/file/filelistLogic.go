package file

import (
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilelistLogic {
	return &FilelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilelistLogic) Filelist(req *types.FileListReq) (resp *types.FileListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
