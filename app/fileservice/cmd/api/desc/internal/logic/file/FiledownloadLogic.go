package file

import (
	"context"

	"CloudMind/app/fileservice/cmd/api/desc/internal/svc"
	"CloudMind/app/fileservice/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FiledownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFiledownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FiledownloadLogic {
	return &FiledownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FiledownloadLogic) Filedownload(req *types.FileDownloadReq) (resp *types.FileDownloadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
