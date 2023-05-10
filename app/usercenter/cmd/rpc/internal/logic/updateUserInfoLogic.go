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
		_, err = l.svcCtx.UserModel.UpdateOneMapById(l.ctx, in.UserId, map[string]interface{}{
			"Sex": in.Filed3,
		})
	case "Memory":
		User, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
		if err != nil {
			return nil, err
		}
		memory := User.Memory

		_, err = l.svcCtx.UserModel.UpdateOneMapById(l.ctx, in.UserId, map[string]interface{}{
			"Memory": memory + in.Filed4,
		})
	case "Flow":
		User, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
		if err != nil {
			return nil, err
		}
		flow := User.Flow

		_, err = l.svcCtx.UserModel.UpdateOneMapById(l.ctx, in.UserId, map[string]interface{}{
			"Flow": flow + in.Filed4,
		})
	case "Money":
		User, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
		if err != nil {
			return nil, err
		}
		money := User.Money
		_, err = l.svcCtx.UserModel.UpdateOneMapById(l.ctx, in.UserId, map[string]interface{}{
			"Money": money + in.Filed4,
		})
	case "Star":
	default:
		err = errors.New("修改信息类型错误")
	}
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserInfoResp{}, nil
}
