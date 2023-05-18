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

type SearchForFilesByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForFilesByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForFilesByIdLogic {
	return &SearchForFilesByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找某个用户收藏/下载/点赞过的文件
func (l *SearchForFilesByIdLogic) SearchForFilesById(in *pb.SearchForFilesByIdReq) (*pb.SearchForFilesByIdResp, error) {
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
		return &pb.SearchForFilesByIdResp{
			Error: fmt.Sprintf("Error encoding query: %s", err),
		}, nil
	}

	// 搜索
	res, err := l.svcCtx.Es.Search(
		l.svcCtx.Es.Search.WithContext(context.Background()),
		l.svcCtx.Es.Search.WithIndex("files"),
		l.svcCtx.Es.Search.WithBody(&buf),
		l.svcCtx.Es.Search.WithTrackTotalHits(true),
		l.svcCtx.Es.Search.WithPretty(),
	)
	if err != nil {
		return &pb.SearchForFilesByIdResp{
			Error: fmt.Sprintf("Error indexing the document: %s", err),
		}, nil
	}
	defer res.Body.Close()
	if res.IsError() {
		return &pb.SearchForFilesByIdResp{
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

	return &pb.SearchForFilesByIdResp{
		Files: Files,
		Error: "",
	}, nil
}
