package main

import "errors"

var ErrInvalidArgs error = errors.New("Error: Please provide exactly one input file and one output file")

func ParseArgs(args []string) (string, string, error) {
	if len(args) != 2 {
		return "", "", ErrInvalidArgs
	}
	return args[0], args[1], nil
}
