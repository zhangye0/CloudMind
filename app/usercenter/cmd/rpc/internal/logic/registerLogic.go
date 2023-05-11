package logic

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/app/usercenter/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
参数: 用户名, 密码, 邮箱,验证码, 登录类型
返回值: 令牌内容(string)，过期时间(int64), 刷新时间(int64)
*/
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {

	// 判断是否已经注册过
	_, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.UserAuth.AuthType, in.UserAuth.AuthKey)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	// 没有注册过
	if err != nil {
		// 新增用户
		_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
			Email:      in.Email,
			Nickname:   in.NickName,
			Password:   in.PassWord,
			CreateTime: time.Now().Unix(),
			Memory:     5120,
			Flow:       5120,
		})
		if err != nil {
			return nil, err
		}

		// 对用户数量进行+1
		UserNumber, err := l.svcCtx.RedisClient.Incr("UserNumber")
		if err != nil {
			return nil, err
		}

		// New生成令牌的函数
		generateToken := NewGenerateTokenLogic(l.ctx, l.svcCtx)

		// 生成令牌
		TokenResp, err := generateToken.GenerateToken(&pb.GenerateTokenReq{
			UserId: UserNumber,
		})
		if err != nil {
			return nil, err
		}

		// 插入授权信息
		_, err = l.svcCtx.UserAuthModel.Insert(l.ctx, &model.UserAuth{
			UserId:     UserNumber,
			AuthKey:    in.UserAuth.AuthKey,
			AuthType:   in.UserAuth.AuthType,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			DeleteTime: time.Now(),
		})
		if err != nil {
			return nil, err
		}

		return &pb.RegisterResp{
			AccessToken:  TokenResp.AccessToken,
			AccessExpire: TokenResp.AccessExpire,
			RefreshAfter: TokenResp.RefreshAfter,
		}, nil

	}
	return nil, errors.New("该邮箱已经注册过，请勿重复注册")
}
