// builtintemplate holds the rendering implementation for builtin (Go style) text file template syntax rendering
/*
MIT License

Copyright (c) 2017 Chris Simpkins

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package renderers

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/chrissimpkins/ink/inkio"
)

// ReplacementStrings is a struct that maintains the strings for text replacements
type ReplacementStrings struct {
	One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten string
	Ink                                                       string
}

// RenderFromInkTemplate is a function that renders a text template on path templatePath with a user specified
// replacement string replaceString (pointer to string) and returns pointer to rendered string and error
func RenderFromInkTemplate(templatePath string, replaceString *string) (*string, error) {
	templateText, readerr := inkio.ReadFileToString(templatePath)
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
