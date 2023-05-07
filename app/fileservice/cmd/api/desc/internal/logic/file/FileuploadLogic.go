package file

import (
	"context"

	"CloudMind/app/fileservice/cmd/api/desc/internal/svc"
	"CloudMind/app/fileservice/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileuploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileuploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileuploadLogic {
	return &FileuploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileuploadLogic) Fileupload(req *types.FileUploadReq) (resp *types.FileUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
