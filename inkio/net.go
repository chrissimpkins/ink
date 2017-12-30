package inkio

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetRequest performs a GET request for a templateURL url string with a 30 second timeout
func GetRequest(templateURL string) (string, error) {
	emptystring := "" // returned with errors
	client := &http.Client{
		Timeout: 30 * time.Second, // GET request timeout is 30 seconds
	}
	req, _ := http.NewRequest(
		"GET", templateURL, nil,
	)
	req.Header.Add("Accept", "text/*")
	req.Header.Add("User-Agent", "ink/1.0")

	resp, resperr := client.Do(req)
	if resperr != nil {
		return emptystring, resperr
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return emptystring, fmt.Errorf("%s returned a non-200 response status: %s", templateURL, resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
