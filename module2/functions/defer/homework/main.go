package main

import (
	"errors"
	"fmt"
)

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {
	job.Merged = make(map[string]string)
	job.IsFinished = true
	for _, m := range job.Dicts {
		for k, v := range m {
			job.Merged[k] = v
		}

	}
	for _, m := range job.Dicts {
		if m == nil {
			return job, errNilDict
		}
	}

	if len(job.Merged) < 2 {
		return job, errNotEnoughDicts
	}

	return job, nil
}

func main() {

	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{}))
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, nil}}))
	fmt.Println(ExecuteMergeDictsJob(&MergeDictsJob{Dicts: []map[string]string{{"a": "b"}, {"b": "c"}}}))
}
