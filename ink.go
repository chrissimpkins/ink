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
	"flag"
	"fmt"
	"os"

	"bytes"
	"io"

	"strings"

	"github.com/chrissimpkins/ink/inkio"
	"github.com/chrissimpkins/ink/renderers"
	"github.com/chrissimpkins/ink/validators"
)

const (
	// Version is the application version string
	Version = "0.2.0"

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

var versionShort, versionLong, helpShort, helpLong, usageLong *bool
var lintFlag, stdOutFlag, trimNLFlag *bool
var findString, replaceString *string

func init() {
	// define available command line flag arguments
	versionShort = flag.Bool("v", false, "Application version")
	versionLong = flag.Bool("version", false, "Application version")
	helpShort = flag.Bool("h", false, "Help")
	helpLong = flag.Bool("help", false, "Help")
	usageLong = flag.Bool("usage", false, "Usage")

	findString = flag.String("find", "", "Optional find string for replacement")
	replaceString = flag.String("replace", "", "Replacement string")
	lintFlag = flag.Bool("lint", false, "Lint the template file(s)")
	stdOutFlag = flag.Bool("stdout", false, "Write to standard output stream")
	trimNLFlag = flag.Bool("trimnl", false, "trim newline characters at the end of the replacement string")
}

func main() {

	// parse command line flag arguments
	flag.Parse()

	// test for at least one argument on command line
	if len(os.Args) < 2 {
		os.Stderr.WriteString("[ink] ERROR: missing arguments to the ink executable. \n")
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

	// parse all non-flag arguments on the command line
	templatePaths := flag.Args()

	/*

		COMMAND LINE VALIDATIONS

	*/
	commandlinefail := false
	for _, templatePath := range templatePaths {
		// test for existence of requested template file on user specified file path
		fileexists, fileerr := validators.FileExists(templatePath)
		if !fileexists {
			fileerrstring := fmt.Sprintf("%v", fileerr)
			os.Stderr.WriteString("[ink] ERROR: " + fileerrstring + "\n")
			commandlinefail = true
		}

		// test that template file is properly specified with `*.in` extension
		if !validators.HasCorrectExtension(templatePath) {
			os.Stderr.WriteString("[ink] ERROR: argument '" + templatePath + "' is not a properly specified template with *.in file extension.\n")
			commandlinefail = true
		}
	}

	// exit with status code 1 if any of the above command line validations failed
	if commandlinefail {
		os.Exit(1)
	}

	/*

		LINT TEMPLATES

	*/
	if *lintFlag {
		failFound := false // flag that tracks presence of linting failure(s), used for exit status code
		for _, templatePath := range templatePaths {
			// Create a new template and parse the letter into it.
			success, err := validators.LintTemplateSuccess(templatePath)
			if success {
				fmt.Println("[✓] " + templatePath + ": Valid template")
			} else {
				errstring := fmt.Sprintf("%v", err)
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

	/*

		RENDER TEMPLATES & WRITE (to file or stdout stream)

	*/
	if len(*replaceString) > 0 {
		for _, templatePath := range templatePaths {
			if *trimNLFlag {
				// trim newlines if the --trimnl flag is included in the command
				// this is used in cases where data contains undesired newline values (e.g. piped from another command)
				*replaceString = strings.TrimRight(*replaceString, "\n")
			}
			renderIt(templatePath, replaceString)
		}
	} else {
		if validators.StdinValidates(os.Stdin) {
			// there was no replace flag at the command line but there were data piped to the stdin stream
			// use standard input stream as the replacement string
			stdinReplaceBytes := new(bytes.Buffer)
			if _, err := io.Copy(stdinReplaceBytes, os.Stdin); err != nil {
				// handle failed read of stdin stream
			}

			stdinReplaceString := stdinReplaceBytes.String()
			for _, templatePath := range templatePaths {
				if *trimNLFlag {
					// trim newlines if the --trimnl flag is included in the command
					// this is used in cases where data contains undesired newline values (e.g. piped from another command)
					stdinReplaceString = strings.TrimRight(stdinReplaceString, "\n")
				}
				renderIt(templatePath, &stdinReplaceString)
			}

		} else {
			// user did not specify a replacement string with the --replace flag on the command line
			os.Stderr.WriteString("[ink] ERROR: please include a replacement string argument with the --replace flag.\n")
			os.Stderr.WriteString(Usage)
			os.Exit(1)
		}
	}
}

func renderIt(templatePath string, replaceString *string) {
	// if user specified --find flag with appropriate argument, perform user template rendering
	if len(*findString) > 0 {
		renderedStringPointer, err := renderers.RenderFromUserTemplate(templatePath, findString, replaceString)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
			os.Exit(1)
		}
		inkio.WriteString(templatePath, *stdOutFlag, renderedStringPointer)
	} else { // otherwise perform builtin template rendering
		renderedStringPointer, err := renderers.RenderFromInkTemplate(templatePath, replaceString)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("%v\n", err))
			os.Exit(1)
		}
		inkio.WriteString(templatePath, *stdOutFlag, renderedStringPointer)
	}
}
