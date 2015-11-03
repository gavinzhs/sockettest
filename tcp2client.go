package main

import (
    "net"
    "log"
    "fmt")

const (
    addr = "127.0.0.1:3333"
)

func main(){
    conn,err := net.Dial("tcp",addr)
    if err != nil {
        log.Printf("连接服务器失败:%v",err)
        return
    }
    log.Println("已连接服务器")
    defer conn.Close()
    Client(conn)
}

func Client(conn net.Conn){
    sms := make([]byte,128)
    for {
        fmt.Print("请输入要发送的内容:")
        _,err := fmt.Scan(&sms)
        fmt.Printf("-------sms:%v",sms)
        if err != nil {
            log.Printf("数据输入异常：%v",err)
        }
        fmt.Printf("sms len:%v",len(sms))
        conn.Write(sms)
        buf := make([]byte,2)
        c,err := conn.Read(buf)
        if err != nil {
            log.Printf("读取服务器数据异常:%v",err)
        }
        log.Println(string(buf[0:c]))
    }
}
