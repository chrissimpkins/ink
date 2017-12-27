# <img src="https://raw.githubusercontent.com/chrissimpkins/ink/images/img/ink-logo-crunch.png">
### A fast, flexible stream editor

[![Build Status](https://semaphoreci.com/api/v1/sourcefoundry/ink/branches/master/badge.svg)](https://semaphoreci.com/sourcefoundry/ink) [![Build status](https://ci.appveyor.com/api/projects/status/21si0rtxx9q36cad/branch/master?svg=true)](https://ci.appveyor.com/project/chrissimpkins/ink/branch/master) [![Go Report Card](https://goreportcard.com/badge/github.com/chrissimpkins/ink)](https://goreportcard.com/report/github.com/chrissimpkins/ink)

# What is ink?

ink is an open source stream editor that supports rendering of text replacements in pre-formatted text files (templates). The `ink` executable is compiled from Go programming language source for use on [Linux, macOS, and Windows platforms](https://github.com/chrissimpkins/ink/releases/latest).  It was designed to provide a simple approach to get command line executable text data into text templates.

It features:

- Unicode support
- line filter style stream editor support with a twist (pipe replacement text from another application to ink, render your template tokens with the piped standard input text, then pipe the rendered text to the standard output stream for further text processing or file writes)
- local and remotely stored (GET request accessible) template support
- parallel multi-template renders
- a simple [built-in text template syntax](#ink-template-specification) with `{{ ink }}` tokens
- a flexible [user-defined text template syntax]((#user-defined-template-specification)) that supports *nearly any text replacement token*™ that you'd like to use.  These tokens are defined at rendering time on the command line
- automated outfile path writes that are built into the template specifications with the `.in` extension appended to the intended render artifact path
- simple find/replace string literal and regular expression match substitutions (with [re2 regex syntax](https://github.com/google/re2/wiki/Syntax)) that are defined on the command line with `--find` and `--replace` options

### Example

##### CSS template file

The CSS template is available on the local file path `hack.css.in` in this example; however, it is possible to store the template file remotely on a server where it is accessible by GET request and use the URL in place of the local file path to accomplish the same local text file render as demonstrated in this example.

```css
/*!
 *  Hack typeface https://github.com/source-foundry/Hack
 *  License: https://github.com/source-foundry/Hack/blob/master/LICENSE.md
 */
/* FONT PATHS
 * -------------------------- */
@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-regular.woff2?sha={{ ink }}') format('woff2'), url('fonts/hack-regular.woff?sha={{ ink }}') format('woff');
  font-weight: 400;
  font-style: normal;
}

@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-bold.woff2?sha={{ ink }}') format('woff2'), url('fonts/hack-bold.woff?sha={{ ink }}') format('woff');
  font-weight: 700;
  font-style: normal;
}

@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-italic.woff2?sha={{ ink }}') format('woff2'), url('fonts/hack-italic.woff?sha={{ ink }}') format('woff');
  font-weight: 400;
  font-style: italic;
}

@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-bolditalic.woff2?sha={{ ink }}') format('woff2'), url('fonts/hack-bolditalic.woff?sha={{ ink }}') format('woff');
  font-weight: 700;
  font-style: italic;
}
```

##### Render template with a git commit SHA1 short code

The following is executed in the repository under git version control:

```
$ git log --pretty=format:'%h' --abbrev-commit -1 | ink hack.css.in
```

##### Rendered CSS file

After execution of the above command, the rendered CSS file with git commit SHA1 short code stamps in the font URL is available on the path `hack.css`:

```css
/*!
 *  Hack typeface https://github.com/source-foundry/Hack
 *  License: https://github.com/source-foundry/Hack/blob/master/LICENSE.md
 */
/* FONT PATHS
 * -------------------------- */
@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-regular.woff2?sha=db337ca') format('woff2'), url('fonts/hack-regular.woff?sha=db337ca') format('woff');
  font-weight: 400;
  font-style: normal;
}

@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-bold.woff2?sha=db337ca') format('woff2'), url('fonts/hack-bold.woff?sha=db337ca') format('woff');
  font-weight: 700;
  font-style: normal;
}

@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-italic.woff2?sha=db337ca') format('woff2'), url('fonts/hack-italic.woff?sha=db337ca') format('woff');
  font-weight: 400;
  font-style: italic;
}

@font-face {
  font-family: 'Hack';
  src: url('fonts/hack-bolditalic.woff2?sha=db337ca') format('woff2'), url('fonts/hack-bolditalic.woff?sha=db337ca') format('woff');
  font-weight: 700;
  font-style: italic;
}

```

## Installation

### Approach 1: Install the pre-compiled binary executable file

Download the latest compiled release file for your operating system and architecture from [the Releases page](https://github.com/chrissimpkins/ink/releases/latest).

#### Linux / macOS

Unpack the tar.gz archive and move the `ink` executable file to a directory on your system PATH (e.g. `/usr/local/bin`).  This can be performed by executing the following command in the root of the unpacked archive:

```
$ mv ink /usr/local/bin/ink
```

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

#### Windows

Unpack the zip archive and move the `ink.exe` executable file to a directory on your system PATH. See [details here](https://stackoverflow.com/questions/4822400/register-an-exe-so-you-can-run-it-from-any-command-line-in-windows) for more information about how to do this.

There are no dependencies contained in the archive.  You can delete all downloaded archive files after the above step.

### Approach 2: Compile from the source code and install

You must install the Go programming language (which includes the `go` tool) in order to compile the project from source.  Follow the [instructions on the Go download page](https://golang.org/dl/) for your platform. 

Once you have installed Go and configured your settings so that Go executables are installed on your system PATH, use the following command to (1) pull the master branch of the ink repository; (2) compile the ink executable from source for your platform/architecture configuration; (3) install the executable on your system:

```
$ go get github.com/chrissimpkins/ink
```

## Usage

### Default Behavior

The default behavior of `ink` is to render text replacements at `{{ ink }}` tokens in one or more local or remote template files with replacement text that is defined by the `--replace=` command line flag value or by text data piped in through the standard input stream.  By default, a file write of the rendered template text takes place on the same directory path as the template file for local templates and the current working directory for remote templates. The rendered file path is defined as the template file path with the `.in` extension removed.

### Syntax

#### Local template file rendering

The following approach uses the built-in ink template syntax and file extension format to identify text replacement sites in the source document (see details below).

```
$ ink --replace=[replacement string] [options] [template path 1]...[template path n]
```

or 

```
$ [executable command stdout stream] | ink [options] [template path 1]...[template path n]
```

#### Remote template file rendering

The following approach uses the built-in ink template syntax and file extension format to identify text replacement sites in the source document (see details below).

```
$ ink --replace=[replacement string] [options] [template URL 1]...[template URL n]
```

or

```
$ [executable command stdout stream] | ink [options] [template URL 1]...[template URL n]
```

#### User-defined token substitutions

`ink` supports user-defined text replacement tokens in the source document with the `--find=` command line option. This option permits you to define your own template token format and render text replacements as you would with template files that follow the ink template specification.  This approach also permits use of ink as a stream editor for routine find/replace text substitutions with string literal or regular expression pattern token matches in the source document. `ink` supports the [re2 regular expression syntax](https://github.com/google/re2/wiki/Syntax). 

```
$ ink --find=[find string value or regular expression pattern] --replace=[replacement string] [options] [template path 1]...[template path n]
```

or

```
$ [executable command stdout replacement stream] | ink --find=[find string value or regular expression pattern] [options] [template path 1]...[template path n]
```

Remote files can be used as the source template text by replacing the local template file paths with one or more GET request accessible URL as shown in the previous usage examples.

Please see further details on the command line syntax for regular expression vs. string literal matching below.

### ink Options

- `--find=` : find string literal value or regular expression pattern for user defined template tokens. Regular expressions must follow the [re2 syntax](https://github.com/google/re2/wiki/Syntax).
- `-h, --help` : application help
- `--lint` : lint a template file for validity using the template file specifications
- `--replace=` : replacement string literal value for text substitutions
- `--stdout` : write rendered text to standard output stream
- `--trimnl` : trim newline value from replacement string (intended for use with data piped through stdin stream)
- `--usage` : application usage
- `-v, --version` : application version

### How to define a replacement string on the command line

The replacement text for your template file can either be piped to `ink` through the standard input stream or you can include the `--replace=[replacement string]` option in the command.  These are mutually exclusive and one of the approaches is mandatory with each command.

The following examples demonstrate how to achieve replacements with the same string literal using each approach:

```
$ echo "abcd123" | ink template.txt.in
$ ink --replace=abcd123 template.txt.in
```

and these examples demonstrate how to evaluate command line expressions and use the standard output data as the replacement text with each approach:

```
$ date | ink template.txt.in
$ ink --replace="$(date)" template.txt.in
```

### How to modify text in the replacement strings from other applications

#### Trim newline characters from replacement strings

Some command line executables include a newline character following the standard output text (including the `echo` executable that was used in examples above).  This is not always desirable in the replacement substring that is used in your template files.  To remove the newline character, include the `--trimnl` option in your command:

```
$ echo "abcd123" | ink --trimnl template.txt.in
```

### How to define a string literal token for text replacements

By default, `ink` uses the syntax `{{ ink }}` to identify text replacement tokens in the template document.  You have the option to define your own replacement tokens with the `--find=` option. Define the string literal value with the `--find=` option like this:

```
$ ink --find="[[ date ]]" --replace="January 1, 2000" template.txt.in
```

The above command identifies sites in the template text that match the string literal value `[[ date ]]` and substitute the replacement text at these locations.

Include double quotes around the string value if you include special shell characters.

### How to define a regular expression pattern token for text replacements

The `--find=` option also supports the use of regular expression patterns ([re2 syntax](https://github.com/google/re2/wiki/Syntax)) to identify text substitution tokens in the source text.  Use the regex definition syntax `--find={{regular expression pattern}}` to define the `--find=` request as a regular expression match substitution.  

For example, to define the regular expression match pattern `\d+\s+` with the `--find=` option, use the following syntax:

```
$ ink --find="{{\d+\s+}}" --replace="five " template.txt.in
```

The opening `{{` and closing `}}` character combination delimiters signify that the contents represent a regular expression pattern.  Do not include space characters between the opening and closing `{{` and `}}` delimiters unless you intend for these characters to be part of the regular expression pattern.  Use double quotes around the regular expression definition on platforms that treat `{` and `}` as special shell characters.

### How to pipe a rendered template to the standard output stream

By default, `ink` writes the rendered text to a file located in the same directory as the template file on a file path that is defined by the removal of the `.in` file extension.  You can modify this behavior to pipe the data through the standard output stream instead of writing to disk by including the `--stdout` option in your command.

For example:

```
$ ink --replace=abcd123 --stdout template.txt.in
```

This will permit you to view the rendered text in your terminal, pipe it to another application for further text processing, or write the file to disk with a file path that you specify.

Here is a Linux/macOS example of a pipeline to and from `ink` with a file write to the path `finalfile.txt` after further processing in the (fake) application `cooltxt`:

```
$ echo "abcd123" | ink --stdout template.txt.in | cooltxt --dothings > finalfile.txt
```

### How to validate a template file

Use the `--lint` option to confirm that a local or remote template file meets the [ink and user-defined template specifications](#template-file-specifications):

```
$ ink --lint template.txt.in
$ ink --lint https://somesite.org/template.txt.in
```
 
## Template File Specifications

"Template" is defined as any local or remote text file that is used as the source for text substitutions by inclusion of text replacement "tokens".

"Token" is defined as the set of ordered, case-sensitive glyphs that are intended as a site for text substitution within "Templates".

"Replacement Text" is defined as the text intended for substitution at the site of a "Token" in a "Template".

"Outfile" is defined as a text file path that is the rendering artifact of the `ink` executable.  Note that this is intentionally distinct from user implemented (e.g. operating system/shell specific shell idiom like `>`) approaches to file writes that are not performed by the ink executable.

### ink template specification

The "ink template" is specified as follows:

- Templates that are rendered to Outfiles MUST be defined by a path that includes the intended file path of the Outfile with the addition of the extension `.in`.
- Templates that are used to pipe rendered text data to the standard output stream do not have a specified file path format.  Users may define any local or remote path when the `--stdout` option is used.  The addition of a `.in` extension to the desired render artifact file path for these Templates is RECOMMENDED when file writes are performed with these streamed data.
- The Template MAY include zero or more Tokens that are defined in a case-sensitive manner as `{{ink}}` or `{{ ink }}`.
- The Template MAY include zero or more Tokens that are defined in a case-sensitive manner as `{{.Ink}}` or `{{ .Ink }}`.
- All Token glyphs up to and including the initial `{` and final `}` glyphs MUST be replaced with Replacement Text during each execution of the renderer.
- All Template Tokens MUST be replaced with Replacement Text during each execution of the renderer.

An example of a valid ink template is:

```
Email: {{ ink }}
Date: January 1, 2000
Subject: About things
```

An example of an invalid ink template is:

```
Email: {{ email }}
Date: January 1, 2000
Subject: About things
```

### User-defined template specification

The "User-defined template" is specified as follows:

- Templates MUST NOT include two adjacent `{` glyphs, zero or more glyphs that are not `}`, followed by two adjacent `}` glyphs in the Template.  The `{{ token }}` formatting is a protected part of the ink template specification.
- Templates that are rendered to Outfiles MUST be defined by a path that includes the intended file path of the rendered text outfile with the addition of the extension `.in`.
- Templates that are used to pipe rendered text data to the standard output stream do not have a specified source file path format.  Users may define any local or remote path when the `--stdout` option is used.  The addition of a `.in` extension to the desired render artifact file path for these Templates is RECOMMENDED when file writes are performed with these streamed data.
- All Token glyphs, in the order and case-sensitive definition specified on the command line, MUST be replaced with the Replacement Text during each execution of the renderer. 
- All Template Tokens MUST be replaced with Replacement Text during each execution of the renderer.

An example of a valid user-defined template is:

```
Email: [[ email ]]
Date: %date%
Subject: << SUBJECT >>
```

An example of an invalid user-defined template is:

```
Email: {{ .Email }}
Date: {{ date }}
Subject: {{ Subject }}
```

## License

[MIT License](https://github.com/chrissimpkins/ink/blob/master/LICENSE)

