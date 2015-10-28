package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	log.Printf("这里是参数: %s", os.Args)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "命令参数不对 : %s\n", os.Args[0])
		os.Exit(1)
	}

	if ip := net.ParseIP(os.Args[1]); ip == nil {
		log.Fatalln("转化失败")
	} else {
		log.Printf("ip 为 : %s", ip)
		log.Printf("ip 为 : %d", ip)
	}
}
