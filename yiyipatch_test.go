package yiyipatch

import (
	"testing"
	"runtime"
	"strconv"
	"time"
	"fmt"
)

func TestNewDispatcher(t *testing.T) {
	dispatcher := NewDispatcher((runtime.NumCPU() * 2) * 4)
	dispatcher.Run()
	start:= time.Now()
	for i := 0; i < 1000; i++ {
		var job Job
		job.key = []byte("job-" + strconv.Itoa(i))
		dispatcher.JobQueue <- job
	}
	end := time.Now().Sub(start)
	fmt.Println(end)

	//wait for sync print
	time.Sleep(5*time.Second)
}