package dal

import (
	"io/ioutil"
	"os"
)

func ReadFromFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
