package main

import (
	"flag"
	"os"
)

const (
	Version = "0.1.0"
	Usage   = "Usage: ink [args]"
	Help    = "=================================================\n" +
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
	flag.Parse()

	// parse command line flags and handle them
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
}
