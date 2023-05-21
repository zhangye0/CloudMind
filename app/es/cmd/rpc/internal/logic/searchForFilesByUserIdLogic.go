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

type SearchForFilesByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForFilesByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForFilesByUserIdLogic {
	return &SearchForFilesByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
查询用户下载/收藏/点赞过的文件
Type: upload(下载), star(收藏), like(点赞)
UserId: 用户id
*/
func (l *SearchForFilesByUserIdLogic) SearchForFilesByUserId(in *pb.SearchForFilesByUserIdReq) (*pb.SearchForFilesByUserIdResp, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"userId": in.UserId,
			},
		},
	}

	// 序列化
	err := json.NewEncoder(&buf).Encode(query)
	if err != nil {
		return &pb.SearchForFilesByUserIdResp{
			Error: fmt.Sprintf("Error encoding query: %s", err),
		}, nil
	}

	// 搜索
	res, err := l.svcCtx.Es.Search(
		l.svcCtx.Es.Search.WithContext(context.Background()),
		l.svcCtx.Es.Search.WithIndex(in.TypeMount+"files"),
		l.svcCtx.Es.Search.WithBody(&buf),
		l.svcCtx.Es.Search.WithTrackTotalHits(true),
		l.svcCtx.Es.Search.WithPretty(),
	)
	if err != nil {
		return &pb.SearchForFilesByUserIdResp{
			Error: fmt.Sprintf("Error indexing the document: %s", err),
		}, nil
	}
	defer res.Body.Close()
	if res.IsError() {
		return &pb.SearchForFilesByUserIdResp{
			Error: fmt.Sprintf("[%s] Error indexing document ID", res.Status()),
		}, nil
	}
	//反序列化搜索结果
	var r map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&r)
	var Files []*pb.File
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		t := pb.File{
			Title: source.(map[string]interface{})["title"].(string),
			Id:    source.(map[string]interface{})["id"].(string),
		}
		Files = append(Files, &t)
	}

	return &pb.SearchForFilesByUserIdResp{
		Files: Files,
		Error: "",
	}, nil
}
