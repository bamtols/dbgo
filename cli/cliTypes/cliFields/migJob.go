package cliFields

import "github.com/bamtols/dbgo/func/parserFn"

type (
	MigJob     string
	MigJobList []MigJob
)

const (
	MigJobAdd   MigJob = "Add"
	MigJobDrop  MigJob = "Drop"
	MigJobAlter MigJob = "Alter"
)

var (
	MigJobAll = MigJobList{
		MigJobAdd,
		MigJobDrop,
		MigJobAlter,
	}
)

func (x *MigJob) String() string {
	return parserFn.GoPublicCamelCase(string(*x))
}

func (x *MigJob) Valid() bool {
	for _, job := range MigJobAll {
		if x.String() == job.String() {
			return true
		}
	}
	return false
}
