package renderers

import (
	"path/filepath"
	"testing"
)

/*
 LOCAL TEMPLATE TESTS
*/

func TestRenderBuiltinLowercaseInkTagLocal(t *testing.T) {
	tests := []struct {
		templatepath string
		replacement  string
		expected     string
	}{
		{filepath.Join("..", "testfiles", "template_1.txt.in"), "abcd123", "sha=abcd123 test=abcd123"},
		{filepath.Join("..", "testfiles", "template_1.txt.in"), "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫"},
		{filepath.Join("..", "testfiles", "template_1.txt.in"), "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂"},
		{filepath.Join("..", "testfiles", "template_4.txt.in"), "abcd123", "饂饂饂=abcd123 åß∂=abcd123"},
		{filepath.Join("..", "testfiles", "template_4.txt.in"), "åß∂ƒç√∫", "饂饂饂=åß∂ƒç√∫ åß∂=åß∂ƒç√∫"},
		{filepath.Join("..", "testfiles", "template_4.txt.in"), "饂饂饂饂", "饂饂饂=饂饂饂饂 åß∂=饂饂饂饂"},
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

func TestRenderBuiltinUppercaseInkTagLocal(t *testing.T) {
	tests := []struct {
		templatepath string
		replacement  string
		expected     string
	}{
		{filepath.Join("..", "testfiles", "template_2.txt.in"), "abcd123", "sha=abcd123 test=abcd123"},
		{filepath.Join("..", "testfiles", "template_2.txt.in"), "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫"},
		{filepath.Join("..", "testfiles", "template_2.txt.in"), "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂"},
		{filepath.Join("..", "testfiles", "template_5.txt.in"), "abcd123", "饂饂饂=abcd123 åß∂=abcd123"},
		{filepath.Join("..", "testfiles", "template_5.txt.in"), "åß∂ƒç√∫", "饂饂饂=åß∂ƒç√∫ åß∂=åß∂ƒç√∫"},
		{filepath.Join("..", "testfiles", "template_5.txt.in"), "饂饂饂饂", "饂饂饂=饂饂饂饂 åß∂=饂饂饂饂"},
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

func TestRenderBuiltinBadLocalFilePathRaisesError(t *testing.T) {
	replacestring := "testing"
	_, err := RenderFromLocalInkTemplate("completelybogus.txt.in", &replacestring)
	if err == nil {
		t.Errorf("[FAIL] Expected error to be raised for invalid file path and the error value was 'nil'")
	}
}

/*
REMOTE TEMPLATE TESTS
*/

func TestRenderBuiltinLowercaseInkTagRemote(t *testing.T) {
	tests := []struct {
		templateurl string
		replacement string
		expected    string
	}{
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_1.txt.in", "abcd123", "sha=abcd123 test=abcd123"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_1.txt.in", "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_1.txt.in", "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_4.txt.in", "abcd123", "饂饂饂=abcd123 åß∂=abcd123"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_4.txt.in", "åß∂ƒç√∫", "饂饂饂=åß∂ƒç√∫ åß∂=åß∂ƒç√∫"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_4.txt.in", "饂饂饂饂", "饂饂饂=饂饂饂饂 åß∂=饂饂饂饂"},
	}

	for _, testcase := range tests {
		haystack, err := RenderFromRemoteInkTemplate(testcase.templateurl, &testcase.replacement)
		if err != nil {
			t.Errorf("[FAIL] RenderFromInkTemplate execution returned error value: %v", err)
		}
		if *haystack != testcase.expected {
			t.Errorf("[FAIL] Expected rendered template value = '%s' and received rendered template value '%s'", testcase.expected, *haystack)
		}
	}
}

func TestRenderBuiltinUppercaseInkTagRemote(t *testing.T) {
	tests := []struct {
		templateurl string
		replacement string
		expected    string
	}{
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_2.txt.in", "abcd123", "sha=abcd123 test=abcd123"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_2.txt.in", "åß∂ƒç√∫", "sha=åß∂ƒç√∫ test=åß∂ƒç√∫"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_2.txt.in", "饂饂饂饂", "sha=饂饂饂饂 test=饂饂饂饂"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_5.txt.in", "abcd123", "饂饂饂=abcd123 åß∂=abcd123"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_5.txt.in", "åß∂ƒç√∫", "饂饂饂=åß∂ƒç√∫ åß∂=åß∂ƒç√∫"},
		{"https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/template_5.txt.in", "饂饂饂饂", "饂饂饂=饂饂饂饂 åß∂=饂饂饂饂"},
	}

	for _, testcase := range tests {
		haystack, err := RenderFromRemoteInkTemplate(testcase.templateurl, &testcase.replacement)
		if err != nil {
			t.Errorf("[FAIL] RenderFromInkTemplate execution returned error value: %v", err)
		}
		if *haystack != testcase.expected {
			t.Errorf("[FAIL] Expected rendered template value = '%s' and received rendered template value '%s'", testcase.expected, *haystack)
		}
	}
}

func TestRenderBuiltinBadURLRaisesError(t *testing.T) {
	replacestring := "testing"
	_, err := RenderFromRemoteInkTemplate("https://raw.githubusercontent.com/chrissimpkins/ink/master/testfiles/completelybogus.txt.in", &replacestring)
	if err == nil {
		t.Errorf("[FAIL] Expected error to be raised for invalid URL and the error value was 'nil'")
	}
}
