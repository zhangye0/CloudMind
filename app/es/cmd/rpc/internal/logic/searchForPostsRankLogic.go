package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esutil"

	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForPostsRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForPostsRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForPostsRankLogic {
	return &SearchForPostsRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchForPostsRankLogic) SearchForPostsRank(in *pb.SearchForPostsRankReq) (*pb.SearchForPostsRankResp, error) {
	//var gte, lte string
	//lte = "now/d"
	//switch in.TypeTime {
	//case "day":
	//	gte = "now-1d/d"
	//case "week":
	//	gte = "now-7d/d"
	//case "month":
	//	gte = "now-30d/d"
	//case "year":
	//	gte = "now-365d/d"
	//}
	query := map[string]interface{}{
		"size": 0,
		//"query": map[string]interface{}{
		//	"range": map[string]interface{}{
		//		"timestamp": map[string]interface{}{
		//			"gte": gte,
		//			"lte": lte,
		//		},
		//	},
		//},
		"aggs": map[string]interface{}{
			"files": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "id.keyword",
					"size":  in.Ranking,
				},
				"aggs": map[string]interface{}{
					"file_data": map[string]interface{}{
						"top_hits": map[string]interface{}{
							"size": in.Ranking,
							"_source": map[string]interface{}{
								"includes": []string{"title", "content", "id"},
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
		l.svcCtx.Es.Search.WithIndex("posts"),
		l.svcCtx.Es.Search.WithBody(esutil.NewJSONReader(query)),
		l.svcCtx.Es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return &pb.SearchForPostsRankResp{
			Error: fmt.Sprintf("Error executing the search: %s", err),
		}, nil
	}
	defer res.Body.Close()

	// 解析 Elasticsearch 查询结果
	if res.IsError() {
		return &pb.SearchForPostsRankResp{
			Error: fmt.Sprintf("Error response returned from Elasticsearch: %s", res.String()),
		}, nil
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return &pb.SearchForPostsRankResp{
			Error: fmt.Sprintf("Error parsing the response body: %s", err),
		}, nil
	}

	files := result["aggregations"].(map[string]interface{})["files"].(map[string]interface{})["buckets"].([]interface{})
	var Posts []*pb.Post
	for i := 0; i < len(files); i++ {
		file := files[i].(map[string]interface{})
		fileData := file["file_data"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
		Posts = append(Posts, &pb.Post{
			Title:   fileData["title"].(string),
			Content: fileData["content"].(string),
			Id:      fileData["id"].(string),
		})
	}

	return &pb.SearchForPostsRankResp{
		Posts: Posts,
		Error: "",
	}, nil
}
