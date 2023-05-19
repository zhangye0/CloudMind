package search

import (
	"CloudMind/app/es/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/es/cmd/api/internal/svc"
	"CloudMind/app/es/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForFilesRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchForFilesRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForFilesRankLogic {
	return &SearchForFilesRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchForFilesRankLogic) SearchForFilesRank(req *types.SearchForFilesRankReq) (resp *types.SearchForFilesRankResp, err error) {
	Resp, err := l.svcCtx.EsRpc.SearchForFileRank(l.ctx, &pb.SearchForFileRankReq{
		Rank:      req.Rank,
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
	return &types.SearchForFilesRankResp{
		Files: Files,
		Error: Resp.Error,
	}, nil
}
