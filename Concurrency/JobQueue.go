package main

import (
	"fmt"
	"sync"
)

type Job struct {
	Id int
}

type JobQueue struct {
	JobCh chan Job
	Wg    sync.WaitGroup
}

func NewJobQueue(capacity int) *JobQueue {
	jq := &JobQueue{
		JobCh: make(chan Job, capacity),
	}
	return jq
}
func (jq *JobQueue) Worker() {
	for j := range jq.JobCh {
		fmt.Println(j.Id)
		jq.Wg.Done()
	}
}

func (jq *JobQueue) AddJob(job Job) {
	jq.Wg.Add(1)
	jq.JobCh <- job

}

func (jq *JobQueue) Wait() {
	jq.Wg.Wait()
	close(jq.JobCh)
}

func main() {
	jq := NewJobQueue(100)
	go jq.Worker()
	for i := 0; i < 20; i++ {
		jq.AddJob(Job{Id: i})
	}
	jq.Wait()
}
