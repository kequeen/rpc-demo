package main

import (
	"fmt"
	"runtime"
	"time"
)

func sender(ch chan int, message int) {
	fmt.Println("Sender: Sending", message)
	ch <- message // 发送数据到通道
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("Sender: Sent", message)
	ch <- message // 发送数据到通道
}

func receiver(ch chan int) {
	// fmt.Println("Receiver: Waiting for data")
	// message := <-ch // 从通道接收数据，此处会阻塞等待
	// fmt.Println("Receiver: Received", message)
	// message = <-ch // 从通道接收数据，此处会阻塞等待
	// fmt.Println("Receiver: Received", message)
}

func main() {
	// ch := make(chan int, 2) // 创建不带缓冲区的通道

	// go sender(ch, 100) // 启动发送协程
	// go receiver(ch)    // 启动接收协程

	// go sender(ch, 100)
	// message := 100
	// fmt.Println("Sender: Sending", message)
	// ch <- message // 发送数据到通道

	// message1 := <-ch // 从通道接收数据，此处会阻塞等待

	// fmt.Println("Recevice", message1)

	// time.Sleep(time.Millisecond) // 等待协程执行完毕

	fmt.Println("CPU 核心数量:", runtime.NumCPU())
	getGmpInfo()
}

func getGmpInfo() {
	numM := runtime.GOMAXPROCS(0)
	fmt.Println("当前 M 的数量:", numM)
}
