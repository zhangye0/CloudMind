package logic

import (
	"context"

	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForLogic {
	return &SearchForLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchForLogic) SearchFor(in *pb.SearchForReq) (*pb.SearchForResp, error) {
	l.svcCtx.Es.Search()
	return &pb.SearchForResp{}, nil
}
