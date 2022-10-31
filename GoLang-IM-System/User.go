package main

import (
	"net"
	"strings"
)

type User struct {
	name    string
	addr    string
	channel chan string
	conn    net.Conn
	server  *Server
}

func NewUser(conn net.Conn, server *Server) *User {

	user := &User{
		addr:    conn.RemoteAddr().String(),
		channel: make(chan string),
		conn:    conn,
		server:  server,
	}
	go user.listenMessage()
	return user
}

func (this *User) listenMessage() {
	for {
		/**
		<-chan读取chan 数据
		chan<- 写数据
		*/
		var message = <-this.channel
		this.conn.Write([]byte(message + "\n"))
	}
}

func (this *User) online() {
	//加锁操作
	this.server.mapLock.Lock()
	//map[key]=value 添加数据
	this.server.onlineMap[this.name] = this
	//释放锁
	this.server.mapLock.Unlock()

	//发送消息
	this.server.broadcastMessage(this, "上线")
}

func (this *User) offline() {
	//加锁操作
	server := this.server
	server.mapLock.Lock()
	//map[key]=value 添加数据
	delete(server.onlineMap, this.name)
	//释放锁
	server.mapLock.Unlock()
	server.broadcastMessage(this, "下线")
}

func (this *User) sendMsg(message string) {
	this.conn.Write([]byte(message))
}

func (this *User) doMessage(message string) {

	if message == "who" {
		this.server.mapLock.Lock()
		for _, user := range this.server.onlineMap {
			onLineMessage := "[" + user.addr + "]" + user.name + ":" + "在线....\r\n"
			this.sendMsg(onLineMessage)
		}
		this.server.mapLock.Unlock()
	} else if len(message) > 7 && message[:7] == "rename|" {
		newName := strings.Split(message, "|")[1]

		_, ok := this.server.onlineMap[newName]
		if ok {
			this.sendMsg("当前用户名已被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.onlineMap, this.name)
			this.server.onlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.name = newName
			this.sendMsg("已更新用户名[" + this.name + "]")
		}
	} else {

		this.server.broadcastMessage(this, message)
	}
}
