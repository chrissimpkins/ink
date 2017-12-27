## Changelog

### v0.7.1

- minor refactor of function parameter in `GetURLFilePath` of utilities/url.go source file
- minor refactor of if block in `renderUserTemplate` function of renderers/usertemplates.go source file
- minor in-application help text revisions

### v0.7.0

- added regular expression substitution support to the `--find` option for local and remote source text files
- added new command line regular expression syntax definition `--find="{{regex pattern}}"`
- modified error messages in ink.go source file - removed render to file wording as not all renders are to files and this did not differentiate, was unnecessary language
- updated in-application help string

### v0.6.3

- added ink template tests for expanded Unicode characters
- added user-defined template tests for expanded Unicode characters
- updated documentation for usage, available options, template specifications

### v0.6.2

- added appveyor.yml settings file
- added Appveyor Windows platform testing x i386 + x86_64 architectures

### v0.6.1

- modified .goreleaser.yml script: eliminated macOS builds x i386 architecture

### v0.6.0

- added support for parallel rendering of remote templates that are accessible via GET requests (includes builtin and user defined formats)
- added utilities/url.go utilities package module with URL file path parsing function
- refactored template file extension validation to only run when `--stdout` flag is not specified.  Assume user is going to manage outfile path (or does not need it) when this flag is used.
- updated usage string
- updated help string
- added new `build.sh` cross-platform compiler shell script

### v0.5.0

- added builtin template support for `{{ ink }}` template tags
- added template linter support for `{{ ink }}` template tag validation

### v0.4.0

- refactored entire approach to parallel template renders/file writes from stdin and command line argument defined template renders
    - eliminated duplicated source code
    - added detection of template rendering errors across all go routines
    - added appropriate exit status code reporting for any rendering errors during parallel renders

### v0.3.1

- added standard input stream error handling

### v0.3.0

- added support for parallel local template rendering
- added support for concurrent template writes

### v0.2.1

- added `--help` documentation with all available options

### v0.2.0

- added support for line filter functionality (i.e. pipe replacement string through std input stream from another command)
- added new `--trimnl` option to remove newline value from replacement string
- refactored ink.go source file template rendering handling to a new function
- added new standard input stream validator function


### v0.1.0

- initial release