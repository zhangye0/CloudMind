package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertLogic {
	return &InsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
参数: 表名(string), 记录(Source)
返回: 错误信息(string)
*/
func (l *InsertLogic) Insert(in *pb.InsertReq) (*pb.InsertResp, error) {
	// 序列化
	data, err := json.Marshal(struct {
		Title  string `json:"title"`
		Id     int64  `json:"id"`
		Avatar string `json:"avatar"`
	}{Title: in.Source.Title,
		Id:     in.Source.Id,
		Avatar: in.Source.Avatar,
	})
	if err != nil {
		return &pb.InsertResp{
			Error: fmt.Sprintf("json Marshal error"),
		}, nil
	}
	// 构建请求

	req := esapi.IndexRequest{
		Index:   in.Index,
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	// 发请求
	res, err := req.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		return &pb.InsertResp{
			Error: fmt.Sprintf("do Error"),
		}, nil
	}
	defer res.Body.Close()

	if res.IsError() {
		return &pb.InsertResp{
			Error: fmt.Sprintf("[%s] Error indexing document ID", res.Status()),
		}, nil
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			return &pb.InsertResp{
				Error: fmt.Sprintf("Error parsing the response body: %s", err),
			}, nil
		}
	}

	return &pb.InsertResp{}, nil
}
