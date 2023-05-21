package logic

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/app/usercenter/model"
	"CloudMind/common/xerr"
	"context"
	"github.com/jinzhu/copier"
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

/*
参数: 登录类型(string), 登录Key(string), 登录密码(string)

返回值: 令牌内容(string)，过期时间(int64), 刷新时间(int64)
*/
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	Resp, err := l.svcCtx.Cache.Take(in.AuthKey+in.Password, func() (interface{}, error) {
		var err error
		var userId int64

		switch in.AuthType {

		case model.UserAuthTypeEmail:
			userId, err = l.EmailLogin(in.AuthKey, in.Password)
			if err == model.ErrNotFound {
				return &pb.LoginResp{
					Error: "邮箱不存在",
				}, nil
			}
			if err != nil {
				return nil, err
			}
			if userId == -1 {
				return &pb.LoginResp{
					Error: "密码错误",
				}, nil
			}

		case model.UserAuthTypeQq:

		case model.UserAuthTypeWx:

		}

		// 构建生成令牌的对象
		generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)

		// 生成JWT令牌
		TokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
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

/*
EmailLogin
参数: 邮箱号(string), 登录密码(string)

返回: 用户id(int64), 错误信息(error)
*/
func (l *LoginLogic) EmailLogin(email, password string) (int64, error) {
	// 去数据库中找
	User, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		return 0, err
	}
	if User.Password != password {
		return -1, nil
	}
	return User.Id, nil
}
