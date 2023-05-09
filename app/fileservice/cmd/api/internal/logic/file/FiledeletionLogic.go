package file

import (
	"context"

	"CloudMind/app/fileservice/cmd/api/internal/svc"
	"CloudMind/app/fileservice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FiledeletionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFiledeletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FiledeletionLogic {
	return &FiledeletionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FiledeletionLogic) Filedeletion(req *types.FileDeletionReq) (resp *types.FileDeletionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
