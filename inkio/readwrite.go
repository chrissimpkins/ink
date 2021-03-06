// readwrite holds the I/O functions for the ink application
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

package inkio

import (
	"io/ioutil"
	"os"
)

// ReadFileToString reads text files from disk and returns (string, error)
func ReadFileToString(filepath string) (string, error) {
	byteString, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(byteString), nil
}

// WriteString writes a rendered string renderedStringPointer to file or to the standard output stream as determined
// by the stdOutFlag boolean parameter value.  File writes occur on a path that is created from templatePath with the
// `.in` file extension suffix removed from the file path
func WriteString(templatePath string, stdOutFlag bool, renderedStringPointer *string) error {
	if stdOutFlag {
		os.Stdout.WriteString(*renderedStringPointer)
	} else {
		outPath := templatePath[0 : len(templatePath)-3]
		f, err1 := os.Create(outPath)
		defer f.Close()
		if err1 != nil {
			return err1
		}
		_, err2 := f.WriteString(*renderedStringPointer)
		if err2 != nil {
			return err2
		}
		f.Sync()
	}
	return nil
}
