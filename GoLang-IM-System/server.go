package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	ip   string
	port int

	//在线用户列表
	onlineMap map[string]*User

	//锁
	mapLock sync.RWMutex

	//message channel
	serverChannel chan string
}

// NewServer 传入 string 类型ip  int类型端口 返回Server地址
func NewServer(ip string, port int) *Server {
	//*Server 是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值 .

	//&Server 取地址符号
	server := &Server{
		ip:            ip,
		port:          port,
		onlineMap:     make(map[string]*User),
		serverChannel: make(chan string),
	}
	//返回的是Sever地址
	return server
}

func (s *Server) listenMessage() {
	for {
		msg := <-s.serverChannel
		s.mapLock.Lock()
		for _, user := range s.onlineMap {
			user.channel <- msg
		}
		s.mapLock.Unlock()
	}
}

func (s *Server) broadcastMessage(user *User, message string) {
	//格式化消息
	sendMsd := fmt.Sprintf("[%s-%s]:%s", user.name, user.addr, message)
	//发送消息  写入channel
	fmt.Println(sendMsd)
	s.serverChannel <- sendMsd
}

// Handler 入参conn 连接 (this *server)隐式的指针
func (s *Server) Handler(conn net.Conn) {
	fmt.Println("当前建立连接成功")

	user := NewUser(conn, s)
	user.online()

	isAlive := make(chan bool)
	go func() {
		bytes := make([]byte, 4096)
		for {
			n, err := conn.Read(bytes)
			if n == 0 {
				user.offline()
			}
			if err != nil && err != io.EOF {
				fmt.Println("connect read err")
			}

			msg := string(bytes[:n-1])
			user.doMessage(msg)

			isAlive <- true
		}
	}()
	//阻塞
	for {
		select {
		case <-isAlive:
		case <-time.After(time.Second * 10):
			user.sendMsg("强制下线")
			close(user.channel)
			err := conn.Close()
			if err != nil {
				return
			}
			return

		}
	}

}

func (s *Server) start() {
	//开启TCP监听ip:port
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		fmt.Println("net, listen err:", err)
		return
	}
	//defer 会在函数结束之前运行
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			return
		}
	}(listen)

	//go 关键字新启动一个线程去执行
	go s.listenMessage()

	//执行 accept 等待连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		fmt.Println("等待客户端连接。。。")

		//启动 线程执行handler
		go s.Handler(conn)
	}
}
