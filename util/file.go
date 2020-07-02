package util

import (
	"io"
	"os"
	"strings"
	//"gfx_project/logsystem"
)

//读取大文件 返回二进制流
func readFile(filePath string) (binary []byte ,err error) {
	file ,err := os.Open(filePath)
	if err != nil {
		return
	}

	defer file.Close()
	//defer os.Remove(filePath)  //删除指定文件

	bEveryTimes := make([]byte, 1048576)
	var n int
	for {
		n, err = file.Read(bEveryTimes)
		if err != nil {
			break
		}
		binary = append(binary, bEveryTimes[:n]...)
	}
	return
}


func PareseFile(path string) (lines []string ,err error){
	binary ,err := readFile(path)
	if err != nil && err != io.EOF{
		return
	}

	fileBytes := []byte{}
	fileBytes = append(fileBytes, binary...)

	str := string(fileBytes)
	lines = strings.Split(strings.TrimSpace(str), "\n")

	return
}




