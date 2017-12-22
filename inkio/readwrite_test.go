package inkio

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestReadFileToStringValidTextFile(t *testing.T) {
	readpath := filepath.Join("..", "testfiles", "simpletext.txt")
	readstring, _ := ReadFileToString(readpath)
	if readstring != "This is simple text" {
		t.Errorf("[FAIL] ReadFileToString expected to return 'This is simple text' and instead returned '%s'", readstring)
	}
}

func TestReadFileToStringInvalidTextFile(t *testing.T) {
	readpath := filepath.Join("..", "testfiles", "totallybogus.txt")
	readstring, err := ReadFileToString(readpath)
	if err == nil {
		t.Errorf("[FAIL] Expected ReadFileToString to return error for missing file, instead it returned nil")
	}
	if len(readstring) > 0 {
		t.Errorf("[FAIL] Expected ReadFileToString to return empty string for missing file, instead it returned '%s'", readstring)
	}
}

func TestWriteStringToStdOut(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)

	// capture stdout stream in go routine
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	teststring := "this is a test"
	WriteString("testpath", true, &teststring)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if out != "this is a test" {
		t.Errorf("[FAIL] Expected WriteString to return 'this is a test' to standard output stream when stdout requested, however the function returned '%s'", out)
	}
}

func TestWriteStringToFile(t *testing.T) {
	mockTemplatePath := "testing.txt.in"
	mockOutPath := "testing.txt"
	teststring := "this is a test"
	err := WriteString(mockTemplatePath, false, &teststring)
	if err != nil {
		t.Errorf("[FAIL] There was an error with the file write for the TestWriteStringToFile test: %v", err)
	}

	if _, err := os.Stat(mockOutPath); !os.IsNotExist(err) {

		readstring, err := ioutil.ReadFile(mockOutPath)
		if err != nil {
			t.Errorf("[FAIL] Unable to read expected text file %s in TestWriteStringToFile test. %v", mockOutPath, err)
		}
		if string(readstring) != teststring {
			t.Errorf("[FAIL] Expected to read '%s' from test file and actually read '%s'", teststring, readstring)
		}
		os.Remove(mockOutPath)
	} else {
		t.Errorf("[FAIL] The expected file write for the TestWriteStringToFile test was not found. %v", err)
	}
}
