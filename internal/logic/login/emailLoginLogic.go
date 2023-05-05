package login

import (
	"CloudMind/internal/svc"
	"CloudMind/internal/types"
	"context"
	"github.com/pkg/errors"

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

func (l *EmailLoginLogic) EmailLogin(req *types.EmailLoginReq) (resp *types.EmailLoginResp, err error) {
	UserInfo, err := l.svcCtx.UserInfoModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil {
		return nil, errors.Errorf("邮箱不存在！")
	}
	if UserInfo.Password != req.PassWord {
		return nil, errors.Errorf("密码错误！")
	}
	return &types.EmailLoginResp{
		AccessToken:  "AA",
		AccessExpire: 5,
		RefreshAfter: 1,
	}, nil
}
