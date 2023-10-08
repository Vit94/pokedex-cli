package main

import (
	"errors"
	"strings"
)

func prepareInput(raw string) ([]string, error) {
	lowered := strings.ToLower(raw)
	strArr := strings.Fields(lowered)
	if len(strArr) == 0 {
		return nil, errors.New("invalid input")
	}
	return strArr, nil
}
