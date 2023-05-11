package file

import (
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilemoveLogic {
	return &FilemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilemoveLogic) Filemove(req *types.FileMoveReq) (resp *types.FileMoveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
