package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	Server *Server
}

// 创建一个用户API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		Server: server,
	}
	// 启动监听 user channel 消息 的 goroutine
	go user.ListenMessage()

	return user
}

// 给当前用户发送消息
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 修改用户名功能
func (this *User) SetName(name string) {
	this.Server.mapLock.Lock()
	// 删除原来名字
	delete(this.Server.OnlieMap, this.Name)
	// 修改名称
	this.Name = name
	this.Server.OnlieMap[name] = this
	this.Server.mapLock.Unlock()
}

// 用了处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 1.查询当前在线用户功能

		this.Server.mapLock.Lock()
		for _, user := range this.Server.OnlieMap {
			msg := "[" + user.Addr + "]" + user.Name + ":" + "在线\n"
			this.sendMsg(msg)
		}
		this.Server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 2.修改用户名功能

		newName := strings.Split(msg, "|")[1]
		// 判断名是否存在
		_, ok := this.Server.OnlieMap[newName]
		if ok {
			this.sendMsg("当前名称已存在！\n")
		} else {
			this.SetName(newName)
			this.sendMsg("更换用户名：" + newName + "成功\n")
		}
	} else {
		this.Server.BroadCast(this, msg)
	}
}

// 监听当前User channel 的方法 一旦有消息 就直接发生给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
