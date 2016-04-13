package main

//go:generate gen

import (
	"time"
)

// +gen slice:"SortBy,Where"
type Process struct {
	Pid  int
	Ppid int
	Name string
	Fds  FdSlice
}

// +gen slice:"SortBy,Where"
type Fd struct {
	Id int
}

func fetchCurrentProcesses() (ProcessSlice, error) {
	defer leave(enter("fetchCurrentProcesses"))

	// simulate long blocking fetch operation
	time.Sleep(time.Duration(5) * time.Second)

	// TODO do the real thing

	// return something
	return ProcessSlice{
		{Pid: 1, Ppid: 1, Name: "httpd", Fds: FdSlice{{Id: 1}, {Id: 2}}},
		{Pid: 2, Ppid: 1, Name: "sshd", Fds: FdSlice{{Id: 3}, {Id: 4}}},
	}, nil
}
