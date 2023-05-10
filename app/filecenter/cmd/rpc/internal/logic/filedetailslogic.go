package logic

import (
	"CloudMind/app/filecenter/model"
	"CloudMind/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDetailsLogic {
	return &FileDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDetailsLogic) FileDetails(in *pb.FileDetailsReq) (*pb.FileDetailsResp, error) {
	Filex, err := l.svcCtx.FileModel.FindOne(l.ctx, in.Id)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " HomestayDetail db err , id : %d ", in.Id)
	}

	l.Logger.Error("没错")

	var pbFile pb.FileDetailsResp
	if Filex != nil {
		_ = copier.Copy(&pbFile, Filex)
	}

	return &pbFile, nil
}
