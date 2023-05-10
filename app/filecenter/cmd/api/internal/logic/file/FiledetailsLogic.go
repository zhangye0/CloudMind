package file

import (
	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"
	"CloudMind/app/filecenter/cmd/rpc/filecenter"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type FiledetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFiledetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FiledetailsLogic {
	return &FiledetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FiledetailsLogic) Filedetails(req *types.FileDetailsReq) (*types.FileDetailsResp, error) {

	FileId := req.Id
	x, err := l.svcCtx.FileRpc.FileDetails(l.ctx, &filecenter.FileDetailsReq{
		Id: FileId,
	})

	if err != nil {
		return nil, err
	}

	var filedetails types.FileDetailsResp
	_ = copier.Copy(&filedetails, x)

	return &filedetails, nil
}
