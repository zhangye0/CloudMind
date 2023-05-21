package logic

import (
	"CloudMind/app/mqueue/cmd/job/internal/svc"
	"CloudMind/app/mqueue/cmd/job/jobtype"
	"context"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//scheduler job
	mux.Handle(jobtype.ScheduleAddFlow, NewAddFlowHandler(l.svcCtx))

	//defer job
	//mux.Handle(jobtype.DeferCloseHomestayOrder,NewCloseHomestayOrderHandler(l.svcCtx))

	//queue job , asynq support queue job
	// wait you fill..

	return mux
}
