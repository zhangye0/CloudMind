package user

import (
	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (*types.GetUserInfoResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	Resp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		UserInfo: types.User{
			Id:       Resp.UserInfo.Id,
			Email:    Resp.UserInfo.Email,
			NickName: Resp.UserInfo.Nickname,
			Sex:      Resp.UserInfo.Sex,
			Avatar:   Resp.UserInfo.Avatar,
			Name:     Resp.UserInfo.Name,
			IdCard:   Resp.UserInfo.IdCard,
		},
	}, nil
}
