package organizer

import (
	"fmt"
	"os"
	"path/filepath"
)

func Organize(Dirpath string) error {

	directory, err := getFiles(Dirpath)
	if err != nil {
		return fmt.Errorf("Error while reading Directory %v", err)
	}
	for i := 0; i < len(directory); i++ {
		if directory[i].IsDir() {
			Organize(filepath.Join(Dirpath, directory[i].Name()))
		}
		ext := filepath.Ext(directory[i].Name())
		if ext != "" {
			ext = ext[1:]
		}
		err := os.MkdirAll(filepath.Join(Dirpath, ext), 0777)
		if err != nil {
			return fmt.Errorf("Error while creating Directory %v", err)
		}
		if ext != "" {
			oldfile := filepath.Join(Dirpath, directory[i].Name())
			newfile := filepath.Join(Dirpath, ext, directory[i].Name())
			err = os.Rename(oldfile, newfile)
		}
		if err != nil {
			return fmt.Errorf("Error while shifting file %v", err)
		}
	}
	return nil
}

func getFiles(Dirpath string) ([]os.DirEntry, error) {
	directory, err := os.ReadDir(Dirpath)
	if err != nil {
		return nil, err
	}
	return directory, err
}

//
//folder
//file1.go file2.go folder
//      			  file3.go file4.go
