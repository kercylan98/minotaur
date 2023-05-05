package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "server.natappfree.cc:37775")
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close() // 关闭TCP连接
	inputReader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			buf := [512]byte{}
			n, err := conn.Read(buf[:])
			if err != nil {
				continue
			}
			fmt.Println(string(buf[:n]))
		}
	}()
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err := conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
	}
}
