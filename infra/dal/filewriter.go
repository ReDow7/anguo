package dal

import (
	"fmt"
	"os"
)

func WriteToFileOverWrite(fileName string, body string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(body)
	if err != nil {
		return err
	}
	return nil
}

func WriteToNewFile(fileName string, body string) error {
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return fmt.Errorf("file %v already exsisted", fileName)
	}
	f, err := os.OpenFile(fileName, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(body)
	if err != nil {
		return err
	}
	return nil
}
