package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//获取命令行参数,用命令传递文件go run send.go go.mod
	list := os.Args
	filepath := list[1]
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("os.Stat err", err)
		return
	}
	filename := fileInfo.Name()
	conn, err := net.Dial("tcp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.Dialt err", err)
		return
	}
	_, err = conn.Write([]byte(filename))
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err", err)
		return
	}
	if string(buf[:n]) == "ok" {
		sendFile(conn, filepath)
	}

}
func sendFile(conn net.Conn, filepath string) {
	//打开要传输的文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("os.Open err", err)
		return
	}
	buf := make([]byte, 4096)
	//循环读取文件内容，写入远程连接
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			fmt.Println("文件发送完成")
			return
		}
		if err != nil {
			fmt.Println("file.Read err:", err)
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}
	}
}
