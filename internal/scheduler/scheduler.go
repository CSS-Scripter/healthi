package scheduler

import "time"

func RunTask(errorChan chan error, task func() error, interval int) {
	err := task()
	if err != nil {
		errorChan <- err
	}
	time.Sleep(time.Second * time.Duration(interval))
	RunTask(errorChan, task, interval)
}
