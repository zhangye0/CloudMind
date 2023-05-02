package user

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateemailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateemailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateemailLogic {
	return &UpdateemailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateemailLogic) Updateemail(req *types.UpdateEmailReq) (resp *types.UpdateEmailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
