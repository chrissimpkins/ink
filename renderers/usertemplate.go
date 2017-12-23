// usertemplate holds the rendering implementation for user specified text file template syntax rendering
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
	"fmt"
	"strings"

	"github.com/chrissimpkins/ink/inkio"
)

// RenderFromLocalUserTemplate is a function that renders a text template file on the path templatePath to
// a rendered string using the findString string pointer replacement target substring with the
// replaceString string pointer replacement substring
func RenderFromLocalUserTemplate(templatePath string, findString *string, replaceString *string) (*string, error) {
	templateText, readerr := inkio.ReadFileToString(templatePath)
	emptystring := "" // returned with errors

	if readerr != nil {
		responseReadErr := fmt.Errorf("unable to read local template file '%s'. %v", templatePath, readerr)
		return &emptystring, responseReadErr
	}

	renderedStringPointer, rendererr := renderUserTemplate(&templateText, findString, replaceString)

	if rendererr != nil {
		renderErr := fmt.Errorf("unable to render local template file '%s'. %v", templatePath, rendererr)
		return &emptystring, renderErr
	}
	return renderedStringPointer, rendererr

}

// RenderFromRemoteUserTemplate is a function that renders a text template at the URL templateURL to
// a rendered string using the findString string pointer replacement target substring with the
// replaceString string pointer replacement substring
func RenderFromRemoteUserTemplate(templateURL string, findString *string, replaceString *string) (*string, error) {
	templateText, geterr := inkio.GetRequest(templateURL)
	emptystring := "" // returned with errors

	if geterr != nil {
		responseReadErr := fmt.Errorf("unable to perform GET request for remote template file '%s'. %v", templateURL, geterr)
		return &emptystring, responseReadErr
	}

	renderedStringPointer, rendererr := renderUserTemplate(&templateText, findString, replaceString)

	if rendererr != nil {
		renderErr := fmt.Errorf("unable to render remote template file '%s'. %v", templateURL, rendererr)
		return &emptystring, renderErr
	}
	return renderedStringPointer, rendererr
}

// renderUserTemplate is a function that performs the text string replacements in user templates across both
// local and remote template files
func renderUserTemplate(templateText *string, findString *string, replaceString *string) (*string, error) {
	// replace all instances of user specified template findString with user specified replaceString
	renderedString := strings.Replace(*templateText, *findString, *replaceString, -1)

	return &renderedString, nil
}
