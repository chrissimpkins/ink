package validators

import (
	"text/template"

	"github.com/chrissimpkins/ink/io"
)

func LintTemplateSuccess(filePath string) (bool, error) {
	templateText, readerr := io.ReadFileToString(filePath)
	if readerr != nil {
		return false, readerr
	}
	_, templateerr := template.New("ink").Parse(templateText)
	if templateerr != nil {
		return false, templateerr
	}

	return true, templateerr
}
