package file

import (
	"context"

	"CloudMind/app/fileservice/cmd/api/internal/svc"
	"CloudMind/app/fileservice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileshareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileshareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileshareLogic {
	return &FileshareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileshareLogic) Fileshare(req *types.FileShareReq) (resp *types.FileShareResp, err error) {
	// todo: add your logic here and delete this line

	return
}
