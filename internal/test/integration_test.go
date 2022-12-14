package test

import (
	"github.com/TwoWaySix/enigma/internal"
	"github.com/TwoWaySix/enigma/internal/decryption"
	"github.com/TwoWaySix/enigma/internal/encryption"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var inputDir = "./testdata/input"
var inputFile1Path = filepath.Join(inputDir, "file1.txt")
var inputFile1Content = "ENIGMA"
var encryptionDir = "./testdata/encrypted"
var encryptedFilePath = filepath.Join(encryptionDir, "file.zip")
var decryptionDir = "./testdata/decrypted"
var decryptionFilePath = filepath.Join(decryptionDir, "file1.txt")
var key = strings.Repeat("a", 16)

func TestMain(m *testing.M) {
	err := os.MkdirAll(inputDir, 0755)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(encryptionDir, 0755)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(decryptionDir, 0755)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(inputFile1Path)
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(inputFile1Content)
	if err != nil {
		panic(err)
	}

	m.Run()

	err = os.Remove(inputFile1Path)
	if err != nil {
		panic(err)
	}
	err = os.Remove(encryptedFilePath)
	if err != nil {
		panic(err)
	}
	err = os.Remove(decryptionFilePath)
	if err != nil {
		panic(err)
	}
}

func TestEncryptionDecryption(t *testing.T) {
	// ENCRYPTION
	config := internal.Config{
		Mode:       "roll",
		InputPath:  inputDir,
		OutputPath: encryptedFilePath,
		Key:        key,
	}

	encryptionJob, err := encryption.NewJob(config)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = encryptionJob.Start()
	if err != nil {
		t.Errorf(err.Error())
	}

	// DECRYPTION
	config.Mode = "unroll"
	config.InputPath = encryptedFilePath
	config.OutputPath = decryptionDir

	decryptionJob, err := decryption.NewJob(config)
	if err != nil {
		t.Errorf(err.Error())
	}
	err = decryptionJob.Start()
	if err != nil {
		t.Errorf(err.Error())
	}

	// Check if after encryption and decryption the file content still is the same
	data, err := os.ReadFile(decryptionFilePath)
	if err != nil {
		t.Errorf(err.Error())
	}
	got := string(data)
	if got != inputFile1Content {
		t.Errorf("got: %s , want: %s", got, inputFile1Content)
	}
}
