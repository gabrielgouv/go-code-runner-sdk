package coderunner

type Job struct {
	Id    *string
	Lang  *string
	Code  *string
	Tests []*Test
}

type Test struct {
	Id          *string
	Description *string
	Input       *string
	Expect      *string
}

type JobResult struct {
	Job     *Job
	Success *bool
}
