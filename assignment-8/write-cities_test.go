package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

const fakeFileName string = "file-write-test.txt"

func TestWriteFileToDisk(t *testing.T) {
	t.Run("writes file to disk", func(t *testing.T) {
		fakeDirectory := writeFakeDirectory(t)
		defer os.RemoveAll(fakeDirectory)

		fakeFilePath := filepath.Join(fakeDirectory, fakeFileName)
		sut := []string{"this", "is", "some", "test", "text"}

		WriteFileToDisk(sut, fakeFilePath)

		checkFileExists(fakeFilePath)
	})
}

func TestReadFileFromDisk(t *testing.T) {
	t.Run("reads file contents", func(t *testing.T) {
		fakeDirectory := writeFakeDirectory(t)
		fmt.Println(fakeDirectory)
		defer os.RemoveAll(fakeDirectory)

		want := "this \n is \n some \n test \n text"
		fakeFilePath := filepath.Join(fakeDirectory, fakeFileName)
		fmt.Println(fakeFilePath)
		fakeFile := writeFakeFileToDirectory(t, want, fakeDirectory, fakeFileName)
		defer fakeFile.Close()

		got := ReadFileFromDisk(fakeFile.Name())

		assertCorrectMessage(t, got, want)
	})
}

func writeFakeFileToDirectory(t testing.TB, contents string, filepath string, filename string) os.File {
	t.Helper()

	tempFile, err := os.CreateTemp(filepath, filename)

	if err != nil {
		log.Fatal(err)
	}

	tempFile.Write([]byte(contents))
	return *tempFile
}

func checkFileExists(filepath string) bool {
	_, err := os.Stat(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return true
}

func writeFakeDirectory(t testing.TB) string {
	t.Helper()
	fakeDirectory, err := os.MkdirTemp("", "testdata")
	if err != nil {
		log.Fatal(err)
	}

	return fakeDirectory
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
