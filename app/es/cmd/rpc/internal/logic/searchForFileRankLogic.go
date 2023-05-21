package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type SearchForFileRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForFileRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForFileRankLogic {
	return &SearchForFileRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
查询文件排行
TypeMount: upload(下载), star(收藏), like(点赞)
Rank: 前多少名
*/
func (l *SearchForFileRankLogic) SearchForFileRank(in *pb.SearchForFileRankReq) (*pb.SearchForFileRankResp, error) {

	FilesJson, err := l.svcCtx.Redis.Get("files" + in.TypeMount + strconv.Itoa(int(in.Rank)))
	if err != nil {
		var Files []*pb.File
		err = json.Unmarshal([]byte(FilesJson), &Files)
		if err != nil {
			return &pb.SearchForFileRankResp{
				Files: Files,
				Error: "",
			}, nil
		}
	}

	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"files": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "id.keyword",
					"size":  in.Rank,
				},
				"aggs": map[string]interface{}{
					"file_data": map[string]interface{}{
						"top_hits": map[string]interface{}{
							"size": in.Rank,
							"_source": map[string]interface{}{
								"includes": []string{"title", "id"},
							},
						},
					},
				},
			},
		},
	}

	// 执行 Elasticsearch 查询
	res, err := l.svcCtx.Es.Search(
		l.svcCtx.Es.Search.WithContext(context.Background()),
		l.svcCtx.Es.Search.WithIndex(in.TypeMount+"files"),
		l.svcCtx.Es.Search.WithBody(esutil.NewJSONReader(query)),
		l.svcCtx.Es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return &pb.SearchForFileRankResp{
			Error: fmt.Sprintf("Error executing the search: %s", err),
		}, nil
	}
	defer res.Body.Close()

	// 解析 Elasticsearch 查询结果
	if res.IsError() {
		return &pb.SearchForFileRankResp{
			Error: fmt.Sprintf("Error response returned from Elasticsearch: %s", res.String()),
		}, nil
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return &pb.SearchForFileRankResp{
			Error: fmt.Sprintf("Error parsing the response body: %s", err),
		}, nil
	}

	files := result["aggregations"].(map[string]interface{})["files"].(map[string]interface{})["buckets"].([]interface{})
	var Files []*pb.File
	for i := 0; i < len(files); i++ {
		file := files[i].(map[string]interface{})
		fileData := file["file_data"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
		Files = append(Files, &pb.File{
			Title: fileData["title"].(string),
			Id:    fileData["id"].(string),
		})
	}

	filesJson, err := json.Marshal(Files)
	if err != nil {
		return nil, err
	}
	val, err := l.svcCtx.Redis.SetnxEx("files"+in.TypeMount+strconv.Itoa(int(in.Rank)), string(filesJson), 3600)
	if err != nil || !val {
		return nil, err
	}
	return &pb.SearchForFileRankResp{
		Files: Files,
		Error: "",
	}, nil
}
