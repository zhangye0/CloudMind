package logic

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/app/usercenter/model"
	"context"
	"github.com/pkg/errors"
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

/*
参数: 用户id(int64), 修改类型(string)，字段1(string),字段2(string),字段3(int64),字段4(float64),
返回值: 空结构体, 错误信息(error)
tips: 修改类型包括PassWord,NickName,Sex,Memory,Flow,Money,Star,AllFlow,AllMemory
如果需要修改密码/用户名，将密码/用户名写在字段1
如果需要修改性别，将性别写在字段3
如果需要修改内存/流量/余额，将内存/流量/余额写在字段4
All开头代表给所有的人都增加
*/
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
	case "Memory", "Flow", "Money":
		_, err := l.svcCtx.UserModel.AddOne(l.ctx, in.UserId, in.UpdateType, in.Filed4)
		if err != nil {
			return nil, err
		}
	case "Star":
	case "AllFlow":
		_, err := l.svcCtx.UserModel.UpdateAll(l.ctx, "Flow", in.Filed4)
		if err != nil {
			return nil, err
		}
	case "AllMemory":
		_, err := l.svcCtx.UserModel.AddAll(l.ctx, "Memory", in.Filed4)
		if err != nil {
			return nil, err
		}
	default:
		err = errors.New("修改类型错误")
	}
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserInfoResp{}, nil
}
