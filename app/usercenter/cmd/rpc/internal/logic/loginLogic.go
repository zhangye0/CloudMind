package logic

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/app/usercenter/model"
	"CloudMind/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrLoginError = xerr.NewErrMsg("账号或密码错误")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error
	switch in.AuthType {
	case model.UserAuthTypeEmail:
		UserInfo, errs := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.AuthKey)
		err = errs
		userId = UserInfo.Id
	case model.UserAuthTypeQq:
	case model.UserAuthTypeWx:
	default:
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, errors.Wrapf(ErrLoginError, "GenerateToken userId: %d", userId)
	}
	geneateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	TokenResp, err := geneateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId: %d", userId)
	}
	//
	return &pb.LoginResp{
		AccessToken:  TokenResp.AccessToken,
		AccessExpire: TokenResp.AccessExpire,
		RefreshAfter: TokenResp.RefreshAfter,
	}, nil
}
