package coderunner

type Subscriber interface {
	OnJobReceived(job Job)
	OnJobFinished(result JobResult)
}
