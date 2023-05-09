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
	case "Avatar":
		// 修改图片的MD5值
		err = l.UpdateUserAvatar(in.UserId, in.Filed2)
	default:
		err = errors.New("修改信息类型错误")
	}
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserInfoResp{}, nil
}

func (l *UpdateUserInfoLogic) UpdateUserAvatar(UserId int64, Md5 string) error {
	_, err := l.svcCtx.UserModel.Update(l.ctx, UserId, &model.User{
		Avatar: Md5,
	})
	// 查找图片表中是否存在
	ok, err := l.svcCtx.Bloom.Exists([]byte(Md5))
	if err != nil {
		return err
	}
	if !ok {
		err := l.svcCtx.Bloom.Add([]byte(Md5))
		if err != nil {
			return err
		}
	}
	return nil
}
