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
	"io"
	"os"
	"strings"
	"sync"

	"github.com/chrissimpkins/ink/inkio"
	"github.com/chrissimpkins/ink/renderers"
	"github.com/chrissimpkins/ink/validators"
)

const (
	// Version is the application version string
	Version = "0.6.0dev1"

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
		"  $ ink [options] [template path 1]...[template path n]\n\n" +
		" Options:\n" +
		"     --find=          Find string value for template render\n" +
		" -h, --help           Application help\n" +
		"     --lint           Lint ink template file\n" +
		"     --replace=       Replacement string value for template render\n" +
		"     --stdout         Write render to standard output stream\n" +
		"     --trimnl         Trim newline value from replacement string\n" +
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
		os.Stderr.WriteString("[ink] ERROR: Missing arguments to the ink executable. \n")
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
			os.Stderr.WriteString("[ink] ERROR: Argument '" + templatePath + "' is not a properly specified template with *.in file extension.\n")
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

	   PREPARE THE REPLACEMENT STRING FOR WRITE

	*/

	if len(*replaceString) > 0 {
		// do nothing, gtg if defined
	} else if validators.StdinValidates(os.Stdin) {
		// there was no replace flag at the command line but there were data piped to the stdin stream
		// use standard input stream as the replacement string
		stdinReplaceBytes := new(bytes.Buffer)
		if _, err := io.Copy(stdinReplaceBytes, os.Stdin); err != nil {
			os.Stderr.WriteString("[ink] ERROR: Unable to read standard input stream. " + fmt.Sprintf("%v\n", err))
			os.Exit(1)
		}

		*replaceString = stdinReplaceBytes.String()

	} else {
		// user did not specify a replacement string with the --replace flag on the command line
		// or pipe replacement string to stdin stream
		os.Stderr.WriteString("[ink] ERROR: Missing replacement string for template render.\n")
		os.Stderr.WriteString(Usage)
		os.Exit(1)
	}

	// Trim newlines if requested on commandline with --trimnl flag
	if *trimNLFlag {
		*replaceString = strings.TrimRight(*replaceString, "\n")
	}

	/*

		RENDER TEMPLATES & WRITE (to file or stdout stream)
		- renders multi-template requests in parallel

	*/

	var wg sync.WaitGroup

	errorc := make(chan bool) // channel used to communicate render/write failures from go routines that are executing them
	for _, templatePath := range templatePaths {
		wg.Add(1)
		go func(templatePath string, replaceString *string) {
			defer wg.Done()
			err := renderIt(templatePath, replaceString)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("[ink] ERROR: Failed to render template %s to file. %v\n", templatePath, err))
				errorc <- true // true = error occurred
			}
			errorc <- false                 // false = error did not occur
			if !*stdOutFlag && err == nil { // print confirmation only if the user did not print to render to stdout stream
				fmt.Printf("[ink] Template %s rendered successfully.\n", templatePath)
			}
		}(templatePath, replaceString)
	}

	// must make the wait and close concurrent with the executing worker go routines
	go func() {
		wg.Wait()
		close(errorc)
	}()

	exitFail := false                   // flag to indicate that a failure occurred for appropriate exit status code on application exit
	for errorOccurred := range errorc { // iterate through the booleans to determine if error occurred during execution
		if errorOccurred == true {
			exitFail = true
		}
	}

	if exitFail {
		os.Exit(1) // fail with exit status code 1 if error occurred during execution of any template renders
	}

	// reachable only if error did not occur
	// indicate render completed successfully if not printing to stdout stream
	// this is intended for user notification in the setting of "long" running multi-template renders
	if !*stdOutFlag { // confirm that the writes are all complete if user did not render to stdout stream
		os.Stdout.WriteString("[ink] Render complete.\n")
	}
}

func renderIt(templatePath string, replaceString *string) error {
	// if user specified --find flag with appropriate argument, perform user template rendering
	if len(*findString) > 0 {
		renderedStringPointer, rendererr := renderers.RenderFromLocalUserTemplate(templatePath, findString, replaceString)
		if rendererr != nil {
			return rendererr
		}
		writeerr := inkio.WriteString(templatePath, *stdOutFlag, renderedStringPointer)
		if writeerr != nil {
			return writeerr
		}
		return nil
	}
	// otherwise perform builtin template rendering
	renderedStringPointer, rendererr := renderers.RenderFromLocalInkTemplate(templatePath, replaceString)
	if rendererr != nil {
		return rendererr
	}
	writeerr := inkio.WriteString(templatePath, *stdOutFlag, renderedStringPointer)
	if writeerr != nil {
		return writeerr
	}
	return nil
}
