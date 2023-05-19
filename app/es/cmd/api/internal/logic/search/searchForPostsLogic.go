package search

import (
	"CloudMind/app/es/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/es/cmd/api/internal/svc"
	"CloudMind/app/es/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchForPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForPostsLogic {
	return &SearchForPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchForPostsLogic) SearchForPosts(req *types.SearchForPostsReq) (*types.SearchForPostsResp, error) {
	Resp, err := l.svcCtx.EsRpc.SearchForPosts(l.ctx, &pb.SearchForPostsReq{
		Content: req.Content,
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
	return &types.SearchForPostsResp{
		Posts: Posts,
		Error: Resp.Error,
	}, nil
}
