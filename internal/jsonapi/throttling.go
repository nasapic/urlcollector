package jsonapi

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"gitlab.com/nasapic/urlcollector/internal/transport"
)

type (
	Queue struct {
		Rate       time.Duration
		JobChannel chan *Job
	}

	Job struct {
		id            uuid.UUID
		Request       *http.Request
		Func          SearchFunc
		SearchRequest *transport.SearchRequest
		DoneChannel   chan (bool)
		Result        *JobResult
		Error         error
	}

	JobResult struct {
		Result *transport.SearchResponse
		Error  error
	}
)

type (
	SearchFunc = func(*transport.SearchRequest) (*transport.SearchResponse, error)
)

func NewQueue(maxReqsPerSec int) *Queue {
	return &Queue{
		Rate:       time.Second / time.Duration(maxReqsPerSec),
		JobChannel: make(chan *Job, 1000),
	}
}

func (q *Queue) processJobs() {
	throttle := time.Tick(q.Rate)

	for job := range q.JobChannel {
		<-throttle
		go job.Process()
	}
}

func NewJob(r *http.Request, sf SearchFunc, sReq *transport.SearchRequest) *Job {
	return &Job{
		id:            uuid.New(),
		Request:       r,
		Func:          sf,
		SearchRequest: sReq,
		DoneChannel:   make(chan (bool)),
	}
}

func (job *Job) Process() {
	searchRes, err := job.Func(job.SearchRequest)

	select {
	case <-job.Request.Context().Done():
		return

	default:
		job.Result = &JobResult{
			Result: searchRes,
			Error:  err,
		}

		job.DoneChannel <- true
	}
}
