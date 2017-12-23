package renderers

import (
	"path/filepath"
	"testing"
)

func TestRenderBuiltinLowercaseInkTag(t *testing.T) {
	tests := []struct {
		templatepath string
		replacement  string
		expected     string
	}{
		{filepath.Join("..", "testfiles", "template_1.txt.in"), "abcd123", "sha=abcd123 test=abcd123"},
		{filepath.Join("..", "testfiles", "template_1.txt.in"), "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫"},
		{filepath.Join("..", "testfiles", "template_1.txt.in"), "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂"},
	}

	for _, testcase := range tests {
		haystack, err := RenderFromLocalInkTemplate(testcase.templatepath, &testcase.replacement)
		if err != nil {
			t.Errorf("[FAIL] RenderFromInkTemplate execution returned error value: %v", err)
		}
		if *haystack != testcase.expected {
			t.Errorf("[FAIL] Expected rendered template value = '%s' and received rendered template value '%s'", testcase.expected, *haystack)
		}
	}
}

func TestRenderBuiltinUppercaseInkTag(t *testing.T) {
	tests := []struct {
		templatepath string
		replacement  string
		expected     string
	}{
		{filepath.Join("..", "testfiles", "template_2.txt.in"), "abcd123", "sha=abcd123 test=abcd123"},
		{filepath.Join("..", "testfiles", "template_2.txt.in"), "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫"},
		{filepath.Join("..", "testfiles", "template_2.txt.in"), "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂"},
	}

	for _, testcase := range tests {
		haystack, err := RenderFromLocalInkTemplate(testcase.templatepath, &testcase.replacement)
		if err != nil {
			t.Errorf("[FAIL] RenderFromInkTemplate execution returned error value: %v", err)
		}
		if *haystack != testcase.expected {
			t.Errorf("[FAIL] Expected rendered template value = '%s' and received rendered template value '%s'", testcase.expected, *haystack)
		}
	}
}

func TestRenderBuiltinBadFilePathRaisesError(t *testing.T) {
	replacestring := "testing"
	_, err := RenderFromLocalInkTemplate("completelybogus.txt.in", &replacestring)
	if err == nil {
		t.Errorf("[FAIL] Expected error to be raised for invalid file path and the error value was 'nil'")
	}
}
