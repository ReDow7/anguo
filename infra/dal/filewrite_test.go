package dal

import (
	"os"
	"testing"
)

func TestFileWriteAndRead(t *testing.T) {
	toWrite := "/abcde\nfghijk"
	pwd, _ := os.Getwd()
	err := WriteToFileOverWrite(pwd+"/tmp.sec", toWrite)
	if err != nil {
		t.Errorf("error when test write file %v", err)
		return
	}
	read, err := ReadFromFile("tmp.sec")
	if err != nil {
		t.Errorf("error when test read file %v", err)
		return
	}
	if read != toWrite {
		t.Errorf("content mismatch %v vs %v", read, toWrite)
		return
	}
}
