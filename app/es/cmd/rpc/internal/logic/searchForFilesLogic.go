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

type SearchForFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForFilesLogic {
	return &SearchForFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
根据Content进行全文模糊搜索， 搜索相关的文件
*/
func (l *SearchForFilesLogic) SearchForFiles(in *pb.SearchForFilesReq) (*pb.SearchForFilesResp, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"multi_match": map[string]interface{}{
							"query":     in.Content,
							"fields":    []string{"*"},
							"fuzziness": "AUTO",
						},
					},
					{
						"match": map[string]interface{}{
							"typeMount": "upload",
						},
					},
				},
			},
		},
	}

	// 序列化
	err := json.NewEncoder(&buf).Encode(query)
	if err != nil {
		return &pb.SearchForFilesResp{
			Error: fmt.Sprintf("Error marshaling the document: %s", err),
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
		return &pb.SearchForFilesResp{
			Error: fmt.Sprintf("Error indexing the document: %s", err),
		}, nil
	}
	defer res.Body.Close()
	if res.IsError() {
		return &pb.SearchForFilesResp{
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

	return &pb.SearchForFilesResp{
		Files: Files,
		Error: "",
	}, nil
}
