package login

import (
	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
通过邮箱号,密码进行登录， 返回JWT令牌
*/
func (l *EmailLoginLogic) EmailLogin(req *types.EmailLoginReq) (*types.EmailLoginResp, error) {
	Resp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &pb.LoginReq{
		AuthType: "email",
		AuthKey:  req.Email,
		Password: req.PassWord,
	})
	if Resp.Error != "" {
		return nil, errorx.NewDefaultError(Resp.Error)
	}
	if err != nil {
		return nil, err
	}

	return &types.EmailLoginResp{
		AccessToken:  Resp.AccessToken,
		AccessExpire: Resp.AccessExpire,
		RefreshAfter: Resp.RefreshAfter,
	}, nil
}
