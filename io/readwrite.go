package io

import "io/ioutil"

func ReadFileToString(filepath string) (string, error) {
	byteString, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(byteString), nil
}
