package search

import (
	"CloudMind/app/es/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/es/cmd/api/internal/svc"
	"CloudMind/app/es/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForPostsRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchForPostsRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForPostsRankLogic {
	return &SearchForPostsRankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchForPostsRankLogic) SearchForPostsRank(req *types.SearchForPostsRankReq) (*types.SearchForPostsRankResp, error) {
	Resp, err := l.svcCtx.EsRpc.SearchForPostsRank(l.ctx, &pb.SearchForPostsRankReq{
		Rank:      req.Rank,
		TypeMount: req.TypeMount,
	})

	if err != nil {
		return nil, err
	}
	if Resp.Error != "" {
		return nil, errorx.NewDefaultError(Resp.Error)
	}
	var Posts []types.Posts
	for _, post := range Resp.Posts {
		Posts = append(Posts, types.Posts{
			Title:   post.Title,
			Content: post.Content,
			Id:      post.Id,
		})
	}
	return &types.SearchForPostsRankResp{
		Posts: Posts,
		Error: Resp.Error,
	}, nil
}
