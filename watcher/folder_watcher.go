package watcher

import (
	"time"
)

type FolderWatcher struct {
	path               string
	includeHiddenFiles bool
	includeSubfolders  bool
	interval           time.Duration
	lastKnownState     *State
}

func NewFolderWatcher(path string, hiddenFiles, subfolders bool) *FolderWatcher {
	return &FolderWatcher{
		path:               path,
		includeHiddenFiles: hiddenFiles,
		includeSubfolders:  subfolders,
		interval:           1 * time.Second,
	}
}

func (f *FolderWatcher) StartWatcher() {
	for {
		f.watch()
		time.Sleep(f.interval)
	}
}

func (f *FolderWatcher) watch() {
	// get current state of directory
	// check against old state
	// figure out differences
	// save differences to db
	// check against cloud hashes
	// upload file and/or metadata
	// delete files and/or metadata
	diff := f.computeStateDiff()
	if diff != nil {

	}
}

func (f *FolderWatcher) computeStateDiff() *State {
	state := f.getCurrentState()
	if f.lastKnownState == nil {
		f.lastKnownState = state
	}

}

func (f *FolderWatcher) getCurrentState() *State {
	return &State{}
}
