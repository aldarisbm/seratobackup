package watcher

import "time"

type FolderWatcher struct {
	path               string
	includeHiddenFiles bool
	includeSubfolders  bool
	interval           time.Duration
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

}
