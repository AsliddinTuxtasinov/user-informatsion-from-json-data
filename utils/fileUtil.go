package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(fileName string) ([]byte, error) {
	if IsEmptyString(fileName) {
		return nil, errors.New("Empty string file name is not used")
	}

	bytes, err := ioutil.ReadFile(fileName)
	CheckError(err)
	return bytes, err

}
