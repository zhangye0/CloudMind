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
		if errs != nil || UserInfo.Password != in.Password {
			return nil, errors.New("账号或密码错误")
		}
		userId = UserInfo.Id
	case model.UserAuthTypeQq:
	case model.UserAuthTypeWx:
	default:
		return nil, errors.New("不存在这种登录方式")
	}
	geneateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	TokenResp, err := geneateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResp{
		AccessToken:  TokenResp.AccessToken,
		AccessExpire: TokenResp.AccessExpire,
		RefreshAfter: TokenResp.RefreshAfter,
	}, nil
}
