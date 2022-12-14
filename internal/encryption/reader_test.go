package encryption

import (
	"github.com/TwoWaySix/enigma/internal"
	"os"
	"path"
	"path/filepath"
	"testing"
)

var testataRoot = "./testdata"

func TestReader(t *testing.T) {
	// Arrange Directories
	testdataDir := path.Join(testataRoot, "one")
	testdataOnePath := path.Join(testdataDir, "one.txt")
	testdataSubDir := path.Join(testdataDir, "two")
	testdataTwoPath := path.Join(testdataSubDir, "two.txt")
	err := os.MkdirAll(testdataSubDir, 0755)
	if err != nil {
		t.Errorf("%v", err)
	}
	f1, err := os.Create(testdataOnePath)
	if err != nil {
		t.Errorf("%v", err)
	}
	f1.Close()
	f2, err := os.Create(testdataTwoPath)
	if err != nil {
		t.Errorf("%v", err)
	}
	f2.Close()

	// Arrange Reader
	readFiles := make(chan internal.UnencryptedFile)
	reader := NewReader(testdataDir, readFiles)

	// Act
	err = reader.Start()
	if err != nil {
		t.Errorf("%v", err)
	}

	// Assert
	file1 := <-readFiles
	want, err := filepath.Rel(testdataDir, testdataOnePath)
	if err != nil {
		t.Errorf("%v", err)
	}
	if file1.Path != want {
		t.Errorf("got: %s, wanted: %s", file1.Path, testdataOnePath)
	}

	file2 := <-readFiles
	want, err = filepath.Rel(testdataDir, testdataTwoPath)
	if err != nil {
		t.Errorf("%v", err)
	}
	if file2.Path != want {
		t.Errorf("got: '%s', wanted: '%s'", file2.Path, want)
	}

	// Cleanup
	err = os.RemoveAll(testdataDir)
	if err != nil {
		t.Errorf("%v", err)
	}
}
