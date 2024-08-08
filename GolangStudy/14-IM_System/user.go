package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	conn   net.Conn
	C      chan string
	server *Server
}

// 创建一个用户API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()

	return user
}

// 监听当前user channel的方法，一旦有消息，就直接发送给客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}

// 用户的上线业务
func (user *User) Online() {
	//用户上线，将用户添加到onlneMap中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	//广播当前用户上线消息
	user.server.BroadCast(user, "already online")
}

// 用户的下线业务
func (user *User) Offline() {
	//用户下线，将用户从onlineMap中删除
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	//广播当前用户下线消息
	user.server.BroadCast(user, "offline")
}

// 给当前user对应的客户端发送消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有哪些
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":online...\n"
			u.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式: rename|张三
		newName := strings.Split(msg, "|")[1]

		//判断newName是否存在
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("present username already exists\n")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.Name = newName
			user.SendMsg("You have already updated username: " + user.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:4] == "to|" {
		//1.获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("msg info formation is not correct.\n")
			return
		}

		//2.根据用户名获取user对象
		remoteUser, ok := user.server.OnlineMap[remoteName]
		if !ok {
			user.SendMsg("user doesn't exist\n")
			return
		}

		//3.获取消息内容，通过对方的user对象将消息内容发送出去
		contend := strings.Split(msg, "|")[2]
		if contend == "" {
			user.SendMsg("contend is empty.\n")
			return
		}
		remoteUser.SendMsg(user.Name + " say :" + contend)
	} else {
		user.server.BroadCast(user, msg)
	}
}
