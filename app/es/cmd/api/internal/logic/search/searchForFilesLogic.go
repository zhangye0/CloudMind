package search

import (
	"CloudMind/app/es/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/es/cmd/api/internal/svc"
	"CloudMind/app/es/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForFilesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchForFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForFilesLogic {
	return &SearchForFilesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchForFilesLogic) SearchForFiles(req *types.SearchForFilesReq) (*types.SearchForFilesResp, error) {
	Resp, err := l.svcCtx.EsRpc.SearchForFiles(l.ctx, &pb.SearchForFilesReq{
		Content: req.Content,
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
	return &types.SearchForFilesResp{
		Files: Files,
		Error: Resp.Error,
	}, nil
}
