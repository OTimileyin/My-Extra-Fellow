package main

import "os"

func ParseInputContent(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", ErrReadFile
	}
	return string(data), nil
}
