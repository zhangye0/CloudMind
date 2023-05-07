package logic

import (
	"CloudMind/app/usercenter/model"
	"context"
	"errors"

	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	UserInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	if err != nil && err == model.ErrNotFound {
		return nil, errors.New("非法访问")
	}
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResp{
		UserInfo: &pb.User{
			Id:       UserInfo.Id,
			Email:    UserInfo.Email,
			Nickname: UserInfo.Nickname,
			Sex:      UserInfo.Sex,
			Avatar:   UserInfo.Avatar,
			Name:     UserInfo.Name,
			IdCard:   UserInfo.Idcard,
		},
	}, nil
}
