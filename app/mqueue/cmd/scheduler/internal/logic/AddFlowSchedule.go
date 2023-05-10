package logic

import (
	"CloudMind/app/mqueue/cmd/job/jobtype"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

// scheduler job ------> go-zero-looklook/app/mqueue/cmd/job/internal/logic/AddFlowSchedule.go.
func (l *MqueueScheduler) AddFlowScheduler() {

	task := asynq.NewTask(jobtype.ScheduleAddFlow, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【AddFlowScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【AddFlowScheduler】 registered an  entry: %q \n", entryID)
}
