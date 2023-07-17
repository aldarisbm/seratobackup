package watcher

import "time"

type State struct {
	folders     []*Folder
	files       []*File
	lastChecked *time.Time
}

type Folder struct{}

type File struct{}
