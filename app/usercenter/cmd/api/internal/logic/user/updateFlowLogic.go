package user

import (
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFlowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFlowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlowLogic {
	return &UpdateFlowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFlowLogic) UpdateFlow(req *types.UpdateFlowReq) (*types.UpdateFlowResp, error) {
	userid := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		UserId:     userid,
		UpdateType: "Flow",
		Filed4:     req.Flow,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
