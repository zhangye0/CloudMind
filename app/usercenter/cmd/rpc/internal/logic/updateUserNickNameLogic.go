package logic

import (
	"CloudMind/app/usercenter/model"
	"context"

	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserNickNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserNickNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserNickNameLogic {
	return &UpdateUserNickNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserNickNameLogic) UpdateUserNickName(in *pb.UpdateUserNickNameReq) (*pb.UpdateUserNickNameResp, error) {
	l.Logger.Error(in.UserId)
	_, err := l.svcCtx.UserModel.Update(l.ctx, in.UserId, &model.User{
		Nickname: in.NickName,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserNickNameResp{}, nil
}
