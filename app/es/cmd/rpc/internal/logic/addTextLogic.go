package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddTextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTextLogic {
	return &AddTextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
参数: 表名(string), 记录(string)

返回: 错误信息(string), 错误error
*/
func (l *AddTextLogic) AddText(in *pb.AddTextReq) (*pb.AddTextResp, error) {
	// 序列化
	data, err := json.Marshal(struct {
		Context string `json:"context"`
	}{Context: in.Text})
	if err != nil {
		return &pb.AddTextResp{
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
		return &pb.AddTextResp{
			Error: fmt.Sprintf("do Error"),
		}, nil
	}
	defer res.Body.Close()

	if res.IsError() {
		return &pb.AddTextResp{
			Error: fmt.Sprintf("[%s] Error indexing document ID", res.Status()),
		}, nil
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			return &pb.AddTextResp{
				Error: fmt.Sprintf("Error parsing the response body: %s", err),
			}, nil
		}
	}

	return &pb.AddTextResp{}, nil
}
