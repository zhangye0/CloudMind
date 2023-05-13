package file

import (
	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"
	"CloudMind/app/filecenter/cmd/rpc/filecenter"
	"context"

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

func (l *FilecreateLogic) Filecreate(req *types.FileCreateReq) (*types.FileCreateResp, error) {

	resp, err := l.svcCtx.FileRpc.FileCreate(l.ctx, &filecenter.FileCreateReq{
		Path: req.Path,
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	if resp != nil {
		return nil, nil // 使用 errorx 报错
	}

	return nil, nil
}
