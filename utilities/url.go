package utilities

import (
	"net/url"
	"strings"
)

// GetURLFilePath parses a Url string and returns the filepath at the final path position of the URL and error
func GetURLFilePath(Url string) (string, error) {
	u, err := url.Parse(Url)
	if err != nil {
		return "", err
	}
	if strings.Contains(u.Path, "/") {
		urlslice := strings.Split(u.Path, "/")
		return urlslice[len(urlslice)-1], err
	}
	return u.Path, err
}
