package logic

import (
	"CloudMind/app/es/cmd/rpc/internal/svc"
	"CloudMind/app/es/cmd/rpc/pb"
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFilesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFilesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFilesLogic {
	return &UpdateFilesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFilesLogic) UpdateFiles(in *pb.UpdateFilesReq) (*pb.UpdateFilesResp, error) {
	requestBody := map[string]interface{}{
		"query": map[string]interface{}{
			"constant_score": map[string]interface{}{
				"filter": map[string]interface{}{
					"term": map[string]interface{}{
						"id": in.File.Id,
					},
				},
			},
		},
		"script": map[string]interface{}{
			"source": "ctx._source.title = params.new_title;",
			"params": map[string]interface{}{
				"new_title": in.File.Title,
			},
		},
	}

	// Create the Update By Query request object
	req := esapi.UpdateByQueryRequest{
		Index: []string{"uploadfiles", "downloadfiles", "starfiles", "likefiles"},
		Body:  esutil.NewJSONReader(requestBody),
	}

	// Send the Update By Query request
	res, err := req.Do(context.Background(), l.svcCtx.Es)
	if err != nil {
		return &pb.UpdateFilesResp{
			Error: fmt.Sprintf("Error executing the Do: %s", err),
		}, nil
	}
	defer res.Body.Close()

	if res.IsError() && res.StatusCode != 404 {
		return &pb.UpdateFilesResp{
			Error: fmt.Sprintf("update by query failed with code %d: %s", res.StatusCode, res.String()),
		}, nil
	}

	return &pb.UpdateFilesResp{}, nil
}
