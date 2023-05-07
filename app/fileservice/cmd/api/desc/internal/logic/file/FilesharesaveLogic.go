package file

import (
	"context"

	"CloudMind/app/fileservice/cmd/api/desc/internal/svc"
	"CloudMind/app/fileservice/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilesharesaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilesharesaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilesharesaveLogic {
	return &FilesharesaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilesharesaveLogic) Filesharesave(req *types.FileShareSaveReq) (resp *types.FileShareSaveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
