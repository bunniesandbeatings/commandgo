package commandgo

import (
	"io/ioutil"
	"os"
	"fmt"
)

type Disk struct {
	HandlesErrors
}

func NewDisk() *Disk {
	disk := &Disk{
		HandlesErrors: NewHandlesErrors(),
	}

	return disk
}

// Returns a temp file path, without an underlying system file
func (disk *Disk) CreateTempFilePath(prefix string) string {
	path := disk.CreateTempFile(prefix, []byte{})
	if err := os.Remove(path); err != nil {
		disk.ErrorHandler(err)
	}
	return path
}

// Returns a temp file, that is empty on disk, or contains the bytes you pass in
func (disk *Disk) CreateTempFile(prefix string, contents []byte) string {
	outputFile, err := ioutil.TempFile("", prefix)
	if err != nil {
		disk.ErrorHandler(err)
	}

	if len(contents) > 0 {
		fmt.Fprint(outputFile, contents)
	}

	if closeError := outputFile.Close(); closeError != nil {
		disk.ErrorHandler(closeError)
	}

	return outputFile.Name()
}
