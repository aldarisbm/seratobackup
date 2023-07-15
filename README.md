
The idea here is to be able to back up a folder recursively.

we will use a daemon to be able to run the backup in the background

we will save the metadata in a different table to 

we do this to avoid reuploading the same files only when metadata has changed

if metadata changes but the hash of the file is the same, we will not reupload the file

we will update the metadata table