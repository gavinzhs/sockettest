package main

import (
    "testing"
    "log"
    "net"
    "bufio"
    "fmt"
    "bytes")

func TestIp(t *testing.T) {
    ipsrc := "123.59.64.205"
    if ip := net.ParseIP(ipsrc); ip == nil {
        log.Fatalln("转化失败")
    } else {
        log.Printf("ip 为 : %s", ip)
        log.Printf("ip 为 : %d", ip)
    }
}

func TestTCPSocketClient(t *testing.T) {
    //    q := make(chan bool)
    ip := "192.168.2.47:9090"
    addr, err := net.ResolveTCPAddr("tcp", ip)
    if err != nil {
        log.Fatalf("resolve tcp addr err : %v", err)
    }

    conn, err := net.DialTCP("tcp", nil, addr)
    if err != nil {
        log.Fatalf("resolve tcp addr err : %v", err)
    }

    //    go func(conn *net.TCPConn) {
    //        defer conn.Close()
    //        reader := bufio.NewReader(conn)
    //        for {
    //            msg, err := reader.ReadString(byte('\n'))
    //            if err != nil {
    //                log.Printf("read string err : %v", err)
    //                q <- true
    //                break
    //            }
    //            log.Printf("收到了服务器给我的消息: %s", msg)
    //
    //            time.Sleep(time.Second * 5)
    //            log.Println("五秒后")
    //            send := "world"
    //            if msg == "world" {
    //                send = "hello"
    //            }
    //
    //            conn.Write([]byte(send + "\n"))
    //        }
    //
    //    }(conn)

    //    conn.Write([]byte("hello\n"))

    sms := make([]byte, 128)
    for {
        fmt.Print("请输入要发送的内容:")
        _, err := fmt.Scan(&sms)
        fmt.Printf("-------sms:%v", sms)
        if err != nil {
            log.Printf("数据输入异常：%v", err)
        }
        fmt.Printf("sms len:%v", len(sms))
        conn.Write(sms)
        buf := make([]byte, 128)
        c, err := conn.Read(buf)
        if err != nil {
            log.Printf("读取服务器数据异常:%v", err)
        }
        log.Println(string(buf[0:c]))
    }

    //    <-q
}

func TestTCPSocketService(t *testing.T) {
    port := ":9090"
    addr, err := net.ResolveTCPAddr("tcp", port)
    if err != nil {
        log.Fatalf("resolve tcp addr err : %v", err)
    }
    listener, err := net.ListenTCP("tcp", addr)
    if err != nil {
        log.Fatalf("resolve tcp addr err : %v", err)
    }

    i := 1
    for {
        conn, err := listener.AcceptTCP()
        log.Printf("你是第%d个接入进来的", i)
        i++
        if err != nil {
            log.Printf("accept err : %v", err)
            continue
        }

        go settleClientConn(conn)
    }
}

func settleClientConn(conn *net.TCPConn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    for {
        msg, err := reader.ReadString(byte('\n'))
        if err != nil {
            log.Printf("read string err : %v", err)
            break
        }
        log.Printf("我收到的msg : 11%s11", msg)
        response := "hello"
        if msg == "hello\n" {
            response = "world"
        }

        conn.Write([]byte(response + "\n"))
    }
}

func TestReadBuffer(t *testing.T) {
    buf := bytes.NewBuffer([]byte("abc"))
    o := make([]byte, 2)
    n, err := buf.Read(o)
    if err != nil {
        log.Fatalf("read err : %v", err)
    }

    log.Printf("o read num : %d, value : %d", n, o)

    o = make([]byte, 2)
    n, err = buf.Read(o)
    if err != nil {
        log.Fatalf("read err : %v", err)
    }

    log.Printf("o read num : %d, value : %d", n, o)

    o = make([]byte, 2)
    n, err = buf.Read(o)
    if err != nil {
        log.Fatalf("read err : %v", err)
    }

    log.Printf("o read num : %d, value : %d", n, o)
}
