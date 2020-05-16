package clear_static

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"wisdom-portal/wisdom-portal/logger"
)

const (
	// 10分钟 600秒
	diffTime = 60 * 10
)

func GetAllFile(pathDir string, fileChan chan string, timeSleep time.Duration) {
	for {
		fileInfo, err := ioutil.ReadDir(pathDir)
		if err != nil {
			logger.Error("getAllFile error!, err: " + err.Error())
			return
		}
		for _, file := range fileInfo {
			strAbsPath, err := filepath.Abs(pathDir + "/" + file.Name())
			if err != nil {
				logger.Error("file abs path fail!, err: " + err.Error())
				return
			}
			fileChan <- strAbsPath
			logger.Debug("fileAbsPath put into the channel success: " + strAbsPath)
		}
		time.Sleep(timeSleep)
	}
}

func CheckFileDiffTime(fileChan chan string) {
	// range会一直堵塞直到有人调用了close
	for pathFile := range fileChan {
		nowTime := time.Now().Unix() // 当前时间
		// path 代表文件绝对路径
		// file 代表文件元数据详情
		err := filepath.Walk(pathFile, func(path string, file os.FileInfo, err error) error {
			if file == nil {
				return err
			}
			fileTime := file.ModTime().Unix()
			if (nowTime - fileTime) >= diffTime {
				logger.Info("take out fileAbsPath from the channel success: " + path)
				logger.Info("delete file success: " + path)
				err = os.RemoveAll(path)
				if err != nil {
					return err
				}
			} else {
				logger.Debug(fmt.Sprintf("%s file did not reach deletion time", path))
			}
			return nil
		})
		if err != nil {
			logger.Error("filepath.Walk() returned, err: " + err.Error())
			return
		}
	}
}
