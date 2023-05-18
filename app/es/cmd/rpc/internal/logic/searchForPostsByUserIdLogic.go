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

type SearchForPostsByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForPostsByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForPostsByUserIdLogic {
	return &SearchForPostsByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
查询用户下载/收藏/点赞过的帖子
Type: look(浏览), star(收藏), like(点赞)
UserId: 用户id
*/
func (l *SearchForPostsByUserIdLogic) SearchForPostsByUserId(in *pb.SearchForPostsByUserIdReq) (*pb.SearchForPostsByUserIdResp, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"userId": in.UserId,
						},
					},
					{
						"term": map[string]interface{}{
							"typeMount": in.TypeMount,
						},
					},
				},
			},
		},
	}

	// 序列化
	err := json.NewEncoder(&buf).Encode(query)
	if err != nil {
		return &pb.SearchForPostsByUserIdResp{
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
		return &pb.SearchForPostsByUserIdResp{
			Error: fmt.Sprintf("Error Search the document: %s", err),
		}, nil
	}
	defer res.Body.Close()
	if res.IsError() {
		return &pb.SearchForPostsByUserIdResp{
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

	return &pb.SearchForPostsByUserIdResp{
		Posts: Posts,
		Error: "",
	}, nil
}
