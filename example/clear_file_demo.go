package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	diffTime = 60 * 30
)

func getAllFile(pathDir string, fileChan chan string, timeSleep time.Duration) {
	for {
		fileInfo, err := ioutil.ReadDir(pathDir)
		if err != nil {
			fmt.Println("getAllFile error!, err: " + err.Error())
			return
		}
		for _, file := range fileInfo {
			strAbsPath, err := filepath.Abs(pathDir + "/" + file.Name())
			if err != nil {
				fmt.Println("file abs path fail!, err: " + err.Error())
				return
			}
			fileChan <- strAbsPath
			fmt.Println("投递成功: " + strAbsPath)
		}
		time.Sleep(timeSleep)
	}
}

func testChan(fileChan chan string, goInt int) {
	// range会一直堵塞直到有人调用了close
	for pathFile := range fileChan {
		fmt.Println("取出成功: " + pathFile + " 处理者: " + strconv.Itoa(goInt))
	}
}

func checkFileDiffTime(fileChan chan string) {
	// range会一直堵塞直到有人调用了close
	for pathFile := range fileChan {
		nowTime := time.Now().Unix() // 当前时间
		err := filepath.Walk(pathFile, func(path string, file os.FileInfo, err error) error {
			if file == nil {
				return err
			}
			fileTime := file.ModTime().Unix()
			if (nowTime - fileTime) > diffTime {
				fmt.Printf("Delete file %v !\r\n", path)
				err = os.RemoveAll(path)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println("filepath.Walk() returned, err: " + err.Error())
		}
	}
}

func main() {
	fileChan := make(chan string, 50)
	go getAllFile("/Users/heyang/GoCode/src/wisdom-portal/static/", fileChan, time.Second*10)
	for i := 0; i <= 5; i++ {
		go checkFileDiffTime(fileChan)
	}

	fmt.Println("test")
	time.Sleep(time.Second * 1000000)
}
