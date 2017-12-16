package validators

import "testing"

// FileExists function tests

func TestFileExistsFilePresent(t *testing.T) {
	result, err := FileExists("../testfiles/template_1.txt.in")
	if result == false {
		t.Errorf("[FAIL] FileExists function test for existing file returned a value of 'false'")
	}
	if err != nil {
		t.Errorf("[FAIL] FileExists function test for existing file returned an error: %v", err)
	}
}

func TestFileExistsFileMissing(t *testing.T) {
	result, err := FileExists("../testfiles/totallybogus.txt")
	if result == true {
		t.Errorf("[FAIL] FileExists function test for missing file returned a value of 'true'")
	}
	if err == nil {
		t.Errorf("[FAIL] FileExists function test for missing file did not return an error")
	}
}

// HasCorrectExtension function tests

func TestHasCorrectExtensionValidFile(t *testing.T) {
	result := HasCorrectExtension("../testfiles/template_1.txt.in")
	if result == false {
		t.Errorf("[FAIL] HasCorrectExtension returned false when the filepath string has correct extension")
	}
}

func TestHasCorrectExtensionInValidFile(t *testing.T) {
	result := HasCorrectExtension("boguspath/template_1.txt")
	if result == true {
		t.Errorf("[FAIL] HasCorrectExtension returned true when the filepath string does not have correct extension")
	}
}
