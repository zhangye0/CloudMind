package user

import (
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePassWordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePassWordLogic {
	return &UpdatePassWordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePassWordLogic) UpdatePassWord(req *types.UpdatePassWordReq) (*types.UpdatePassWordResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		UserId:     userId,
		UpdateType: "PassWord",
		Field1:     req.PassWord,
		Field2:     0,
	})

	if err != nil {
		return nil, err
	}
	return &types.UpdatePassWordResp{}, nil
}
