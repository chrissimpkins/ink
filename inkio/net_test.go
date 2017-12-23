package inkio

import "testing"

func TestGetRequestValidURL(t *testing.T) {
	responsePointer, err := GetRequest("https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/simpletext.txt")
	if responsePointer != "This is simple text" {
		t.Errorf("[FAIL] GetRequest function should return 'This is simple text' and instead returned '%s'", responsePointer)
	}
	if err != nil {
		t.Errorf("[FAIL] GetRequest function should not have returned an error and returned: %v", err)
	}
}

func TestGetRequestInvalidURL(t *testing.T) {
	responsePointer, err := GetRequest("https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/completelybogus.txt")
	if responsePointer != "" {
		t.Errorf("[FAIL] GetRequest function should have returned empty string for invalid URL, instead returned %s", responsePointer)
	}
	if err == nil {
		t.Errorf("[FAIL] GetRequest function should have returned an error on an invalid URL, instead returned nil")
	}
}
