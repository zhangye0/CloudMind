package logic

import (
	"context"
	"strconv"

	Espb "CloudMind/app/es/cmd/rpc/pb"
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *pb.LogoutReq) (*pb.LogoutResp, error) {
	_, err := l.svcCtx.UserModel.UpdateOneMapById(l.ctx, in.UserId, map[string]interface{}{
		"del_state": 0,
	})

	if err != nil {
		return nil, err
	}
	// 退出登录
	l.svcCtx.EsRpc.AddText(l.ctx, &Espb.AddTextReq{
		Index: "logout",
		Text:  strconv.Itoa(int(in.UserId)),
	})
	return &pb.LogoutResp{}, nil
}
