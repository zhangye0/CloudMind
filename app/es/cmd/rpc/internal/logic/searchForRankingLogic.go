package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForRankingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForRankingLogic {
	return &SearchForRankingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchForRankingLogic) SearchForRanking(in *pb.SearchForRankingReq) (*pb.SearchForRankingResp, error) {
	query := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"files": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "id",
					"size":  in.Ranking,
				},
				"aggs": map[string]interface{}{
					"file_data": map[string]interface{}{
						"top_hits": map[string]interface{}{
							"size": in.Ranking,
							"_source": map[string]interface{}{
								"includes": []string{"title", "id", "avatar"},
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
		l.svcCtx.Es.Search.WithIndex(in.Index),
		l.svcCtx.Es.Search.WithBody(esutil.NewJSONReader(query)),
		l.svcCtx.Es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return &pb.SearchForRankingResp{
			Sources: nil,
			Error:   fmt.Sprintf("Error executing the search: %s", err),
		}, nil
	}
	defer res.Body.Close()

	// 解析 Elasticsearch 查询结果
	if res.IsError() {
		return &pb.SearchForRankingResp{
			Sources: nil,
			Error:   fmt.Sprintf("Error response returned from Elasticsearch: %s", res.String()),
		}, nil
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return &pb.SearchForRankingResp{
			Sources: nil,
			Error:   fmt.Sprintf("Error parsing the response body: %s", err),
		}, nil
	}

	files := result["aggregations"].(map[string]interface{})["files"].(map[string]interface{})["buckets"].([]interface{})
	var sources []*pb.Source
	for i := 0; i < len(files); i++ {
		file := files[i].(map[string]interface{})
		fileData := file["file_data"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
		sources = append(sources, &pb.Source{
			Title:  fileData["title"].(string),
			Id:     int64(fileData["id"].(float64)),
			Avatar: fileData["avatar"].(string),
		})
	}

	return &pb.SearchForRankingResp{
		Sources: sources,
		Error:   "",
	}, nil
}
