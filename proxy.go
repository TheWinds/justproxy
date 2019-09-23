package justproxy

import (
	"fmt"
	"io"
	"log"
	"net"
)

func Proxy(confFileName string) {
	config, err := LoadConfig(confFileName)
	if err != nil {
		log.Fatalf("can not read config file: %v", err)
	}
	if len(config.Proxys) == 0 {
		fmt.Println("nothing to proxy,exit")
		return
	}
	fmt.Println("all proxy:")
	for _, item := range config.Proxys {
		fmt.Printf("\t %s=>%s\n", item.Src, item.Dest)
	}
	proxyAll(config)
}

func proxyAll(config *Config) {
	for _, proxyItem := range config.Proxys {
		go func(item *ProxyItem) {
			listener, err := net.Listen("tcp", item.Src)
			if err != nil {
				log.Fatalf("error during listen %s: %v\n", item.Src, err)
			}
			for {
				srcConn, err := listener.Accept()
				if err != nil {
					log.Printf("error during accept src conn(%s): %v\n", item.Src, err)
					continue
				}
				go handleProxy(srcConn, item)
			}
		}(proxyItem)
	}
	select {}
}

func handleProxy(srcConn net.Conn, proxyItem *ProxyItem) {
	defer srcConn.Close()
	destConn, err := net.Dial("tcp", proxyItem.Dest)
	if err != nil {
		log.Printf("error during dail dest conn(%s): %v\n", proxyItem.Dest, err)
	}
	defer destConn.Close()
	overChan := make(chan bool)
	go cpoyIO(destConn, srcConn, overChan)
	go cpoyIO(srcConn, destConn, overChan)
	<-overChan
}

func cpoyIO(dst io.Writer, src io.Reader, overChan chan bool) {
	_, err := io.Copy(dst, src)
	if err != nil {
		overChan <- false
		log.Printf("buffer copy error: %v\n", err)
		return
	}
	overChan <- true
}
