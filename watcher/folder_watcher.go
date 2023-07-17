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
	f.computeStateDiff()
}

func (f *FolderWatcher) computeStateDiff() {
	state := f.getCurrentState()
	if f.lastKnownState == nil {
		f.lastKnownState = state
	}
}

func (f *FolderWatcher) getCurrentState() *State {
	files := f.getFiles()
	folders := f.getFolders()
	return &State{
		files:   files,
		folders: folders,
	}
}

func (f *FolderWatcher) getFiles() []*File {
	return []*File{}
}

func (f *FolderWatcher) getFolders() []*Folder {
	return []*Folder{}
}
