package logic

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/app/usercenter/model"
	"CloudMind/common/xerr"
	"context"
	"github.com/jinzhu/copier"
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
	Resp, err := l.svcCtx.Cache.Take(in.AuthKey+in.Password, func() (interface{}, error) {
		var err error
		var userId int64
		switch in.AuthType {
		case model.UserAuthTypeEmail:
			userId, err = l.EmailLogin(in.AuthKey, in.Password)
			if err != nil {
				return nil, err
			}
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
	})
	var resp pb.LoginResp
	_ = copier.Copy(&resp, Resp)
	return &resp, err
}

func (l *LoginLogic) EmailLogin(email, password string) (int64, error) {
	// 去数据库中找
	User, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil && err == model.ErrNotFound {
		return 0, errors.New("邮箱不存在")
	}
	if err != nil {
		return 0, err
	}
	if User.Password != password {
		return 0, errors.New("密码错误")
	}
	return User.Id, nil
}
