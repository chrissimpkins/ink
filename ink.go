// ink holds the main function for the ink templating application executable
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

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/chrissimpkins/ink/io"
	"github.com/chrissimpkins/ink/validators"
)

const (
	// Version is the application version string
	Version = "0.1.0"

	// Usage is the application usage string
	Usage = "Usage: ink [options] [template path 1]...[template path n]\n"

	// Help is the application help string
	Help = "=================================================\n" +
		" ink v" + Version + "\n" +
		" Copyright 2017 Christopher Simpkins\n" +
		" MIT License\n\n" +
		" Source: https://github.com/chrissimpkins/ink\n" +
		"=================================================\n\n" +
		" Usage:\n" +
		"  $ ink [args]\n\n" +
		" Options:\n" +
		" -h, --help           Application help\n" +
		"     --usage          Application usage\n" +
		" -v, --version        Application version\n\n"
)

// ReplacementStrings is a struct that maintains the strings for text replacements
type ReplacementStrings struct {
	One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten string
	Ink                                                       string
}

var versionShort, versionLong, helpShort, helpLong, usageLong *bool
var lintFlag, stdOutFlag *bool
var replaceString *string

func init() {
	// define available command line flag arguments
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")

	replaceString = flag.String("replace", "", "Optional string replacement")
	lintFlag = flag.Bool("lint", false, "Lint the template file(s)")
	stdOutFlag = flag.Bool("stdout", false, "Write to standard output stream")
}

func main() {

	// parse command line flag arguments
	flag.Parse()

	// test for at least one argument on command line
	if len(os.Args) < 2 {
		os.Stderr.WriteString("[Error] Please include at least one argument for your Unicode code point search\n")
		os.Stderr.WriteString(Usage)
		os.Exit(1)
	}

	// parse help, version, usage command line flags and handle them
	switch {
	case *versionShort, *versionLong:
		os.Stdout.WriteString("ink v" + Version + "\n")
		os.Exit(0)
	case *helpShort, *helpLong:
		os.Stdout.WriteString(Help)
		os.Exit(0)
	case *usageLong:
		os.Stdout.WriteString(Usage)
		os.Exit(0)
	}

	// TODO: add file path checks
	// TODO: add template file path check for `.in` extension
	// TODO: improve user messages for fails, successes

	//templatePath := os.Args[len(os.Args)-1]
	templatePaths := flag.Args()

	// LINT: lint the template files
	if *lintFlag {
		failFound := false // flag that tracks presence of linting failure(s), used for exit status code
		for _, templatePath := range templatePaths {
			// Create a new template and parse the letter into it.
			success, err := validators.LintTemplateSuccess(templatePath)
			if success {
				fmt.Println("[✓] " + templatePath + ": Valid template")
			} else {
				errstring := fmt.Sprintf("%s", err)
				os.Stderr.WriteString("[X] " + templatePath + ": FAIL --- " + errstring + "\n")
				failFound = true
			}
		}
		// if found a linting failure for any requested template file, exit with status code 1
		if failFound {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}

	// Handle single pass replacement from command line `--replace=[string]` flag
	if len(*replaceString) > 0 {
		for _, templatePath := range templatePaths {
			templateText, readerr := io.ReadFileToString(templatePath)
			if readerr != nil {
				errstring := fmt.Sprintf("%v", readerr)
				os.Stderr.WriteString("[ink] ERROR: unable to read template file '" + templatePath + "'. " + errstring + "\n")
			}
			//funcs := template.FuncMap{"ink": ink}
			//t, err := template.New("ink").Funcs(funcs).Parse(templateText)
			t, err := template.New("ink").Parse(templateText)
			if err != nil {
				os.Stderr.WriteString("[ink] ERROR: ink template could not be parsed. " + fmt.Sprintf("%v", err))
				os.Exit(1)
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
			outPath := templatePath[0 : len(templatePath)-3]

			buf := new(bytes.Buffer)
			executeerr := t.Execute(buf, r)
			if executeerr != nil {
				errstring := fmt.Sprintf("[ink] ERROR: while executing template '%s' encountered error: %v\n", templatePath, executeerr)
				os.Stderr.WriteString(errstring)
				os.Exit(1)
			}

			if *stdOutFlag {
				os.Stdout.WriteString(buf.String())
			} else {
				f, err := os.Create(outPath)
				if err != nil {
					os.Stderr.WriteString("**** TODO ****")
					os.Exit(1)
				}
				f.WriteString(buf.String())
				f.Sync()
				f.Close()
			}
		}
	}
}
