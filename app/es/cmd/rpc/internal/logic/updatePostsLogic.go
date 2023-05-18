package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostsLogic {
	return &UpdatePostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostsLogic) UpdatePosts(in *pb.UpdatePostsReq) (*pb.UpdatePostsResp, error) {
	requestBody := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"id": in.Post.Id,
			},
		},
		"script": map[string]interface{}{
			"source": "ctx._source.title = params.new_title; ctx._source.content = params.new_content;",
			"params": map[string]interface{}{
				"new_title":   in.Post.Title,
				"new_content": in.Post.Content,
			},
		},
	}

	// Create the Update By Query request object
	req := esapi.UpdateByQueryRequest{
		Index: []string{"posts"},
		Body:  esutil.NewJSONReader(requestBody),
	}

	// Send the Update By Query request
	res, err := req.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		return &pb.UpdatePostsResp{
			Error: fmt.Sprintf("Error executing the Do: %s", err),
		}, nil
	}
	defer res.Body.Close()

	if res.IsError() {
		return &pb.UpdatePostsResp{
			Error: fmt.Sprintf("update by query failed with code %d: %s", res.StatusCode, res.String()),
		}, nil
	}

	return &pb.UpdatePostsResp{}, nil
}
