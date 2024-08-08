package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       -1,
	}

	//连接server
	conn, err := net.Dial("tcp", fmt.Sprint("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	//返回对象
	return client
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1. public chat mode")
	fmt.Println("2. secret chat mode")
	fmt.Println("3. update username")
	fmt.Println("0. exit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>please input valid number>>>")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>Please input username:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

func (client *Client) PublicChat() {
	//提示用户输入信息
	var chatMsg string

	fmt.Println("Please input content: 'exit' to exit")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//消息不为空时发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println("Please input content: 'exit' to exit")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err:", err)
		return
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>Please input user [username]: 'exit' to exit")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>Please input chat content: 'exit' to exit")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>Please input chat content: 'exit' to exit")
		fmt.Scanln(&chatMsg)
	}
	client.SelectUsers()
	fmt.Println(">>>Please input user [username]: 'exit' to exit")
	fmt.Scanln(&remoteName)
}

func (client *Client) Run() {
	for client.flag != 0 {
		for !client.menu() {
		}

		//根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			//公聊模式
			client.PublicChat()
		case 2:
			//私聊模式
			client.PrivateChat()
		case 3:
			//更新用户名
			client.UpdateName()
		}
	}
}

// 处理server回应的消息，直接显示到标准输出
func (client *Client) DealResponse() {
	//一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "120.0.0.1", "set server ip address")
	flag.IntVar(&serverPort, "port", 8888, "set server port value")
}

func main_() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>连接服务器失败>>>>>>")
		return
	}

	//单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>>连接服务器成功>>>>>")

	//启动客户端的业务
	client.Run()
}
