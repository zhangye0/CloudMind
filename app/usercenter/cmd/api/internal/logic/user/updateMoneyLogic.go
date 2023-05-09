package user

import (
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMoneyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMoneyLogic {
	return &UpdateMoneyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMoneyLogic) UpdateMoney(req *types.UpdateMoneyReq) (*types.UpdateMoneyResp, error) {
	userid := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		UserId:     userid,
		UpdateType: "Money",
		Filed4:     req.Money,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
