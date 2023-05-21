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

type InsertFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertFileLogic {
	return &InsertFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
用户下载/收藏/点赞了一个文件
TypeMount: download(下载), star(收藏), like(点赞), upload(上传)
UserId: 用户id
File: title(文件名), id(文件id)
*/
func (l *InsertFileLogic) InsertFile(in *pb.InsertFileReq) (*pb.InsertFileResp, error) {
	// 序列化
	data, err := json.Marshal(struct {
		Title     string `json:"title"`
		Id        string `json:"id"`
		UserId    int64  `json:"userId"`
		TypeMount string `json:"typeMount"`
	}{Title: in.File.Title,
		Id:        in.File.Id,
		UserId:    in.UserId,
		TypeMount: in.TypeMount,
	})
	if err != nil {
		return &pb.InsertFileResp{
			Error: fmt.Sprintf("Error marshaling the document: %s", err),
		}, nil
	}
	// 构建请求

	req := esapi.IndexRequest{
		Index: in.TypeMount + "files",
		Body:  bytes.NewReader(data),
	}

	// 发请求
	res, err := req.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		return &pb.InsertFileResp{
			Error: fmt.Sprintf("Error indexing the document: %s", err),
		}, nil
	}
	defer res.Body.Close()

	if res.IsError() {
		return &pb.InsertFileResp{
			Error: fmt.Sprintf("[%s] Error indexing document ID", res.Status()),
		}, nil
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			return &pb.InsertFileResp{
				Error: fmt.Sprintf("Error parsing the response body: %s", err),
			}, nil
		}
	}

	return &pb.InsertFileResp{}, nil
}
