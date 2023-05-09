package logic

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/app/usercenter/model"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	var err error
	switch in.UpdateType {
	case "PassWord":
		_, err = l.svcCtx.UserModel.Update(l.ctx, in.UserId, &model.User{
			Password: in.Field1,
		})
	case "NickName":
		_, err = l.svcCtx.UserModel.Update(l.ctx, in.UserId, &model.User{
			Nickname: in.Field1,
		})
	case "Sex":
		_, err = l.svcCtx.UserModel.Update(l.ctx, in.UserId, &model.User{
			Sex: in.Filed3,
		})
	default:
		err = errors.New("修改信息类型错误")
	}
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserInfoResp{}, nil
}
