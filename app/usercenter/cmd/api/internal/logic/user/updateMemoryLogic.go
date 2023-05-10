package user

import (
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMemoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemoryLogic {
	return &UpdateMemoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMemoryLogic) UpdateMemory(req *types.UpdateMemoryReq) (*types.UpdateMemoryResp, error) {
	userid := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		UserId:     userid,
		UpdateType: "Memory",
		Filed4:     req.Memory,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
