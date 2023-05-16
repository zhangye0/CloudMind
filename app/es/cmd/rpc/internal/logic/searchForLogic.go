package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchForLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchForLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchForLogic {
	return &SearchForLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
参数： 表名(string), 搜索内容(string), 搜索模式(string)

返回: 搜索到的数量

tips： 分为精确搜索(match)和模糊搜索(fuzz)
*/
func (l *SearchForLogic) SearchFor(in *pb.SearchForReq) (*pb.SearchForResp, error) {
	var buf bytes.Buffer
	//精确搜索
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"fuzzy": map[string]interface{}{
				"content": map[string]interface{}{
					"value":     in.Content,
					"fuzziness": "AUTO",
				},
			},
		},
	}

	// 序列化
	err := json.NewEncoder(&buf).Encode(query)
	if err != nil {
		return &pb.SearchForResp{
			Error: fmt.Sprintf("序列化失败"),
		}, nil
	}

	// 搜索
	res, err := l.svcCtx.Es.Search(
		l.svcCtx.Es.Search.WithContext(context.Background()),
		l.svcCtx.Es.Search.WithIndex(in.Index),
		l.svcCtx.Es.Search.WithBody(&buf),
		l.svcCtx.Es.Search.WithTrackTotalHits(true),
		l.svcCtx.Es.Search.WithPretty(),
	)
	if err != nil {
		return &pb.SearchForResp{
			Error: fmt.Sprintf("search error"),
		}, nil
	}
	defer res.Body.Close()
	if res.IsError() {
		return &pb.SearchForResp{
			Error: fmt.Sprintf("index not found"),
		}, nil
	}

	//反序列化搜索结果
	var r map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&r)

	var Sources []*pb.Source
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]
		t := pb.Source{
			Title:  source.(map[string]interface{})["title"].(string),
			Id:     source.(map[string]interface{})["id"].(int64),
			Avatar: source.(map[string]interface{})["avatar"].(string),
		}
		Sources = append(Sources, &t)
	}

	return &pb.SearchForResp{
		Sources: Sources,
		Error:   "",
	}, nil
}
