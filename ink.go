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

const (
	// Application version string
	Version = "0.1.0"

	// Application usage string
	Usage = "Usage: ink [args]"

	// Application help string
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

func main() {
	// test for at least one argument on command line
	if len(os.Args) < 2 {
		os.Stderr.WriteString("[Error] Please include at least one argument for your Unicode code point search\n")
		os.Stderr.WriteString(Usage)
		os.Exit(1)
	}

	// define available command line flags
	var versionShort = flag.Bool("v", false, "Application version")
	var versionLong = flag.Bool("version", false, "Application version")
	var helpShort = flag.Bool("h", false, "Help")
	var helpLong = flag.Bool("help", false, "Help")
	var usageLong = flag.Bool("usage", false, "Usage")

	var replaceString = flag.String("replace", "", "Optional string replacement")
	var lintFlag = flag.Bool("lint", false, "Lint the template file(s)")
	var stdOutFlag = flag.Bool("stdout", false, "Write to standard output stream")

	flag.Parse()

	// parse help, version, usage command line flags and handle them
	switch {
	case *versionShort:
		os.Stdout.WriteString("uni v" + Version + "\n")
		os.Exit(0)
	case *versionLong:
		os.Stdout.WriteString("uni v" + Version + "\n")
		os.Exit(0)
	case *helpShort:
		os.Stdout.WriteString(Help)
		os.Exit(0)
	case *helpLong:
		os.Stdout.WriteString(Help)
		os.Exit(0)
	case *usageLong:
		os.Stdout.WriteString(Usage)
		os.Exit(0)
	}

	// # TODO: add template file path check for `.in` extension

	type ReplacementString struct {
		One, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten string
		Ink                                                       string
	}

	templatePath := os.Args[len(os.Args)-1]

	// LINT: lint the template files
	if *lintFlag {
		// Create a new template and parse the letter into it.
		success, response := lintTemplateSuccess(templatePath)
		if success {
			fmt.Printf("%s\n", response)
			os.Exit(0)
		} else {
			os.Stderr.WriteString(response)
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
		r := ReplacementString{
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

func lintTemplateSuccess(filePath string) (bool, string) {
	templateText := readFileToString(filePath)
	_, err := template.New("ink").Parse(templateText)
	if err != nil {
		return false, fmt.Sprintf("[ink] LINT ERROR \nFile: '%s' \nError: %v\n", filePath, err)
	}

	return true, fmt.Sprintf("[ink] LINT \nFile: '%s': no template errors\n", filePath)
}

func readFileToString(filepath string) string {

	byteString, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return string(byteString)
}
