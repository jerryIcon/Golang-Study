package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户列表
	OnlieMap map[string]*User
	mapLock  sync.RWMutex

	// 广播Message
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:       ip,
		Port:     port,
		OnlieMap: make(map[string]*User),
		Message:  make(chan string),
	}
	return server
}

// 监听Message 全局广播,将Message发送给所有用户
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		// 将Message发送给所有用户
		this.mapLock.Lock()
		for _, user := range this.OnlieMap {
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听收到的 读消息
func (this *Server) ListenRead(conn net.Conn, user *User, isLive chan bool) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			// 用户下线
			this.UserOff(user)
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("Conn Read err:", err)
			return
		}
		// 重置定时器
		isLive <- true
		// 提取用户的消息（去除 \n）
		msg := string(buf[:n-1])
		// 广播
		user.DoMessage(msg)

	}
}

func (this *Server) Handler(conn net.Conn) {
	// fmt.Println("accept doHandler()......")

	// 创建用户
	user := NewUser(conn, this)

	// 创建销毁用户定时器
	isLive := make(chan bool)

	// 用户上线
	this.UserOnlie(user)

	// 接收用户的广播消息
	go this.ListenRead(conn, user, isLive)

	// 阻塞
	select {
	case <-isLive:
		// 当前用户是活跃的 应该重置定时器
		// 不做任何事情 为了激活select 更新下面的定时器
	// 定时器
	case <-time.After(time.Second * 10):
		// 超时

		// 将User关闭
		user.sendMsg("自动下线\n")

		close(user.C)

		conn.Close()

		//退出当前hanlder
		return // rutime.Goexit()
	}
}

// 用户上线
func (this *Server) UserOnlie(user *User) {
	// 用户上线， 将用户加入到online中
	this.mapLock.Lock()
	this.OnlieMap[user.Name] = user
	this.mapLock.Unlock()
	// 广播当前用户上线消息
	this.BroadCast(user, "已上线")
}

// 用户下线
func (this *Server) UserOff(user *User) {
	// 用户上线， 将用户加入到online中
	this.mapLock.Lock()
	delete(this.OnlieMap, user.Name)
	this.mapLock.Unlock()
	this.BroadCast(user, "已下线")
}

// 开始服务
func (this *Server) Start() {

	defer fmt.Println("tcp server end")

	// 监听服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Printf("tcp start error %s:%d err:%s\n", this.Ip, this.Port, err)
		return
	}
	fmt.Printf("tcp server start %s:%d.....\n", this.Ip, this.Port)

	// close
	defer listener.Close()

	// 启动监听Message的goroutine
	go this.ListenMessage()

	// accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connection accept error:", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}

}
