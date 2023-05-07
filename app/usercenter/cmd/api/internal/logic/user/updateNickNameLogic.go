package user

import (
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNickNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNickNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNickNameLogic {
	return &UpdateNickNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNickNameLogic) UpdateNickName(req *types.UpdateNickNameReq) (*types.UpdateNickNameResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.UsercenterRpc.UpdateUserNickName(l.ctx, &pb.UpdateUserNickNameReq{
		UserId:   userId,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateNickNameResp{}, nil
}
