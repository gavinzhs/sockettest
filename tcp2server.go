package main

import (
    "net"
    "log"
    "fmt"
    "io")

const (
    ip = ""
    port = 3333
)

func main(){
    listen,err := net.ListenTCP("tcp",&net.TCPAddr{net.ParseIP(ip),port,""})
    if err != nil {
        log.Printf("监听失败,err:%v",err)
        return
    }

    log.Println("已初始化连接，等待客户端连接。。。")
    Server(listen)
}

func Server(listen *net.TCPListener) {
    for {
        conn,err := listen.AcceptTCP()
        if err != nil {
            log.Printf("接受客户端连接异常：%v",err)
            continue
        }
        log.Printf("客户端来自：%s",conn.RemoteAddr().String())
        defer conn.Close()
        go func (){
            data := make([]byte,2)
            fmt.Printf("data len:%v\n",len(data))
            for {
//                i,err := conn.Read(data)
                i, err := io.ReadFull(conn, data)
                fmt.Printf("data len:%v\n",len(data))
                log.Printf("客户端%s.发来数据：%s",conn.RemoteAddr().String(),string(data[0:i]))
                if err != nil {
                    log.Printf("读取客户端数据错误:%v",err)
                    break
                }
                conn.Write([]byte{'f','i','n','s','h'})
            }
        }()
    }
}