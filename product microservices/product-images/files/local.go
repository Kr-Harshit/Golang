package files

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//Local is an implementation of the storage which works with the
//local disk on the current machine
type Local struct {
	maxFileSize int // maximum number of bytes of files
	basePath    string
}

// NewLocal creates a new Local Filesystem with the given base path
// basepath is tha base directory to save the file to.
// maxSize is the max number of bytes that a file can be
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p, maxFileSize: maxSize}, nil
}

// returns the absolute path
func (l *Local) fullPath(path string) string {
	// append the given path to the base path
	return filepath.Join(l.basePath, path)
}

// Save the contents pf the Writer to the given path
// path is a relative path, basePath will be appended
func (l *Local) Save(path string, contents io.Reader) error {
	// get the full path for the file
	fp := l.fullPath(path)

	// get the directory and make sure it exists
	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return fmt.Errorf("unable to create directory: %w", err)
	}

	// if the file exist delete it
	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return fmt.Errorf("unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		// if this is anything other than a not exists error
		return fmt.Errorf("unable to get file info: %w", err)
	}

	// create a new file at the path
	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("unable to Create file: %w", err)
	}
	defer f.Close()

	// write the content to the new file
	// ensure that we are not writing greater than max bytes
	_, err = io.Copy(f, contents)
	if err != nil {
		return fmt.Errorf("unable to write to file: %w", err)
	}
	return nil
}

// Get the file at the given path and return a Reader
// the calling function is Responsible for clossing the Reader
func (l *Local) Get(Path string) (*os.File, error) {
	// get the full path for the file
	fp := l.fullPath(Path)

	// open the file
	f, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("unable to Open the file: %w", err)
	}
	return f, nil
}
