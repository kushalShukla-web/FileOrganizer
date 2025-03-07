package organizer

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func createfile(a []string, b string) error {
	for i := range a {
		file, err := os.OpenFile(filepath.Join(b, a[i]), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("Error while Opening file %v", err)
		}
		defer file.Close()
		if err != nil {
			return fmt.Errorf("Error while closing the file %v", err)
		}
		_, err = file.WriteString("GREAT!!\n")
		if err != nil {
			return fmt.Errorf("Error while Writing in file %v", err)
		}
	}
	return nil
}
func TestOrganizer(t *testing.T) {
	testDir := "./testdir"
	err := os.RemoveAll(testDir)
	if err != nil {
		t.Fatalf("Failed to remove test directory: %v", err)
	}
	err = os.MkdirAll(testDir, 0775)
	if err != nil {
		t.Fatalf("failed to create test directory")
	}
	testdata := []string{
		"messi.go",
		"neymar.go",
	}

	err = createfile(testdata, testDir)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	err = Organize(testDir)
	if err != nil {
		t.Fatalf("Error %v", err)
	}
	for _, file := range testdata {
		ext := filepath.Ext(file)
		if ext == "" {
			ext = "no_extension"
		} else {
			ext = ext[1:]
		}

		destDir := filepath.Join(testDir, ext)
		destPath := filepath.Join(destDir, file)
		if _, err := os.Stat(destPath); os.IsNotExist(err) {
			t.Errorf("File %s was not moved to %s", file, destDir)
			t.Fatalf("Error: %v", err)
		}
	}
}
