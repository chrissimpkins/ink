package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)


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
		os.Stdout.WriteString("uni v" + Version + "\n")
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

	templatePath := os.Args[len(os.Args)-1]

	// LINT: lint the template files
	if *lintFlag {
		// Create a new template and parse the letter into it.
		success := lintTemplateSuccess(templatePath)
		if success {
			fmt.Println("Linting was successful")
			os.Exit(0)
		} else {
			os.Stderr.WriteString("Linting failed")
			os.Exit(1)
		}

	}

	// Handle single pass replacement from command line `--replace=[string]` flag
	if len(*replaceString) > 0 {
		templateText := readFileToString(templatePath)
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
			defer f.Close()
			f.WriteString(buf.String())
			f.Sync()
		}

	}

}

func lintTemplateSuccess(filePath string) bool {
	templateText := readFileToString(filePath)
	_, err := template.New("ink").Parse(templateText)
	if err != nil {
		return false
	}

	return true
}

func readFileToString(filepath string) string {

	byteString, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return string(byteString)
}
