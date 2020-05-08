package schedule

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// ConfigureSchedule sets up the job that will be run
func ConfigureSchedule(f func()) {
	c := cron.New()

	c.AddFunc("@every 10m", f)

	dt := time.Now()

	startingJob := fmt.Sprint(dt.Format("Jan-02-06 03:04:05 PM"), "\n", "Starting cron job")

	fmt.Println(startingJob)
	c.Start()
}
