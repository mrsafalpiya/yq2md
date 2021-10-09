package main

import (
	"errors"
	"os"
)

func getFileNameFromArgs() (string, error) {
	for _, arg := range os.Args[1:] {
		if !isFirstCharacterADash(&arg) {
			err := getFileStat(&arg)
			if err != nil {
				return "", err
			}

			return arg, nil
		}
	}

	return "", errors.New("No filename argument is provided")
}

func isFirstCharacterADash(str *string) bool {
	if (*str)[0] == '-' {
		return true
	}

	return false
}

func getFileStat(fileLoc *string) error {
	_, err := os.Stat(*fileLoc)
	if err != nil {
		return err
	}

	return nil
}
