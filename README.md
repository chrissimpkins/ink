# <img src="https://raw.githubusercontent.com/chrissimpkins/ink/images/img/ink-logo-crunch.png">
### A fast, flexible text template renderer

[![Build Status](https://semaphoreci.com/api/v1/sourcefoundry/ink/branches/master/badge.svg)](https://semaphoreci.com/sourcefoundry/ink) [![Go Report Card](https://goreportcard.com/badge/github.com/chrissimpkins/ink)](https://goreportcard.com/report/github.com/chrissimpkins/ink)

# What is ink?

ink is an open source command line text file renderer that is built with Go. The ink executable is compiled for use on [Linux, macOS, and Windows platforms]((https://github.com/chrissimpkins/ink/releases/latest)).  It was designed to provide a simple approach to get your command line executable text data into pre-formatted text files.

It features:

- line filter support (pipe replacement text from other applications to ink, render your template with the standard input piped text, then pipe the rendered text to the standard output stream for file writes or further text processing)
- support for parallel multi-file renders from local and remotely stored (GET request accessible) templates
- a simple built-in text template format using `{{ ink }}` text labels
- extremely flexible user defined text template formatting that supports *any template format*™ that you'd like to use (you define it on the command line)

### Example

##### CSS template file

The CSS template is available on the path `hack.css.in` in this example; however, it is possible to store the template file remotely on a server where it is accessible by GET request and use the URL in place of the local file path to accomplish the same local text file render as demonstrated in this example.

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





