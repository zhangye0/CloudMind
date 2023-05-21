package search

import (
	"CloudMind/app/es/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/es/cmd/api/internal/svc"
	"CloudMind/app/es/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForPostsByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchForPostsByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForPostsByUserIdLogic {
	return &SearchForPostsByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchForPostsByUserIdLogic) SearchForPostsByUserId(req *types.SearchForPostsByUserIdReq) (*types.SearchForPostsByUserIdResp, error) {
	Resp, err := l.svcCtx.EsRpc.SearchForPostsByUserId(l.ctx, &pb.SearchForPostsByUserIdReq{
		UserId:    req.UserId,
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
	return &types.SearchForPostsByUserIdResp{
		Posts: Posts,
		Error: Resp.Error,
	}, nil
}
