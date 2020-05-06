package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"reflect"
	"runtime"
	"time"
)

// 单例
var Corn *cron.Cron

func Run(job func() error) {
	from := time.Now().Unix()
	err := job()
	to := time.Now().UnixNano()
	jobname := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err == nil {
		fmt.Printf("%s success: %dms \n", jobname, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s error: %dms \n", jobname, (to-from)/int64(time.Millisecond))

	}
}

func CronJob() {
	if Corn == nil {
		Corn = cron.New()
	} else {
		Corn.AddFunc("0 0 0 * * *", func() {
			Run(RestartDailyRank)
		})
		Corn.Start()
		fmt.Println("CronJob Start.......")

	}
}
