package efile

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	//写入log文件
	file := LogFileName("requests")
	reqByte := []byte("sdafafsad")
	err := WriteFile(file, reqByte, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		t.Log(err)
	}
}
