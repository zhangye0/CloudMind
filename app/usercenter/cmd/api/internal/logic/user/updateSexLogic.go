package user

import (
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSexLogic {
	return &UpdateSexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSexLogic) UpdateSex(req *types.UpdateSexReq) (*types.UpdateSexResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		UserId:     userId,
		UpdateType: "Sex",
		Field1:     "",
		Field2:     req.Sex,
	})

	if err != nil {
		return nil, err
	}
	return &types.UpdateSexResp{}, nil
}
