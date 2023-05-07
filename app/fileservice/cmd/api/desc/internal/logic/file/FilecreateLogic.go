package file

import (
	"context"

	"CloudMind/app/fileservice/cmd/api/desc/internal/svc"
	"CloudMind/app/fileservice/cmd/api/desc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilecreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilecreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilecreateLogic {
	return &FilecreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilecreateLogic) Filecreate(req *types.FileCreateReq) (resp *types.FileCreateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
