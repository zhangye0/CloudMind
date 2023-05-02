package user

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetinfoLogic {
	return &GetinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetinfoLogic) Getinfo(req *types.GetInfoReq) (resp *types.GetInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
