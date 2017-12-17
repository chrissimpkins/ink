package validators

import (
	"path/filepath"
	"testing"
)

func TestLintTemplateSuccessValidTemplate(t *testing.T) {
	result, _ := LintTemplateSuccess(filepath.Join("..", "testfiles", "template_1.txt.in"))
	if result == false {
		t.Errorf("[FAIL] LintTemplateSuccess returned false for a valid template, expected true.")
	}
}

func TestLintTemplateSuccessInvalidTemplate(t *testing.T) {
	result, _ := LintTemplateSuccess(filepath.Join("..", "testfiles", "template_2.txt.in"))
	if result == true {
		t.Errorf("[FAIL] LintTemplateSuccess returned true for an invalid template, expected false.")
	}
}

func TestLintTemplateSuccessBadFilePath(t *testing.T) {
	result, err := LintTemplateSuccess(filepath.Join("..", "testfiles", "totallybogus.txt"))
	if result == true {
		t.Errorf("[FAIL] LintTemplateSuccess returned true for a file path to a missing file, expected false.")
	}
	if err == nil {
		t.Errorf("[FAIL] LintTemplateSuccess returned nil for error when a file path to a missing file was tested, expected error message.")
	}
}
