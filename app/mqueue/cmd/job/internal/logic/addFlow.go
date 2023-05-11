package logic

import (
	"CloudMind/app/mqueue/cmd/job/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"context"
	"fmt"
	"github.com/hibiken/asynq"
)

// SettleRecordHandler   shcedule billing to home business
type AddFlowHandler struct {
	svcCtx *svc.ServiceContext
}

func NewAddFlowHandler(svcCtx *svc.ServiceContext) *AddFlowHandler {
	return &AddFlowHandler{
		svcCtx: svcCtx,
	}
}

// every one minute exec : if return err != nil , asynq will retry
func (l *AddFlowHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	_, err := l.svcCtx.UsercenterRpc.UpdateUserInfo(ctx, &pb.UpdateUserInfoReq{
		UpdateType: "AllFlow",
		Field1:     "",
		Filed2:     "",
		Filed3:     0,
		Filed4:     100.0,
	})
	if err != nil {
		return err
	}
	fmt.Printf("流量增加成功!!!\n")
	return nil
}
