## Changelog

### v0.4.0

- refactored entire approach to parallel template renders/file writes from stdin and command line argument defined template renders
    - eliminated duplicated source code
    - added detection of template rendering errors across all go routines
    - added appropriate exit status code reporting for any rendering errors during parallel renders

### v0.3.1

- [main] added standard input stream error handling

### v0.3.0

- added support for parallel template rendering
- added support for concurrent template writes

### v0.2.1

- [main] added `--help` documentation with all available options

### v0.2.0

- [main] added support for line filter functionality (i.e. pipe replacement string through std input stream from another command)
- [main] added new `--trimnl` option to remove newline value from replacement string
- [main] refactored ink.go source file template rendering handling to a new function
- [validators] added new standard input stream validator function


### v0.1.0

- initial release