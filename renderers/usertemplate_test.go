package renderers

import (
	"path/filepath"
	"testing"
)

func TestRenderUserBrackets(t *testing.T) {
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
		haystack, err := RenderFromUserTemplate(testcase.templatepath, &testcase.find, &testcase.replacement)
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
	_, err := RenderFromUserTemplate("completelybogus.txt.in", &findstring, &replacestring)
	if err == nil {
		t.Errorf("[FAIL] Expected error to be raised for invalid file path and the error value was 'nil'")
	}
}
