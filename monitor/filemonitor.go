package monitor

import (
	"gfx_project/datacollection"
	"io"
	"os"
	"strings"
	//"sync"
	//"time"
	"fmt"

	"github.com/fsnotify/fsnotify"
)


type logFile struct {
	binary []byte
	timestamp int64
}

//读取文件
func readFile(fileName string) (logFile, error) {
	var logFile logFile

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("os open file error")
		return logFile, err
	}
	defer file.Close()
	//defer os.Remove(fileName)  //删除指定文件

	bEveryTimes := make([]byte, 1048576)
	var n int
	for {
		n, err = file.Read(bEveryTimes)
		if err != nil {
			if err != io.EOF {
				return logFile, err
			}
			break
		}
		logFile.binary = append(logFile.binary, bEveryTimes[:n]...)
	}

	logFile.timestamp = int64(32)
	//fmt.Printf("show file data = %v", logFile)
	return logFile, nil
}

//添加监听
func FileMonitor(filePath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		//logsystem.Gfxlog.Emerg(err.Error())
		fmt.Printf("watcher error = %v" , err)
	}
	defer watcher.Close()

	err = watcher.Add(filePath)
	if err != nil {
		//logsystem.Gfxlog.Emerg(err.Error())
		fmt.Printf( "watcher.Add error = %v \n" , err)
	}

	for {
		//fileBytes := []byte{}
		select {
		case event := <-watcher.Events:
			if event.Op & fsnotify.Create == fsnotify.Create {
				fmt.Println("event:" + event.String() +"\n")
				fmt.Println("created file:" + event.Name+"\n")

				typefile := GetFileType(event.Name)
				if typefile != -1 {
					go datacollection.AddData(typefile,event.Name)
				}
			}
		case err := <-watcher.Errors:
			//epgLog.err(err.Error())
			fmt.Printf("currnet error = %v" , err)
		}
	}
}

func GetFileType(str string ) int {
	arr := strings.Split(str ,".")
	if len(arr) == 1 {
		//光距
		return 2
	}

	if arr[1] == "csv" {   //日志文件
		return 0
	}else  if arr[1] == "txt"{ //分光器
		return  1
	}

	//光距
	return -1
}