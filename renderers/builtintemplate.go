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

// InkmarkReplaceString is a global variable that holds the replacement string data to support the use of
//  the {{ ink }} template tag as a built in value
var InkmarkReplaceString = ""

// RenderFromLocalInkTemplate is a function that renders a text template on path templatePath with a user specified
// replacement string replaceStringPointer (pointer to string) and returns pointer to rendered string and error
func RenderFromLocalInkTemplate(templatePath string, replaceStringPointer *string) (*string, error) {
	templateText, readerr := inkio.ReadFileToString(templatePath)
	emptystring := "" // returned with errors

	if readerr != nil {
		responseReadErr := fmt.Errorf("unable to read local template file '%s'. %v", templatePath, readerr)
		return &emptystring, responseReadErr
	}

	renderedStringPointer, rendererr := renderInkTemplate(&templateText, replaceStringPointer)

	if rendererr != nil {
		templateRenderErr := fmt.Errorf("unable to render local template file '%s'. %v", templatePath, rendererr)
		return &emptystring, templateRenderErr
	}

	return renderedStringPointer, rendererr
}

// RenderFromRemoteInkTemplate is a function that renders a text template at URL templateURL with a user specified
// replacement string replaceStringPointer (pointer to string) and returns pointer to rendered string and error
func RenderFromRemoteInkTemplate(templateURL string, replaceStringPointer *string) (*string, error) {
	templateText, geterr := inkio.GetRequest(templateURL)
	emptystring := "" //returned with errors

	if geterr != nil {
		responseGetErr := fmt.Errorf("unable to perform GET request for remote template file '%s'. %v", templateURL, geterr)
		return &emptystring, responseGetErr
	}

	renderedStringPointer, rendererr := renderInkTemplate(&templateText, replaceStringPointer)

	if rendererr != nil {
		templateRenderErr := fmt.Errorf("unable to render remote template pulled by GET request from '%s'. %v", templateURL, rendererr)
		return &emptystring, templateRenderErr
	}

	return renderedStringPointer, rendererr
}

// renderTemplate handles renders of the template text replacements for local and remote template files and returns
// a pointer to the rendered template string + error
func renderInkTemplate(templateText *string, replaceString *string) (*string, error) {
	// set global variable with replacement string variable (used to support `{{ ink }}` template tags via ink() function below)
	InkmarkReplaceString = *replaceString
	emptystring := ""
	funcs := template.FuncMap{"ink": ink}
	t, err := template.New("ink").Funcs(funcs).Parse(*templateText)

	if err != nil {
		return &emptystring, err
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
		return &emptystring, executeerr
	}

	renderedString := buf.String()
	return &renderedString, nil
}

// ink function supports use of the `ink` template tag as part of the builtin templating
func ink() string { return InkmarkReplaceString }
