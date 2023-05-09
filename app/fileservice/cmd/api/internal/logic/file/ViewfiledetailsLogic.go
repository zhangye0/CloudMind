package file

import (
	"CloudMind/app/fileservice/cmd/rpc/filecenter"
	"context"
	"github.com/jinzhu/copier"

	"CloudMind/app/fileservice/cmd/api/internal/svc"
	"CloudMind/app/fileservice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewfiledetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewfiledetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewfiledetailsLogic {
	return &ViewfiledetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewfiledetailsLogic) Viewfiledetails(req *types.FileDetailsReq) (resp *types.FileDetailsResp, err error) {

	FileId := req.Id
	fileDetailsResp, err := l.svcCtx.FileRpc.FileDetails(l.ctx, &filecenter.FileDetailsReq{
		Id: FileId,
	})

	if err != nil {
		return nil, err
	}

	var filedetails types.FileDetailsResp
	_ = copier.Copy(&filedetails, fileDetailsResp)

	return &filedetails, nil
}
