package search

import (
	"CloudMind/app/es/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/es/cmd/api/internal/svc"
	"CloudMind/app/es/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForFilesByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchForFilesByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForFilesByUserIdLogic {
	return &SearchForFilesByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchForFilesByUserIdLogic) SearchForFilesByUserId(req *types.SearchForFilesByUserIdReq) (*types.SearchForFilesByUserIdResp, error) {
	Resp, err := l.svcCtx.EsRpc.SearchForFilesByUserId(l.ctx, &pb.SearchForFilesByUserIdReq{
		UserId:    req.UserId,
		TypeMount: req.TypeMount,
	})

	if err != nil {
		return nil, err
	}
	if Resp.Error != "" {
		return nil, errorx.NewDefaultError(Resp.Error)
	}
	var Files []types.File
	for _, file := range Resp.Files {
		Files = append(Files, types.File{
			Title: file.Title,
			Id:    file.Id,
		})
	}
	return &types.SearchForFilesByUserIdResp{
		Files: Files,
		Error: Resp.Error,
	}, nil
}
