package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

func (l *AddTextLogic) AddText(in *pb.AddTextReq) (*pb.AddTextResp, error) {
	// 封装添加请求的字段
	data, err := json.Marshal(struct {
		Context string `json:"context"`
	}{Context: in.Text})

	if err != nil {
		return nil, errors.New("封装请求字段失败")
	}
	req := esapi.IndexRequest{
		Index:   in.Index,
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		return nil, errors.New("发送增加请求失败")
	}
	defer res.Body.Close()

	return &pb.AddTextResp{}, nil
}
