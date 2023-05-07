package file

import (
	"context"

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

	return
}
