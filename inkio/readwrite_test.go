package inkio

import (
	"bytes"
	"io"
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
