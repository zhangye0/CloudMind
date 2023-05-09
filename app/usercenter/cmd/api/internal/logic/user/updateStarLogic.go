package user

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStarLogic {
	return &UpdateStarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStarLogic) UpdateStar(req *types.UpdateStarReq) (resp *types.UpdateStarResp, err error) {
	// todo: add your logic here and delete this line
	return
}
