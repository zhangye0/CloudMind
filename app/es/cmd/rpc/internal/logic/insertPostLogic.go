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

type InsertPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertPostLogic {
	return &InsertPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/*
用户浏览/收藏/点赞了一个帖子
TypeMount: look(浏览), star(收藏), like(点赞)
UserId: 用户id
File: title(帖子名), content(帖子内容), id(帖子id)
*/
func (l *InsertPostLogic) InsertPost(in *pb.InsertPostReq) (*pb.InsertPostResp, error) {
	// 序列化
	data, err := json.Marshal(struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		Id        string `json:"id"`
		UserId    int64  `json:"userId"`
		TypeMount string `json:"typeMount"`
	}{Title: in.Post.Title,
		Content:   in.Post.Content,
		Id:        in.Post.Id,
		UserId:    in.UserId,
		TypeMount: in.TypeMount,
	})
	if err != nil {
		return &pb.InsertPostResp{
			Error: fmt.Sprintf("Error marshaling the document: %s", err),
		}, nil
	}
	// 构建请求

	req := esapi.IndexRequest{
		Index:   in.TypeMount + "posts",
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	// 发请求
	res, err := req.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		return &pb.InsertPostResp{
			Error: fmt.Sprintf("Error indexing the document: %s", err),
		}, nil
	}
	defer res.Body.Close()

	if res.IsError() {
		return &pb.InsertPostResp{
			Error: fmt.Sprintf("[%s] Error indexing document ID", res.Status()),
		}, nil
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			return &pb.InsertPostResp{
				Error: fmt.Sprintf("Error parsing the response body: %s", err),
			}, nil
		}
	}

	return &pb.InsertPostResp{}, nil
}
