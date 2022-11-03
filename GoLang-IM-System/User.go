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

func (u *User) listenMessage() {
	for {
		/**
		<-chan读取chan 数据
		chan<- 写数据
		*/
		var message = <-u.channel
		_, err := u.conn.Write([]byte(message + "\n"))
		if err != nil {
			return
		}
	}
}

func (u *User) online() {
	//加锁操作
	u.server.mapLock.Lock()
	//map[key]=value 添加数据
	u.server.onlineMap[u.name] = u
	//释放锁
	u.server.mapLock.Unlock()

	//发送消息
	u.server.broadcastMessage(u, "上线")
}

func (u *User) offline() {
	//加锁操作
	server := u.server
	server.mapLock.Lock()
	//map[key]=value 添加数据
	delete(server.onlineMap, u.name)
	//释放锁
	server.mapLock.Unlock()
	server.broadcastMessage(u, "下线")
}

func (u *User) sendMsg(message string) {
	_, err := u.conn.Write([]byte(message))
	if err != nil {
		return
	}
}

func (u *User) doMessage(message string) {

	if message == "who" {
		u.server.mapLock.Lock()
		for _, user := range u.server.onlineMap {
			onLineMessage := "[" + user.addr + "]" + user.name + ":" + "在线....\r\n"
			u.sendMsg(onLineMessage)
		}
		u.server.mapLock.Unlock()
	} else if len(message) > 7 && message[:7] == "rename|" {
		newName := strings.Split(message, "|")[1]

		_, ok := u.server.onlineMap[newName]
		if ok {
			u.sendMsg("当前用户名已被使用\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.onlineMap, u.name)
			u.server.onlineMap[newName] = u
			u.server.mapLock.Unlock()
			u.name = newName
			u.sendMsg("已更新用户名[" + u.name + "]")
		}
	} else {

		u.server.broadcastMessage(u, message)
	}
}
