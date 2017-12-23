package utilities

import "testing"

func TestGetURLFilePathInkTemplateDepthOne(t *testing.T) {
	response, err := GetURLFilePath("http://test.com/inktemplate.txt.in")
	if err != nil {
		t.Errorf("[FAIL] Did not expect error returned from GetURLFilePath for valid file, received: %v", err)
	}
	if response != "inktemplate.txt.in" {
		t.Errorf("[FAIL] Expected return of path 'inktemplate.txt.in', received: %s", response)
	}
}

func TestGetURLFilePathInkTemplateDepthTwo(t *testing.T) {
	response, err := GetURLFilePath("http://test.com/testing/inktemplate.txt.in")
	if err != nil {
		t.Errorf("[FAIL] Did not expect error returned from GetURLFilePath for valid file, received: %v", err)
	}
	if response != "inktemplate.txt.in" {
		t.Errorf("[FAIL] Expected return of path 'inktemplate.txt.in', received: %s", response)
	}
}

func TestGetURLFilePathInkTemplateDepthThree(t *testing.T) {
	response, err := GetURLFilePath("http://test.com/testing/more/inktemplate.txt.in")
	if err != nil {
		t.Errorf("[FAIL] Did not expect error returned from GetURLFilePath for valid file, received: %v", err)
	}
	if response != "inktemplate.txt.in" {
		t.Errorf("[FAIL] Expected return of path 'inktemplate.txt.in', received: %s", response)
	}
}

func TestGetURLFilePathInkTemplateWithQuery(t *testing.T) {
	response, err := GetURLFilePath("http://test.com/inktemplate.txt.in?q=test")
	if err != nil {
		t.Errorf("[FAIL] Did not expect error returned from GetURLFilePath for valid file, received: %v", err)
	}
	if response != "inktemplate.txt.in" {
		t.Errorf("[FAIL] Expected return of path 'inktemplate.txt.in', received: %s", response)
	}
}

func TestGetURLFilePathInkTemplateWithMultiQuery(t *testing.T) {
	response, err := GetURLFilePath("http://test.com/inktemplate.txt.in?q=test&t=more")
	if err != nil {
		t.Errorf("[FAIL] Did not expect error returned from GetURLFilePath for valid file, received: %v", err)
	}
	if response != "inktemplate.txt.in" {
		t.Errorf("[FAIL] Expected return of path 'inktemplate.txt.in', received: %s", response)
	}
}

// when protocol is not included, still returns correct file path
func TestGetURLFilePathInkTemplateNoProtocol(t *testing.T) {
	response, err := GetURLFilePath("test.com/inktemplate.txt.in")
	if err != nil {
		t.Errorf("[FAIL] Did not expect error returned from GetURLFilePath for valid file, received: %v", err)
	}
	if response != "inktemplate.txt.in" {
		t.Errorf("[FAIL] Expected return of path 'inktemplate.txt.in', received: %s", response)
	}
}

// when protocol and path structure is not included, still returns correct file path
func TestGetURLFilePathInkTemplateNoProtocolOrDirs(t *testing.T) {
	response, err := GetURLFilePath("inktemplate.txt.in")
	if err != nil {
		t.Errorf("[FAIL] Did not expect error returned from GetURLFilePath for valid file, received: %v", err)
	}
	if response != "inktemplate.txt.in" {
		t.Errorf("[FAIL] Expected return of path 'inktemplate.txt.in', received: %s", response)
	}
}
