package efile

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
	"work_report/config"
)

//生成任意文件
func FileName(dirname string, fileName string, ext string) string {
	file := fileName + "." + ext
	return filepath.Join(dirname, file)
}

//生成log文件路径
func LogFileName(prefix string) string {
	fileName := config.AppName + "_" + prefix + "_" + time.Now().Format("2006-01-02") + ".log"
	return filepath.Join(config.LogPath, fileName)
}

//写入文件
func WriteFile(fileDir string, data interface{}, op int, mode os.FileMode) error {
	txtByte, _ := json.Marshal(data)
	fl, err := os.OpenFile(fileDir, op, mode)
	if err != nil {
		return err
	}
	defer fl.Close()
	n, err := fl.Write(txtByte)
	if err == nil && n < len(txtByte) {
		return err
	}
	return nil
}

//判断文件是否存在  存在返回 true 不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
