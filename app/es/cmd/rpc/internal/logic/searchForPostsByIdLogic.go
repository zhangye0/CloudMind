package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForPostsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForPostsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForPostsByIdLogic {
	return &SearchForPostsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchForPostsByIdLogic) SearchForPostsById(in *pb.SearchForPostsByIdReq) (*pb.SearchForPostsByIdResp, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"userId": in.Id,
						},
					},
					{
						"term": map[string]interface{}{
							"typeMount": in.Type,
						},
					},
				},
			},
		},
	}

	// 序列化
	err := json.NewEncoder(&buf).Encode(query)
	if err != nil {
		return &pb.SearchForPostsByIdResp{
			Error: fmt.Sprintf("Error encoding query: %s", err),
		}, nil
	}

	// 搜索
	res, err := l.svcCtx.Es.Search(
		l.svcCtx.Es.Search.WithContext(context.Background()),
		l.svcCtx.Es.Search.WithIndex("posts"),
		l.svcCtx.Es.Search.WithBody(&buf),
		l.svcCtx.Es.Search.WithTrackTotalHits(true),
		l.svcCtx.Es.Search.WithPretty(),
	)
	if err != nil {
		return &pb.SearchForPostsByIdResp{
			Error: fmt.Sprintf("Error Search the document: %s", err),
		}, nil
	}
	defer res.Body.Close()
	if res.IsError() {
		return &pb.SearchForPostsByIdResp{
			Error: fmt.Sprintf("[%s] Error Search document ID", res.Status()),
		}, nil
	}
	//反序列化搜索结果
	var r map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&r)
	var Posts []*pb.Post
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		t := pb.Post{
			Title:   source.(map[string]interface{})["title"].(string),
			Content: source.(map[string]interface{})["content"].(string),
			Id:      source.(map[string]interface{})["id"].(string),
		}
		Posts = append(Posts, &t)
	}

	return &pb.SearchForPostsByIdResp{
		Posts: Posts,
		Error: "",
	}, nil
}
