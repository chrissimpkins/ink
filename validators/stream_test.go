package validators

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestStdinValidatesMockStdinWithData(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "stdin")
	defer os.Remove(file.Name())

	// Write to it
	file.WriteString("test data in standard input stream")
	result := StdinValidates(file)
	if result == false {
		t.Errorf("[FAIL] StdinValidates function tested with mocked standard input data returns false, expected true.")
	}
}

func TestStdinValidatesMockStdinWithoutData(t *testing.T) {
	file, _ := ioutil.TempFile(os.TempDir(), "stdin")
	defer os.Remove(file.Name())

	// Write to it
	file.WriteString("")
	result := StdinValidates(file)
	if result == true {
		t.Errorf("[FAIL] StdinValidates function tested with mocked empty standard input data and returns true, expected false.")
	}
}
