package user

import (
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAvatarLogic) UpdateAvatar(req *types.UpdateAvatarReq) (*types.UpdateAvatarResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		UserId:     userId,
		UpdateType: "Avatar",
		Field1:     req.Avatar,
		Filed2:     req.Md5,
		Filed3:     0,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateAvatarResp{}, nil
}
