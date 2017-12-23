package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// test version string formatting
func TestVersionString(t *testing.T) {
	r, _ := regexp.Compile(`\d{1,2}.\d{1,2}.\d{1,2}`)
	if r.MatchString(Version) == false {
		t.Errorf("[FAIL] Failed to match regex pattern to version string %s", Version)
	}
}

// test usage string formatting
func TestUsageString(t *testing.T) {
	if strings.HasPrefix(Usage, "Usage:") == false {
		t.Errorf("[FAIL] Improperly formatted usage string.  Expected string to start with 'Usage:' and received %s", Usage)
	}
}

// test help string formatting
func TestHelpString(t *testing.T) {
	if strings.HasPrefix(Help, "====") == false {
		t.Errorf("[FAIL] Improperly formatted usage string. Expected to start with '===' and received %s", Help)
	}
}

// --------------------------------------------
// test default command line argument variables
// --------------------------------------------

func TestDefaultVersionShort(t *testing.T) {
	if *versionShort == true {
		t.Errorf("[FAIL] Expected *versionShort == false as default, got true")
	}
}

func TestDefaultVersionLong(t *testing.T) {
	if *versionLong == true {
		t.Errorf("[FAIL] Expected *versionLong == false as default, got true")
	}
}

func TestDefaultHelpShort(t *testing.T) {
	if *helpShort == true {
		t.Errorf("[FAIL] Expected *helpShort == false as default, got true")
	}
}

func TestDefaultHelpLong(t *testing.T) {
	if *helpLong == true {
		t.Errorf("[FAIL] Expected *helpLong == false as default, got true")
	}
}

func TestDefaultUsageLong(t *testing.T) {
	if *usageLong == true {
		t.Errorf("[FAIL] Expected *usageLong == false as default, got true")
	}
}

func TestDefaultFindString(t *testing.T) {
	if len(*findString) > 0 {
		t.Errorf("[FAIL] Expected empty *findString value by default, received string %s", *findString)
	}
}

func TestDefaultReplaceString(t *testing.T) {
	if len(*replaceString) > 0 {
		t.Errorf("[FAIL] Expected empty *replaceString value by default, received string %s", *replaceString)
	}
}

func TestDefaultLintFlag(t *testing.T) {
	if *lintFlag == true {
		t.Errorf("[FAIL] Expected *lintFlag == false as default, got true")
	}
}

func TestDefaultStdOutFlag(t *testing.T) {
	if *stdOutFlag == true {
		t.Errorf("[FAIL] Expected *stdOutFlag == false as default, got true")
	}
}

func TestDefaultTrimNLFlag(t *testing.T) {
	if *trimNLFlag == true {
		t.Errorf("[FAIL] Expected *trimNLFlag == false as default, got true")
	}
}

// --------------------------------------------
// test render functions
// --------------------------------------------

func TestRenderLocalBuiltinTemplateFileWrite(t *testing.T) {
	templatePath := filepath.Join("testfiles", "template_1.txt.in")
	outPath := filepath.Join("testfiles", "template_1.txt")
	replaceString := "test"
	expectedString := "sha=test test=test"
	mockStdoutFlag := false
	fileerr := renderLocal(templatePath, &replaceString, &mockStdoutFlag)

	_, staterr := os.Stat(outPath)
	if !os.IsNotExist(staterr) {

		readstring, readerr := ioutil.ReadFile(outPath)
		if readerr != nil {
			t.Errorf("[FAIL] Unable to read expected text file %s in test. %v", outPath, readerr)
		}
		if string(readstring) != expectedString {
			t.Errorf("[FAIL] Expected to read '%s' from test file and actually read '%s'", expectedString, readstring)
		}
		os.Remove(outPath)
	} else {
		t.Errorf("[FAIL] The expected file write for the test was not found. %v", fileerr)
	}
}

func TestRenderLocalUserTemplateFileWrite(t *testing.T) {
	templatePath := filepath.Join("testfiles", "template_3.txt.in")
	outPath := filepath.Join("testfiles", "template_3.txt")
	replaceString := "test"
	*findString = "[[user]]"
	expectedString := "sha=test test=test"
	mockStdoutFlag := false
	fileerr := renderLocal(templatePath, &replaceString, &mockStdoutFlag)
	*findString = "" // reset to default value or this interferes with other tests

	_, staterr := os.Stat(outPath)
	if !os.IsNotExist(staterr) {

		readstring, readerr := ioutil.ReadFile(outPath)
		if readerr != nil {
			t.Errorf("[FAIL] Unable to read expected text file %s in test. %v", outPath, readerr)
		}
		if string(readstring) != expectedString {
			t.Errorf("[FAIL] Expected to read '%s' from test file and actually read '%s'", expectedString, readstring)
		}
		os.Remove(outPath)
	} else {
		t.Errorf("[FAIL] The expected file write for the test was not found. %v", fileerr)
	}
}

func TestRenderLocalBuiltinTemplateStdout(t *testing.T) {
	templatePath := filepath.Join("testfiles", "template_1.txt.in")
	testString := "test"
	mockStdoutFlag := true
	expectedString := "sha=test test=test"

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

	fileerr := renderLocal(templatePath, &testString, &mockStdoutFlag)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if out != expectedString {
		t.Errorf("[FAIL] Expected test to return '%s' to standard output stream when stdout requested, however the function returned '%s'", expectedString, out)
	}
	if fileerr != nil {
		t.Errorf("[FAIL] Unexpected error raised during execution: %v", fileerr)
	}

}

func TestRenderLocalUserTemplateStdout(t *testing.T) {
	templatePath := filepath.Join("testfiles", "template_3.txt.in")
	testString := "test"
	*findString = "[[user]]"
	expectedString := "sha=test test=test"
	mockStdoutFlag := true

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

	fileerr := renderLocal(templatePath, &testString, &mockStdoutFlag)
	*findString = "" // reset to default value or this interferes with other tests

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if out != expectedString {
		t.Errorf("[FAIL] Expected test to return '%s' to standard output stream when stdout requested, however the function returned '%s'", expectedString, out)
	}
	if fileerr != nil {
		t.Errorf("[FAIL] Unexpected error raised during execution: %v", fileerr)
	}

}

func TestRenderRemoteBuiltinTemplateFileWrite(t *testing.T) {
	templatePath := "https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_1.txt.in"
	outPath := "template_1.txt"
	replaceString := "test"
	expectedString := "sha=test test=test"
	mockStdoutFlag := false
	fileerr := renderRemote(templatePath, &replaceString, &mockStdoutFlag)

	_, staterr := os.Stat(outPath)
	if !os.IsNotExist(staterr) {

		readstring, readerr := ioutil.ReadFile(outPath)
		if readerr != nil {
			t.Errorf("[FAIL] Unable to read expected text file %s. %v", outPath, readerr)
		}
		if string(readstring) != expectedString {
			t.Errorf("[FAIL] Expected to read '%s' from test file and actually read '%s'", expectedString, readstring)
		}
		os.Remove(outPath)
	} else {
		t.Errorf("[FAIL] The expected file write for the test was not found. %v", fileerr)
	}
}

func TestRenderRemoteUserTemplateFileWrite(t *testing.T) {
	templatePath := "https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_3.txt.in"
	outPath := "template_3.txt"
	replaceString := "test"
	*findString = "[[user]]"
	expectedString := "sha=test test=test"
	mockStdoutFlag := false
	fileerr := renderRemote(templatePath, &replaceString, &mockStdoutFlag)
	*findString = "" // reset to default value or this interferes with other tests

	_, staterr := os.Stat(outPath)
	if !os.IsNotExist(staterr) {

		readstring, readerr := ioutil.ReadFile(outPath)
		if readerr != nil {
			t.Errorf("[FAIL] Unable to read expected text file %s in test. %v", outPath, readerr)
		}
		if string(readstring) != expectedString {
			t.Errorf("[FAIL] Expected to read '%s' from test file and actually read '%s'", expectedString, readstring)
		}
		os.Remove(outPath)
	} else {
		t.Errorf("[FAIL] The expected file write for the test was not found. %v", fileerr)
	}
}

func TestRenderRemoteBuiltinTemplateStdout(t *testing.T) {
	templatePath := "https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_1.txt.in"
	testString := "test"
	mockStdoutFlag := true
	expectedString := "sha=test test=test"

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

	fileerr := renderRemote(templatePath, &testString, &mockStdoutFlag)

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if out != expectedString {
		t.Errorf("[FAIL] Expected test to return '%s' to standard output stream when stdout requested, however the function returned '%s'", expectedString, out)
	}
	if fileerr != nil {
		t.Errorf("[FAIL] Unexpected error raised during execution: %v", fileerr)
	}

}

func TestRenderRemoteUserTemplateStdout(t *testing.T) {
	templatePath := "https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_3.txt.in"
	testString := "test"
	*findString = "[[user]]"
	expectedString := "sha=test test=test"
	mockStdoutFlag := true

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

	fileerr := renderRemote(templatePath, &testString, &mockStdoutFlag)
	*findString = "" // reset to default value or this interferes with other tests

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	if out != expectedString {
		t.Errorf("[FAIL] Expected test to return '%s' to standard output stream when stdout requested, however the function returned '%s'", expectedString, out)
	}
	if fileerr != nil {
		t.Errorf("[FAIL] Unexpected error raised during execution: %v", fileerr)
	}

}
