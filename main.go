package main

import (
	"github.com/aldarisbm/seratobackup/watcher"
)

func main() {
	folderWatcher := watcher.NewFolderWatcher(
		"/Users/berrio/Development/seratobackup/test_folder",
		true,
		true,
	)

	folderWatcher.StartWatcher()
}
