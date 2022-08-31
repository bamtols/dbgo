package cliTypes

type (
	MigrationJob     string
	MigrationJobList []MigrationJob
)

const (
	MigrationJobAdd   MigrationJob = "Add"
	MigrationJobAlter MigrationJob = "Alter"
	MigrationJobDrop  MigrationJob = "Drop"
)

var (
	MigrationJobAll = MigrationJobList{
		MigrationJobAdd,
		MigrationJobAlter,
		MigrationJobDrop,
	}
)

func (x *MigrationJob) String() string {
	return string(*x)
}

func (x *MigrationJob) Valid() bool {
	for _, job := range MigrationJobAll {
		if job == *x {
			return true
		}
	}
	return false
}
