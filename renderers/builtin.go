package renderers

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/chrissimpkins/ink/io"
)

// ReplacementStrings is a struct that maintains the strings for text replacements
type ReplacementStrings struct {
	One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten string
	Ink                                                       string
}

// RenderFromInkTemplate is a function that renders a text template on path templatePath with a user specified
// replacement string replaceString (pointer to string) and returns pointer to rendered string and error
func RenderFromInkTemplate(templatePath string, replaceString *string) (*string, error) {
	templateText, readerr := io.ReadFileToString(templatePath)
	emptystring := "" // returned with errors

	if readerr != nil {
		responseReadErr := fmt.Errorf("[ink] ERROR: unable to read template file '%s'. %v\n", templatePath, readerr)
		return &emptystring, responseReadErr
	}
	//funcs := template.FuncMap{"ink": ink}
	//t, err := template.New("ink").Funcs(funcs).Parse(templateText)
	t, err := template.New("ink").Parse(templateText)
	if err != nil {
		templateParseErr := fmt.Errorf("[ink] ERROR: ink template '%s' could not be parsed. %v\n", templatePath, err)
		return &emptystring, templateParseErr
	}

	r := ReplacementStrings{
		*replaceString,
		"{{.One}}",
		"{{.Two}}",
		"{{.Three}}",
		"{{.Four}}",
		"{{.Five}}",
		"{{.Six}}",
		"{{.Seven}}",
		"{{.Eight}}",
		"{{.Nine}}",
		*replaceString,
	}
	buf := new(bytes.Buffer)
	executeerr := t.Execute(buf, r)
	if executeerr != nil {
		templateRenderErr := fmt.Errorf("[ink] ERROR: while executing template '%s' encountered error: %v\n", templatePath, executeerr)
		return &emptystring, templateRenderErr
	}

	renderedString := buf.String()
	return &renderedString, nil
}
