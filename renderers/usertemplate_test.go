package renderers

import (
	"path/filepath"
	"testing"
)

/*
Local user template tests
*/

func TestRenderUserBracketsLocal(t *testing.T) {
	tests := []struct {
		templatepath string
		replacement  string
		expected     string
		find         string
	}{
		{filepath.Join("..", "testfiles", "template_3.txt.in"), "abcd123", "sha=abcd123 test=abcd123", "[[user]]"},
		{filepath.Join("..", "testfiles", "template_3.txt.in"), "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫", "[[user]]"},
		{filepath.Join("..", "testfiles", "template_3.txt.in"), "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂", "[[user]]"},
	}

	for _, testcase := range tests {
		haystack, err := RenderFromLocalUserTemplate(testcase.templatepath, &testcase.find, &testcase.replacement)
		if err != nil {
			t.Errorf("[FAIL] RenderFromInkTemplate execution returned error value: %v", err)
		}
		if *haystack != testcase.expected {
			t.Errorf("[FAIL] Expected rendered template value = '%s' and received rendered template value '%s'", testcase.expected, *haystack)
		}
	}
}

func TestRenderUserBadFilePathRaisesError(t *testing.T) {
	replacestring := "testing"
	findstring := "[[user]]"
	_, err := RenderFromLocalUserTemplate("completelybogus.txt.in", &findstring, &replacestring)
	if err == nil {
		t.Errorf("[FAIL] Expected error to be raised for invalid file path and the error value was 'nil'")
	}
}

/*
Remote user template tests
*/

func TestRenderUserBracketsRemote(t *testing.T) {
	tests := []struct {
		templateurl string
		replacement string
		expected    string
		find        string
	}{
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_3.txt.in", "abcd123", "sha=abcd123 test=abcd123", "[[user]]"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_3.txt.in", "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫", "[[user]]"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_3.txt.in", "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂", "[[user]]"},
	}

	for _, testcase := range tests {
		haystack, err := RenderFromRemoteUserTemplate(testcase.templateurl, &testcase.find, &testcase.replacement)
		if err != nil {
			t.Errorf("[FAIL] RenderFromInkTemplate execution returned error value: %v", err)
		}
		if *haystack != testcase.expected {
			t.Errorf("[FAIL] Expected rendered template value = '%s' and received rendered template value '%s'", testcase.expected, *haystack)
		}
	}
}

func TestRenderUserBadURLRaisesError(t *testing.T) {
	replacestring := "testing"
	findstring := "[[user]]"
	_, err := RenderFromRemoteUserTemplate("https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/completelybogus.txt.in", &findstring, &replacestring)
	if err == nil {
		t.Errorf("[FAIL] Expected error to be raised for invalid URL and the error value was 'nil'")
	}
}
