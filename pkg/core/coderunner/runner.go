package coderunner

import "fmt"

type Runner struct {
	id  string
	job Job
}

func (r *Runner) Execute(job Job) {
	r.id = "TODO: generate"
	r.job = job

	fmt.Println("Running job...")
}
