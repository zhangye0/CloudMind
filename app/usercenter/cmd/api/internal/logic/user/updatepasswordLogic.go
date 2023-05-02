package user

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatepasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatepasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatepasswordLogic {
	return &UpdatepasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatepasswordLogic) Updatepassword(req *types.UpdatePassWordReq) (resp *types.UpdatePassWordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
