package login

import (
	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"context"
	"github.com/jinzhu/copier"
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

func (l *EmailLoginLogic) EmailLogin(req *types.EmailLoginReq) (*types.EmailLoginResp, error) {
	Resp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &pb.LoginReq{
		AuthType: "email",
		AuthKey:  req.Email,
		Password: req.PassWord,
	})
	if err != nil {
		return nil, err
	}
	var resp types.EmailLoginResp
	_ = copier.Copy(&resp, Resp)
	return &resp, nil

}
